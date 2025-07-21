package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/docker/mcp-registry/internal/mcp"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
)

func main() {
	listTools := flag.Bool("tools", false, "List the tools")
	pullCommunity := flag.Bool("pull-community", false, "Pull images that are not in the mcp/ namespace")

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: task build -- <server>")
		os.Exit(1)
	}

	if err := run(context.Background(), args[0], *listTools, *pullCommunity); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, name string, listTools bool, pullCommunity bool) error {
	server, err := servers.Read(filepath.Join("servers", name, "server.yaml"))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("server %s not found (did you already create it with `task create`?)", name)
		}
		return err
	}

	isMcpImage := strings.HasPrefix(server.Image, "mcp/")

	if isMcpImage {
		if err := buildMcpImage(ctx, server); err != nil {
			return err
		}
	} else {
		if !pullCommunity {
			return fmt.Errorf("server is not docker built (ie, in the 'mcp/' namespace), you must either build it yourself or pull it with `docker pull %s` if you want to use it", server.Image)
		}
		if err := pullCommunityImage(ctx, server); err != nil {
			return err
		}
	}

	if listTools {
		tools, err := mcp.Tools(ctx, server, false, false, false)
		if err != nil {
			return err
		}

		if len(tools) == 0 {
			fmt.Println()
			fmt.Println("No tools found.")
		} else {
			fmt.Println()
			fmt.Println(len(tools), "tools found.")
		}
	}
	fmt.Printf("\n-----------------------------------------\n\n")

	if isMcpImage {
		fmt.Println("✅ Image built as", server.Image)
	} else {
		fmt.Println("✅ Image pulled as", server.Image)
	}

	return nil
}

func buildDockerEnv(additionalEnv ...string) []string {
	env := []string{"PATH=" + os.Getenv("PATH")}
	
	// On Windows, Docker also needs ProgramW6432
	// See https://github.com/docker/mcp-registry/issues/79 for more details
	programW6432 := os.Getenv("ProgramW6432")
	if runtime.GOOS == "windows" && programW6432 != "" {
		env = append(env, "ProgramW6432="+programW6432)
	}
	
	return append(env, additionalEnv...)
}

func buildMcpImage(ctx context.Context, server servers.Server) error {
	projectURL := server.Source.Project
	branch := server.Source.Branch
	directory := server.Source.Directory

	client := github.New()

	repository, err := client.GetProjectRepository(ctx, projectURL)
	if err != nil {
		return err
	}

	if branch == "" {
		branch = repository.GetDefaultBranch()
	}

	sha, err := client.GetCommitSHA1(ctx, projectURL, branch)
	if err != nil {
		return err
	}

	gitURL := projectURL + ".git#"
	if branch != "" {
		gitURL += branch
	}
	if directory != "" && directory != "." {
		gitURL += ":" + directory
	}

	var cmd *exec.Cmd
	token := os.Getenv("GITHUB_TOKEN")

	if token != "" {
		cmd = exec.CommandContext(ctx, "docker", "buildx", "build", "--secret", "id=GIT_AUTH_TOKEN", "-f", server.GetDockerfile(), "-t", "check", "-t", server.Image, "--label", "org.opencontainers.image.revision="+sha, gitURL)
		cmd.Env = buildDockerEnv("GIT_AUTH_TOKEN=" + token)
	} else {
		cmd = exec.CommandContext(ctx, "docker", "buildx", "build", "-f", server.GetDockerfile(), "-t", "check", "-t", server.Image, "--label", "org.opencontainers.image.revision="+sha, gitURL)
		cmd.Env = buildDockerEnv()
	}

	cmd.Dir = os.TempDir()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func pullCommunityImage(ctx context.Context, server servers.Server) error {
	cmd := exec.CommandContext(ctx, "docker", "pull", server.Image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

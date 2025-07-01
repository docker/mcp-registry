package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docker/mcp-registry/internal/mcp"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
)

func main() {
	listTools := flag.Bool("tools", false, "List the tools")

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: task build -- <server>")
		os.Exit(1)
	}

	if err := run(context.Background(), args[0], *listTools); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, name string, listTools bool) error {
	server, err := servers.Read(filepath.Join("servers", name, "server.yaml"))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("server %s not found (did you already create it with `task create`?)", name)
		}
		return err
	}

	if !strings.HasPrefix(server.Image, "mcp/") {
		return fmt.Errorf("server is not docker built (ie, in the 'mcp/' namespace), you must either build it yourself or pull it with `docker pull %s` if you want to use it", server.Image)
	}

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
		cmd = exec.CommandContext(ctx, "docker", "buildx", "build", "--secret", "id=GIT_AUTH_TOKEN", "-t", "check", "-t", server.Image, "--label", "org.opencontainers.image.revision="+sha, gitURL)
		cmd.Env = []string{"GIT_AUTH_TOKEN=" + token, "PATH=" + os.Getenv("PATH")}
	} else {
		cmd = exec.CommandContext(ctx, "docker", "buildx", "build", "-t", "check", "-t", server.Image, "--label", "org.opencontainers.image.revision="+sha, gitURL)
		cmd.Env = []string{"PATH=" + os.Getenv("PATH")}
	}

	cmd.Dir = os.TempDir()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
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

	fmt.Println("✅ Image built as", server.Image)

	return nil
}

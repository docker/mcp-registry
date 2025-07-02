/*
Copyright © 2025 Docker, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"gopkg.in/yaml.v3"

	"github.com/docker/mcp-registry/internal/licenses"
	"github.com/docker/mcp-registry/internal/mcp"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
)

func main() {
	name := flag.String("name", "", "Name of the mcp server, name is guessed if not provided")
	category := flag.String("category", "", "Category for the mcp server (required) - [ai, data-visualization, database, devops, ecommerce, finance, games, communication, monitoring, productivity, search]")
	image := flag.String("image", "", "Image to use for the mcp server, instead of building from the repository")
	build := flag.Bool("build", true, "Build the image")
	listTools := flag.Bool("tools", true, "List the tools")

	flag.Parse()
	args := flag.Args()

	if *category == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	url := ""
	var additionalArgs []string
	if len(args) > 0 {
		url = args[0]
		additionalArgs = args[1:]
	}

	if err := run(ctx, url, *name, *category, *image, *build, *listTools, additionalArgs); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, buildURL, name, category, userProvidedImage string, build, listTools bool, args []string) error {
	projectURL := buildURL

	client := github.New()
	repository, err := client.GetProjectRepository(ctx, projectURL)
	if err != nil {
		return err
	}

	tags := github.FindTags(repository.Topics)

	detectedInfo, err := github.DetectBranchAndDirectory(projectURL, repository)
	if err != nil {
		return err
	}

	branch := detectedInfo.Branch
	directory := detectedInfo.Directory
	projectURL = detectedInfo.ProjectURL

	upstream := ""
	if repository.GetParent() != nil {
		upstream = repository.GetParent().GetHTMLURL()
	}

	sha, err := client.GetCommitSHA1(ctx, projectURL, branch)
	if err != nil {
		return err
	}

	refProjectURL := projectURL
	if upstream != "" {
		refProjectURL = upstream
	}

	guessedName := guessName(projectURL)
	if name == "" {
		name = strings.ToLower(guessedName)
	}

	tag := "mcp/" + name
	if userProvidedImage != "" {
		tag = userProvidedImage
	}

	title := strings.ToUpper(guessedName[0:1]) + guessedName[1:]
	if !strings.Contains(repository.GetDescription(), title+" MCP Server") {
		title += " (TODO)"
	}

	if !licenses.IsValid(repository.License) {
		fmt.Println("[WARNING] Project", projectURL, "is licensed under", repository.License.GetName(), "which may be incompatible with some tools")
	}

	if build && userProvidedImage == "" {
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
			cmd = exec.CommandContext(ctx, "docker", "buildx", "build", "--secret", "id=GIT_AUTH_TOKEN", "-t", "check", "-t", tag, "--label", "org.opencontainers.image.revision="+sha, gitURL)
			cmd.Env = []string{"GIT_AUTH_TOKEN=" + token, "PATH=" + os.Getenv("PATH")}
		} else {
			cmd = exec.CommandContext(ctx, "docker", "buildx", "build", "-t", "check", "-t", tag, "--label", "org.opencontainers.image.revision="+sha, gitURL)
			cmd.Env = []string{"PATH=" + os.Getenv("PATH")}
		}

		cmd.Dir = os.TempDir()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Find the working directory
	cmd := exec.CommandContext(ctx, "docker", "inspect", "--format", "{{.Config.WorkingDir}}", "check")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("getting working directory: %w\n%s", err, out)
	}

	var (
		secrets []servers.Secret
		env     []servers.Env
		command []string
	)
	for i := 0; i < len(args); i += 1 {
		if args[i] == "-e" {
			kv := args[i+1]
			parts := strings.SplitN(kv, "=", 2)

			if strings.HasSuffix(parts[0], "_TOKEN") || strings.HasSuffix(parts[0], "_KEY") || strings.HasSuffix(parts[0], "_PASSWORD") {
				secrets = append(secrets, servers.Secret{
					Name:    secretName(name, parts[0]),
					Env:     parts[0],
					Example: "<" + parts[0] + ">",
				})
			} else {
				env = append(env, servers.Env{
					Name:    parts[0],
					Example: parts[1],
				})
			}
			i += 1
		} else {
			command = append(command, args[i])
		}
	}

	icon, err := client.FindIcon(ctx, refProjectURL)
	if err != nil {
		return err
	}

	if branch == "main" {
		branch = ""
	}

	env, schema := servers.CreateSchema(name, env)

	server := servers.Server{
		Name:  name,
		Image: tag,
		Type:  "server",
		Meta: servers.Meta{
			Category: category,
			Tags:     tags,
		},
		About: servers.About{
			Icon:        icon,
			Title:       title,
			Description: "TODO (only to provide a better description than the upstream project)",
		},
		Source: servers.Source{
			Project:   projectURL,
			Upstream:  upstream,
			Branch:    branch,
			Directory: directory,
		},
		Run: servers.Run{
			Command: command,
		},
		Config: servers.Config{
			Description: "Configure the connection to TODO",
			Secrets:     secrets,
			Env:         env,
			Parameters:  schema,
		},
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

		fmt.Printf("\n-----------------------------------------\n\n")
	}
	if exists, err := checkLocalServerExists(name); err != nil {
		return err
	} else if exists {
		fmt.Printf("[WARNING] Server for %s already exists, overwriting...\n", name)
	}

	serverDir := filepath.Join("servers", server.Name)
	_ = os.Mkdir(serverDir, 0755)

	serverFile := filepath.Join(serverDir, "server.yaml")

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(server); err != nil {
		return err
	}

	if err := os.WriteFile(serverFile, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing server config: %w", err)
	}

	fmt.Printf("Server definition written to %s.\n", serverFile)

	// Default to mcp build instructions
	step2 := fmt.Sprintf(`
  2. Test out your server in Docker Desktop building the image, generating a catalog, and importing it:

     task build -- %[1]s
     task catalog -- %[1]s
     docker mcp catalog import $PWD/catalogs/%[1]s/catalog.yaml
`, name)

	if userProvidedImage != "" {
		step2 = fmt.Sprintf(`
  2. Test out your server in Docker Desktop by generating a catalog and importing it:

     task catalog -- %[1]s
     docker mcp catalog import $PWD/catalogs/%[1]s/catalog.yaml
`, name)
	}

	fmt.Printf(`
-----------------------------------------

What to do next?

  1. Review %[2]s and make sure no TODO remains.
%[3]s
  3. After doing so, you should be able to test it with the MCP Toolkit. Repeat step 2 as needed while making changes.

  4. Reset your catalog after testing:

     docker mcp catalog reset

  5. Open a Pull Request with the %[2]s file.
`, name, serverFile, step2)

	return nil
}

func guessName(projectURL string) string {
	parts := strings.Split(strings.ToLower(projectURL), "/")
	name := parts[len(parts)-1]

	name = strings.TrimPrefix(name, "mcp-server-")
	name = strings.TrimPrefix(name, "mcp-")
	name = strings.TrimPrefix(name, "server-")

	name = strings.TrimSuffix(name, "-mcp-server")
	name = strings.TrimSuffix(name, "-mcp")
	name = strings.TrimSuffix(name, "-server")

	return name
}

func secretName(server, name string) string {
	return server + "." + strings.TrimPrefix(strings.ToLower(name), strings.ToLower(server)+"_")
}

func checkLocalServerExists(name string) (bool, error) {
	entries, err := os.ReadDir("servers")
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == name {
				return true, nil
			}
		}
	}

	return false, nil
}

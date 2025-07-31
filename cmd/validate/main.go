package main

import (
	"context"
	"flag"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/docker/mcp-registry/internal/licenses"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

func main() {
	name := flag.String("name", "", "Name of the mcp server, name is guessed if not provided")
	flag.Parse()

	if err := run(*name); err != nil {
		log.Fatal(err)
	}
}

func run(name string) error {
	if err := isNameValid(name); err != nil {
		return err
	}

	if err := isDirectoryValid(name); err != nil {
		return err
	}

	if err := areSecretsValid(name); err != nil {
		return err
	}

	if err := isConfigEnvValid(name); err != nil {
		return err
	}

	if err := IsLicenseValid(name); err != nil {
		return err
	}
	if err := isIconValid(name); err != nil {
		return err
	}

	if err := isDockerfileValid(name); err != nil {
		return err
	}

	return nil
}

// check if the name is a valid
func isNameValid(name string) error {
	// check if name has only letters, numbers, and hyphens
	if !regexp.MustCompile(`^[a-z0-9-]+$`).MatchString(name) {
		return fmt.Errorf("name is not valid. It must be a lowercase string with only letters, numbers, and hyphens")
	}

	fmt.Println("✅ Name is valid")
	return nil
}

// check if the directory is valid
// servers/<NAME>/server.yaml exists
func isDirectoryValid(name string) error {
	_, err := os.Stat(filepath.Join("servers", name, "server.yaml"))
	if err != nil {
		return err
	}
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// check if the server.yaml file has a valid name
	if server.Name != name {
		return fmt.Errorf("server.yaml file has a invalid name. It must be %s", name)
	}

	fmt.Println("✅ Directory is valid")
	return nil
}

// check if the secrets are valid
// secrets must be prefixed with the name of the server
func areSecretsValid(name string) error {
	// read the server.yaml file
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// check if the server.yaml file has a valid secrets
	if len(server.Config.Secrets) > 0 {
		for _, secret := range server.Config.Secrets {
			if !strings.HasPrefix(secret.Name, name+".") {
				return fmt.Errorf("secret %s is not valid. It must be prefixed with the name of the server", secret.Name)
			}
		}
	}

	fmt.Println("✅ Secrets are valid")
	return nil
}

// Check parameter usage is valid
func isConfigEnvValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	for _, e := range server.Config.Env {
		if !strings.HasPrefix(e.Value, "{{") {
			continue
		}
		if !strings.HasPrefix(e.Value, "{{"+server.Name+".") {
			return fmt.Errorf("server uses unknown parameter %q: %q", server.Name, e.Value)
		}
	}

	fmt.Println("✅ Config env is valid")
	return nil
}

// check if the license is valid
// the license must be valid
func IsLicenseValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// Skip validation for servers without source project
	if server.Source.Project == "" {
		fmt.Println("✅ License validation skipped (no source project)")
		return nil
	}

	ctx := context.Background()
	client := github.New()
	repository, err := client.GetProjectRepository(ctx, server.Source.Project)
	if err != nil {
		return err
	}

	if !licenses.IsValid(repository.License) {
		return fmt.Errorf("project %s is licensed under %s which may be incompatible with some tools", server.Source.Project, repository.License.GetName())
	}
	fmt.Println("✅ License is valid")

	return nil
}

func isIconValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	if server.About.Icon == "" {
		fmt.Println("🛑 No icon found")
		return nil
	}
	// fetch the image and check the size
	resp, err := http.Get(server.About.Icon)
	if err != nil {
		fmt.Println("🛑 Icon could not be fetched")
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("🛑 Icon could not be fetched, status code: %d, url: %s\n", resp.StatusCode, server.About.Icon)
		return nil
	}
	if resp.ContentLength > 2*1024*1024 {
		fmt.Println("🛑 Icon is too large. It must be less than 2MB")
		return nil
	}
	img, format, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return err
	}
	if format != "png" {
		fmt.Println("🛑 Icon is not a png. It must be a png")
		return nil
	}

	if img.Width > 512 || img.Height > 512 {
		fmt.Println("🛑 Icon is too large. It must be less than 512x512")
		return nil
	}

	fmt.Println("✅ Icon is valid")
	return nil
}

func isDockerfileValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// Skip validation for servers without source project
	if server.Source.Project == "" {
		fmt.Println("✅ Dockerfile validation skipped (no source project)")
		return nil
	}

	// Determine Dockerfile path - default to "Dockerfile" if not specified
	dockerfilePath := server.Source.Dockerfile
	if dockerfilePath == "" {
		dockerfilePath = "Dockerfile"
	}

	// Validate Dockerfile path format
	if err := validateDockerfilePath(dockerfilePath); err != nil {
		return fmt.Errorf("invalid Dockerfile path: %w", err)
	}

	// Skip GitHub API validation for local development or if GITHUB_TOKEN is not set
	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Println("🛑 Dockerfile content validation skipped (no GITHUB_TOKEN)")
		return nil
	}

	// Fetch and validate Dockerfile content from GitHub
	ctx := context.Background()
	client := github.NewFromServer(server)

	// Construct full path including directory if specified
	fullPath := dockerfilePath
	if server.Source.Directory != "" {
		fullPath = filepath.Join(server.Source.Directory, dockerfilePath)
	}

	content, err := client.GetFileContent(ctx, server.Source.Project, server.Source.Branch, fullPath)
	if err != nil {
		return fmt.Errorf("failed to fetch Dockerfile: %w", err)
	}

	// Validate Dockerfile content
	if err := validateDockerfileContent(content); err != nil {
		return fmt.Errorf("invalid Dockerfile content: %w", err)
	}

	fmt.Println("✅ Dockerfile is valid")
	return nil
}

func validateDockerfilePath(path string) error {
	// Must be relative path
	if strings.HasPrefix(path, "/") {
		return fmt.Errorf("Dockerfile path must be relative, not absolute")
	}

	// Should not contain directory traversal
	if strings.Contains(path, "..") {
		return fmt.Errorf("Dockerfile path should not contain directory traversal (..) components")
	}

	// Should be a valid filename
	if strings.Contains(path, "\x00") {
		return fmt.Errorf("Dockerfile path contains invalid characters")
	}

	return nil
}

func validateDockerfileContent(content string) error {
	lines := strings.Split(content, "\n")
	
	// Remove empty lines and comments for validation
	var validLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
			validLines = append(validLines, trimmed)
		}
	}

	if len(validLines) == 0 {
		return fmt.Errorf("Dockerfile is empty or contains only comments")
	}

	// First instruction must be FROM
	firstLine := strings.ToUpper(strings.TrimSpace(validLines[0]))
	if !strings.HasPrefix(firstLine, "FROM ") {
		return fmt.Errorf("Dockerfile must start with FROM instruction")
	}

	// Check for basic security issues
	for _, line := range validLines {
		upperLine := strings.ToUpper(line)
		
		// Check for potential security issues
		if strings.Contains(upperLine, "PASSWORD=") || 
		   strings.Contains(upperLine, "SECRET=") ||
		   strings.Contains(upperLine, "API_KEY=") ||
		   strings.Contains(upperLine, "TOKEN=") {
			return fmt.Errorf("Dockerfile should not contain hardcoded credentials")
		}
	}

	// Should have an ENTRYPOINT or CMD instruction
	hasEntrypoint := false
	hasCmd := false
	for _, line := range validLines {
		upperLine := strings.ToUpper(strings.TrimSpace(line))
		if strings.HasPrefix(upperLine, "ENTRYPOINT ") {
			hasEntrypoint = true
		}
		if strings.HasPrefix(upperLine, "CMD ") {
			hasCmd = true
		}
	}

	if !hasEntrypoint && !hasCmd {
		return fmt.Errorf("Dockerfile must contain either ENTRYPOINT or CMD instruction")
	}

	return nil
}

func readServerYaml(name string) (servers.Server, error) {
	serverYaml, err := os.ReadFile(filepath.Join("servers", name, "server.yaml"))
	if err != nil {
		return servers.Server{}, err
	}
	var server servers.Server
	err = yaml.Unmarshal(serverYaml, &server)
	if err != nil {
		return servers.Server{}, err
	}
	return server, nil
}

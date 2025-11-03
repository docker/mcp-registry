package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/docker/mcp-registry/internal/licenses"
	"github.com/docker/mcp-registry/internal/mcp"
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

	if err := isTitleValid(name); err != nil {
		return err
	}

	if err := isYamlIndentationValid(name); err != nil {
		return err
	}

	if err := isCommitPinnedIfNecessary(name); err != nil {
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
	if err := isRemoteValid(name); err != nil {
		return err
	}

	if err := isOAuthDynamicValid(name); err != nil {
		return err
	}

	if err := isPociValid(name); err != nil {
		return err
	}

	return nil
}

// legacyNameExceptions enumerates catalog entries added before current naming rules.
var legacyNameExceptions = map[string]bool{
	"SQLite":              true,
	"osp_marketing_tools": true,
	"youtube_transcript":  true,
}

// check if the name is a valid
func isNameValid(name string) error {
	// check if name has only letters, numbers, and hyphens
	if !regexp.MustCompile(`^[a-z0-9-]+$`).MatchString(name) {
		if legacyNameExceptions[name] {
			fmt.Printf("⚠️ Name %s is grandfathered and bypasses naming rules.\n", name)
			return nil
		}
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

// check if the title is valid
// titles should not contain "MCP" or "Server" and every word should be capitalized
func isTitleValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	title := server.About.Title

	// Check for "MCP" or "Server" in the title
	if strings.Contains(title, "MCP") {
		return fmt.Errorf("title should not contain 'MCP': %s", title)
	}
	if strings.Contains(title, "Server") {
		return fmt.Errorf("title should not contain 'Server': %s", title)
	}

	// Check that every word is capitalized
	words := strings.Fields(title)
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		// Check if the first character is uppercase
		firstChar := []rune(word)[0]
		if string(firstChar) != strings.ToUpper(string(firstChar)) {
			return fmt.Errorf("title must have every word capitalized: %s (word: %s)", title, word)
		}
	}

	fmt.Println("✅ Title is valid")
	return nil
}

// check if the YAML file is formatted correctly using prettier
func isYamlIndentationValid(name string) error {
	yamlPath := filepath.Join("servers", name, "server.yaml")

	// Use npx to run prettier without requiring local installation
	cmd := exec.Command("npx", "--yes", "prettier", "--check", yamlPath)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("YAML file is not formatted correctly. Run 'npx prettier --write %s' to fix:\n%s", yamlPath, string(output))
	}

	fmt.Println("✅ YAML formatting is valid")
	return nil
}

var commitSHA1Pattern = regexp.MustCompile(`^[a-f0-9]{40}$`)

// isCommitPinnedIfNecessary ensures that every local server is pinned to a specific commit.
func isCommitPinnedIfNecessary(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	if server.Type != "server" {
		fmt.Println("✅ Commit pin not required (non-local server)")
		return nil
	}

	if server.Source.Commit == "" {
		return fmt.Errorf("local server must specify source.commit to pin the audited revision")
	}

	if !commitSHA1Pattern.MatchString(strings.ToLower(server.Source.Commit)) {
		return fmt.Errorf("source.commit must be a 40-character lowercase SHA1 (got %q)", server.Source.Commit)
	}

	fmt.Println("✅ Commit is pinned")
	return nil
}

// secretNamePattern validates that secret names match the expected prefix.name
// format requirement.
var secretNamePattern = regexp.MustCompile(`^[A-Za-z0-9_-]+\.[A-Za-z0-9._-]+$`)

// legacySecretNameExceptions enumerates secrets defined before the current
// naming rules were introduced.
var legacySecretNameExceptions = map[string]map[string]bool{
	"nasdaq-data-link": {
		"nasdaq_data_link_api_key": true,
	},
	"sec-edgar": {
		"sec_edgar_user_agent": true,
	},
}

// check if the secrets are valid
func areSecretsValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// Ensure that all secrets match the expected format. We no longer require
	// that the prefix matches the server name.
	for _, secret := range server.Config.Secrets {
		if !secretNamePattern.MatchString(secret.Name) {
			if legacySecretNameExceptions[name][secret.Name] {
				fmt.Printf("⚠️ Secret %s for %s is grandfathered and bypasses naming rules.\n", secret.Name, name)
				continue
			}
			return fmt.Errorf("secret %s is not valid. It must use prefix.name format with alphanumeric characters, hyphen, period, or underscore", secret.Name)
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
	ctx := context.Background()
	client := github.New()
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// Skip license validation for remote servers without source
	if server.Source.Project == "" {
		fmt.Println("✅ License validation skipped (remote server)")
		return nil
	}

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
		fmt.Println("⚠️ No icon found")
		return nil
	}
	// fetch the image and check the size
	resp, err := http.Get(server.About.Icon)
	if err != nil {
		fmt.Println("⚠️ Icon could not be fetched")
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("⚠️ Icon could not be fetched, status code: %d, url: %s\n", resp.StatusCode, server.About.Icon)
		return nil
	}
	if resp.ContentLength > 2*1024*1024 {
		fmt.Println("⚠️ Icon is too large. It must be less than 2MB")
		return nil
	}

	// Check content type for SVG, favicon, and WebP support
	contentType := resp.Header.Get("Content-Type")
	switch contentType {
	case "image/svg+xml":
		fmt.Println("✅ Icon is valid (SVG)")
		return nil
	case "image/x-icon":
		fmt.Println("✅ Icon is valid (favicon)")
		return nil
	case "image/webp":
		fmt.Println("✅ Icon is valid (WebP)")
		return nil
	}

	img, format, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return err
	}
	if format != "png" && format != "jpeg" {
		fmt.Println("⚠️ Icon is not a png or svg. It must be a png or svg")
		return nil
	}

	if img.Width > 512 || img.Height > 512 {
		fmt.Println("⚠️ Icon is too large. It must be less than 512x512")
		return nil
	}

	fmt.Println("✅ Icon is valid")
	return nil
}

// check if the remote configuration is valid
func isRemoteValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// Skip validation for non-remote servers
	if server.Remote.URL == "" {
		fmt.Println("✅ Remote validation skipped (not a remote server)")
		return nil
	}

	// Check that transport_type is not empty for remote servers
	if server.Remote.TransportType == "" {
		return fmt.Errorf("remote server must have a transport_type specified")
	}

	// Validate transport_type is one of the allowed values
	validTransports := []string{"stdio", "sse", "streamable-http"}
	isValid := false
	for _, valid := range validTransports {
		if server.Remote.TransportType == valid {
			isValid = true
			break
		}
	}
	if !isValid {
		return fmt.Errorf("remote server transport_type must be one of: stdio, sse, streamable-http (got: %s)", server.Remote.TransportType)
	}

	if err := hasValidTools(server); err != nil {
		return err
	}

	fmt.Println("✅ Remote is valid")
	return nil
}

// Check that there is either a tools.json, dynamic tools, or can fetch remote tools
func hasValidTools(server servers.Server) error {
	defaultErr := fmt.Errorf("server must have either a tools.json, dynamic tools, or can fetch remote tools")

	// Dynamic tools are valid
	if server.Dynamic != nil && server.Dynamic.Tools {
		fmt.Println("✅ Dynamic tools are valid")
		return nil
	}

	// Tools.json is valid
	tools, err := readToolsJson(server.Name)
	if err == nil {
		toolCount := len(tools)
		fmt.Printf("✅ tools.json is valid. Found %d tools.\n", toolCount)
		return nil
	}
	if !os.IsNotExist(err) {
		fmt.Printf("🛑 Tools.json could not be read: %v\n", err)
		return defaultErr
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Remote tools are valid
	remoteTools, err := mcp.RemoteTools(ctx, server)
	if err != nil {
		fmt.Printf("🛑 Remote tools could not be fetched (if auth is required, specify \ndynamic:\n  tools: true\n): %v\n", err)
		return defaultErr
	}

	toolCount := len(remoteTools)

	fmt.Printf("✅ Remote tools are valid. Found %d tools.\n", toolCount)
	return nil
}

// Some special entries bypass the dynamic tools requirement.
var oauthDynamicToolExceptions = map[string]bool{
	"github-official": true,
}

// check if servers with OAuth have dynamic tools enabled
func isOAuthDynamicValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// If server has OAuth configuration, it must have dynamic tools enabled
	if len(server.OAuth) > 0 {
		if server.Dynamic == nil || !server.Dynamic.Tools {
			if oauthDynamicToolExceptions[name] {
				fmt.Printf("⚠️ OAuth dynamic rule bypassed for %s (special configuration).\n", name)
			} else {
				return fmt.Errorf("server with OAuth must have 'dynamic: tools: true' configuration")
			}
		}
	}

	fmt.Println("✅ OAuth dynamic configuration is valid")
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

func readToolsJson(name string) ([]mcp.Tool, error) {
	path := filepath.Join("servers", name, "tools.json")
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tools []mcp.Tool
	if err := json.Unmarshal(buf, &tools); err != nil {
		return nil, err
	}

	return tools, nil
}

func isPociValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	if server.Type != "poci" {
		return nil
	}

	for _, tool := range server.Tools {
		if tool.Container.Image != "" {
			if err := pullPociImage(tool.Container.Image); err != nil {
				fmt.Printf("🛑 Could not pull poci image %s: %v\n", tool.Container.Image, err)
				return err
			}
		}
	}

	fmt.Println("✅ Poci image is valid")
	return nil
}

func pullPociImage(image string) error {
	cmd := exec.Command("docker", "pull", image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

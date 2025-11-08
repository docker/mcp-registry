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
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/docker/mcp-registry/pkg/servers"
)

// writeJSONFile stores the provided value as indented JSON at the given path.
func writeJSONFile(path string, value any) error {
	payload, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, payload, 0o644)
}

// readJSONFile populates value with JSON data read from the provided path.
func readJSONFile(path string, value any) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, value)
}

// removeIfPresent deletes the file at the path when it exists.
func removeIfPresent(path string) {
	if path == "" {
		return
	}
	if _, err := os.Stat(path); err == nil {
		_ = os.Remove(path)
	}
}

// loadServerYAMLFromWorkspace loads a server YAML file located in the workspace.
func loadServerYAMLFromWorkspace(workspace, relative string) (servers.Server, error) {
	fullPath := filepath.Join(workspace, relative)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return servers.Server{}, err
	}
	return decodeServerDocument(content)
}

// loadServerYAMLAt loads a server YAML file from the git history at the commit.
func loadServerYAMLAt(workspace, commit, relative string) (servers.Server, error) {
	out, err := runGitCommand(workspace, "show", fmt.Sprintf("%s:%s", commit, relative))
	if err != nil {
		return servers.Server{}, err
	}
	return decodeServerDocument([]byte(out))
}

// decodeServerDocument converts raw YAML bytes into a servers.Server.
func decodeServerDocument(raw []byte) (servers.Server, error) {
	var doc servers.Server
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return servers.Server{}, err
	}
	return doc, nil
}

// isLocalServer returns true when the definition corresponds to a server that
// should be security reviewed. This includes both local servers (mcp/ namespace)
// and external servers that have a source repository.
func isLocalServer(doc servers.Server) bool {
	if !strings.EqualFold(doc.Type, "server") {
		return false
	}

	// Include servers built in the mcp/ namespace.
	if strings.HasPrefix(strings.TrimSpace(doc.Image), "mcp/") {
		return true
	}

	// Include external servers that have a source repository with a commit pin.
	// We can't guarantee provenance between source and image, but reviewing the
	// source is better than nothing.
	project := strings.TrimSpace(doc.Source.Project)
	commit := strings.TrimSpace(doc.Source.Commit)
	return project != "" && commit != ""
}

// gitDiff runs git diff for server YAML files and returns the resulting paths.
func gitDiff(workspace, base, head, mode string) ([]string, error) {
	args := []string{"diff", mode, base, head, "--", "servers/*/server.yaml"}
	out, err := runGitCommand(workspace, args...)
	if err != nil {
		return nil, err
	}

	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// runGitCommand executes git with the given arguments inside the directory.
func runGitCommand(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, string(output))
	}
	return string(output), nil
}

// splitList normalizes a delimited string into lowercase server names.
func splitList(raw string) []string {
	if raw == "" {
		return nil
	}
	var values []string
	for _, segment := range strings.FieldsFunc(raw, func(r rune) bool {
		return r == ',' || r == '\n' || r == ' ' || r == '\t'
	}) {
		value := strings.TrimSpace(segment)
		if value != "" {
			values = append(values, strings.ToLower(value))
		}
	}
	return values
}

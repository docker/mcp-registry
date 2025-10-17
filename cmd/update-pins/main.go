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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
)

// main orchestrates the pin refresh process, updating server definitions when
// upstream branches advance.
func main() {
	ctx := context.Background()

	// Enumerate the server directories that contain YAML definitions.
	entries, err := os.ReadDir("servers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading servers directory: %v\n", err)
		os.Exit(1)
	}

	var updated []string
	for _, entry := range entries {
		// Ignore any files that are not server directories.
		if !entry.IsDir() {
			continue
		}

		serverPath := filepath.Join("servers", entry.Name(), "server.yaml")
		server, err := servers.Read(serverPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading %s: %v\n", serverPath, err)
			continue
		}

		if server.Type != "server" {
			continue
		}

		if !strings.HasPrefix(server.Image, "mcp/") {
			continue
		}

		if server.Source.Project == "" {
			continue
		}

		// Only GitHub repositories are supported by the current workflow.
		if !strings.Contains(server.Source.Project, "github.com/") {
			fmt.Printf("Skipping %s: project is not hosted on GitHub.\n", server.Name)
			continue
		}

		// Unpinned servers have to undergo a separate security audit first.
		existing := strings.ToLower(server.Source.Commit)
		if existing == "" {
			fmt.Printf("Skipping %s: no pinned commit present.\n", server.Name)
			continue
		}

		// Resolve the current branch head for comparison.
		branch := server.GetBranch()
		client := github.NewFromServer(server)

		latest, err := client.GetCommitSHA1(ctx, server.Source.Project, branch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetching commit for %s: %v\n", server.Name, err)
			continue
		}

		latest = strings.ToLower(latest)

		changed, err := writeCommit(serverPath, latest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "updating %s: %v\n", server.Name, err)
			continue
		}

		if existing != latest {
			fmt.Printf("Updated %s: %s -> %s\n", server.Name, existing, latest)
		} else if changed {
			fmt.Printf("Reformatted pinned commit for %s at %s\n", server.Name, latest)
		}

		if changed {
			updated = append(updated, server.Name)
		}
		if existing == latest && !changed {
			continue
		}
	}

	if len(updated) == 0 {
		fmt.Println("No commit updates required.")
		return
	}

	sort.Strings(updated)
	fmt.Println("Servers with updated pins:", strings.Join(updated, ", "))
}

// writeCommit inserts or updates the commit field inside the source block of
// a server definition while preserving the surrounding formatting. The bool
// return value indicates whether the file contents were modified.
func writeCommit(path string, updated string) (bool, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	lines := strings.Split(string(content), "\n")
	sourceIndex := -1
	for i, line := range lines {
		if strings.HasPrefix(line, "source:") {
			sourceIndex = i
			break
		}
	}
	if sourceIndex == -1 {
		return false, fmt.Errorf("no source block found")
	}

	commitIndex := -1
	indent := ""
	commitPattern := regexp.MustCompile(`^([ \t]+)commit:\s*[a-fA-F0-9]{40}\s*$`)
	for i := sourceIndex + 1; i < len(lines); i++ {
		line := lines[i]
		if !strings.HasPrefix(line, "  ") {
			break
		}

		if match := commitPattern.FindStringSubmatch(line); match != nil {
			commitIndex = i
			indent = match[1]
			break
		}
	}

	if commitIndex < 0 {
		return false, fmt.Errorf("no commit line found in source block")
	}

	newLine := indent + "commit: " + updated
	lines[commitIndex] = newLine

	output := strings.Join(lines, "\n")
	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}

	if output == string(content) {
		return false, nil
	}

	if err := os.WriteFile(path, []byte(output), 0o644); err != nil {
		return false, err
	}
	return true, nil
}

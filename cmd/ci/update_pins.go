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
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
)

// runUpdatePins refreshes pinned commits for local servers by resolving the
// latest upstream revision on the tracked branch and updating server YAML
// definitions in place. Accepts a --servers flag to limit which servers to
// update. Emits a summary of modified servers on stdout; errors are reported
// per server so that a single failure does not abort the entire sweep.
func runUpdatePins(args []string) error {
	fs := flag.NewFlagSet("update-pins", flag.ExitOnError)
	serversFlag := fs.String("servers", "", "Comma-separated list of servers to update (leave blank for all)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	ctx := context.Background()

	// Parse the servers filter if provided.
	allowedServers := make(map[string]bool)
	if *serversFlag != "" {
		for _, name := range splitList(*serversFlag) {
			allowedServers[name] = true
		}
	}

	entries, err := os.ReadDir("servers")
	if err != nil {
		return fmt.Errorf("reading servers directory: %w", err)
	}

	var updated []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// Skip servers not in the filter list (if provided).
		serverName := strings.ToLower(entry.Name())
		if len(allowedServers) > 0 && !allowedServers[serverName] {
			continue
		}

		serverPath := filepath.Join("servers", entry.Name(), "server.yaml")

		// Parse the server definition so that we can evaluate eligibility and
		// discover the backing GitHub repository and branch information.
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

		// Only repositories hosted on GitHub can be advanced by this command,
		// because the helper client relies on the GitHub API for commit lookup.
		if !strings.Contains(server.Source.Project, "github.com/") {
			fmt.Printf("Skipping %s: project is not hosted on GitHub.\n", server.Name)
			continue
		}

		existing := strings.ToLower(server.Source.Commit)
		if existing == "" {
			fmt.Printf("Skipping %s: no pinned commit present.\n", server.Name)
			continue
		}

		client := github.NewFromServer(server)
		// Resolve the latest commit on the configured branch so we can refresh
		// the pin if it has advanced since the last sweep.
		latest, err := client.GetCommitSHA1(ctx, server.Source.Project, server.GetBranch())
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetching commit for %s: %v\n", server.Name, err)
			continue
		}
		latest = strings.ToLower(latest)

		changed, err := writePinnedCommit(serverPath, latest)
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
	}

	if len(updated) == 0 {
		fmt.Println("No commit updates required.")
		return nil
	}

	sort.Strings(updated)
	fmt.Println("Servers with updated pins:", strings.Join(updated, ", "))
	return nil
}

// writePinnedCommit replaces the commit field inside the source block with the
// provided SHA while preserving formatting. A boolean indicates whether the
// file changed.
func writePinnedCommit(path string, updated string) (bool, error) {
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
		return false, fmt.Errorf("no source block found in %s", path)
	}

	// Scan the nested source block until we locate the commit attribute,
	// capturing its indentation so that formatting survives the rewrite.
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
		return false, fmt.Errorf("no commit line found in %s", path)
	}

	// Replace only the commit value so that other keys maintain their
	// original ordering and indentation.
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

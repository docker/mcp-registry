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
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// runPrepareNewServers checks out repositories for newly added local servers,
// given a JSON context file. It expects --context-file and --output-dir flags
// and prepares per-server metadata and source trees.
func runPrepareNewServers(args []string) error {
	flags := flag.NewFlagSet("prepare-new-servers", flag.ContinueOnError)
	contextFile := flags.String("context-file", "", "path to JSON context file")
	outputDir := flags.String("output-dir", "", "directory to receive prepared artifacts")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *contextFile == "" || *outputDir == "" {
		return errors.New("context-file and output-dir are required")
	}

	var targets []newServerTarget
	if err := readJSONFile(*contextFile, &targets); err != nil {
		return err
	}

	if len(targets) == 0 {
		return nil
	}

	if err := os.MkdirAll(*outputDir, 0o755); err != nil {
		return err
	}

	for _, target := range targets {
		if err := prepareNewServerTarget(*outputDir, target); err != nil {
			return fmt.Errorf("prepare new server %s: %w", target.Server, err)
		}
	}

	return nil
}

// prepareNewServerTarget clones the upstream repository at the pinned commit
// for a new server and records metadata for downstream review.
func prepareNewServerTarget(outputDir string, target newServerTarget) error {
	serverDir := filepath.Join(outputDir, target.Server)
	repoDir := filepath.Join(serverDir, "repo")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		return err
	}

	if err := initGitRepository(repoDir, target.Project); err != nil {
		return err
	}
	if err := fetchCommit(repoDir, target.Commit); err != nil {
		return err
	}
	if _, err := runGitCommand(repoDir, "checkout", target.Commit); err != nil {
		return err
	}

	metadata := map[string]string{
		"server":     target.Server,
		"repository": target.Project,
		"commit":     target.Commit,
		"directory":  target.Directory,
	}
	if err := writeJSONFile(filepath.Join(serverDir, "metadata.json"), metadata); err != nil {
		return err
	}

	summary := buildNewServerDetail(target)
	return os.WriteFile(filepath.Join(serverDir, "README.md"), []byte(summary), 0o644)
}

// buildNewServerDetail returns a Markdown overview describing the cloned
// server, suitable for inclusion in review prompts.
func buildNewServerDetail(target newServerTarget) string {
	builder := strings.Builder{}
	builder.WriteString("# New Server Security Review\n\n")
	builder.WriteString(fmt.Sprintf("- Server: %s\n", target.Server))
	builder.WriteString(fmt.Sprintf("- Repository: %s\n", target.Project))
	builder.WriteString(fmt.Sprintf("- Commit: %s\n", target.Commit))
	if target.Directory != "" {
		builder.WriteString(fmt.Sprintf("- Directory: %s\n", target.Directory))
	} else {
		builder.WriteString("- Directory: (repository root)\n")
	}
	return builder.String()
}

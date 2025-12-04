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

// runCollectUpdatedPins gathers metadata for servers that updated their commit
// pins between two git revisions. It expects --base, --head, --workspace,
// --output-json, and --summary-md arguments. The identified targets are written
// to the JSON file while a Markdown summary is produced for humans.
func runCollectUpdatedPins(args []string) error {
	flags := flag.NewFlagSet("collect-updated-pins", flag.ContinueOnError)
	base := flags.String("base", "", "base git commit SHA")
	head := flags.String("head", "", "head git commit SHA")
	workspace := flags.String("workspace", ".", "path to repository workspace")
	outputJSON := flags.String("output-json", "", "path to write JSON context")
	summaryMD := flags.String("summary-md", "", "path to write Markdown summary")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *base == "" || *head == "" || *outputJSON == "" || *summaryMD == "" {
		return errors.New("base, head, output-json, and summary-md are required")
	}

	targets, err := collectUpdatedPinTargets(*workspace, *base, *head)
	if err != nil {
		return err
	}

	if len(targets) == 0 {
		removeIfPresent(*outputJSON)
		removeIfPresent(*summaryMD)
		return nil
	}

	if err := writeJSONFile(*outputJSON, targets); err != nil {
		return err
	}

	summary := buildPinSummary(targets)
	return os.WriteFile(*summaryMD, []byte(summary), 0o644)
}

// collectUpdatedPinTargets identifies local servers whose pinned commits differ
// between the supplied git revisions and returns their metadata for further
// processing.
func collectUpdatedPinTargets(workspace, base, head string) ([]pinTarget, error) {
	paths, err := gitDiff(workspace, base, head, "--name-only")
	if err != nil {
		return nil, err
	}

	var targets []pinTarget
	for _, relative := range paths {
		if !strings.HasPrefix(relative, "servers/") || !strings.HasSuffix(relative, "server.yaml") {
			continue
		}

		baseDoc, err := loadServerYAMLAt(workspace, base, relative)
		if err != nil {
			continue
		}
		headDoc, err := loadServerYAMLFromWorkspace(workspace, relative)
		if err != nil {
			continue
		}

		if !isLocalServer(headDoc) || !isLocalServer(baseDoc) {
			continue
		}

		oldCommit := strings.TrimSpace(baseDoc.Source.Commit)
		newCommit := strings.TrimSpace(headDoc.Source.Commit)
		project := strings.TrimSpace(headDoc.Source.Project)
		if oldCommit == "" || newCommit == "" || oldCommit == newCommit || project == "" {
			continue
		}

		targets = append(targets, pinTarget{
			Server:    filepath.Base(filepath.Dir(relative)),
			File:      relative,
			Image:     strings.TrimSpace(headDoc.Image),
			Project:   project,
			Directory: strings.TrimSpace(headDoc.Source.Directory),
			OldCommit: oldCommit,
			NewCommit: newCommit,
		})
	}

	return targets, nil
}

// buildPinSummary renders a Markdown section describing updated pin targets so
// that review tooling and humans can understand what changed.
func buildPinSummary(targets []pinTarget) string {
	builder := strings.Builder{}
	builder.WriteString("## Updated Commit Pins\n\n")

	for _, target := range targets {
		builder.WriteString(fmt.Sprintf("### %s\n", target.Server))
		builder.WriteString(fmt.Sprintf("- Repository: %s\n", target.Project))
		if target.Directory != "" {
			builder.WriteString(fmt.Sprintf("- Directory: %s\n", target.Directory))
		} else {
			builder.WriteString("- Directory: (repository root)\n")
		}
		builder.WriteString(fmt.Sprintf("- Previous commit: `%s`\n", target.OldCommit))
		builder.WriteString(fmt.Sprintf("- New commit: `%s`\n", target.NewCommit))
		builder.WriteString(fmt.Sprintf("- Diff path: /tmp/security-review/pins/%s/diff.patch\n\n", target.Server))
	}

	return builder.String()
}

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
	"io/fs"
	"path/filepath"
	"strings"
)

// runCollectFullAudit enumerates local servers (optionally filtered) and writes
// their metadata to a JSON file for manual auditing. It expects --workspace,
// --servers, and --output-json flags.
func runCollectFullAudit(args []string) error {
	flags := flag.NewFlagSet("collect-full-audit", flag.ContinueOnError)
	workspace := flags.String("workspace", ".", "path to repository workspace")
	filter := flags.String("servers", "", "optional comma-separated server filter")
	outputJSON := flags.String("output-json", "", "path to write JSON context")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *outputJSON == "" {
		return errors.New("output-json is required")
	}

	targets, err := collectAuditTargets(*workspace, *filter)
	if err != nil {
		return err
	}

	if len(targets) == 0 {
		removeIfPresent(*outputJSON)
		return nil
	}

	return writeJSONFile(*outputJSON, targets)
}

// collectAuditTargets returns audit targets for all local servers or a filtered
// subset based on the supplied comma-separated list.
func collectAuditTargets(workspace, filter string) ([]auditTarget, error) {
	filterSet := make(map[string]struct{})
	for _, name := range splitList(filter) {
		filterSet[name] = struct{}{}
	}

	var targets []auditTarget
	err := filepath.WalkDir(filepath.Join(workspace, "servers"), func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, "server.yaml") {
			return nil
		}

		relative := strings.TrimPrefix(path, workspace+string(filepath.Separator))
		doc, err := loadServerYAMLFromWorkspace(workspace, relative)
		if err != nil || !isLocalServer(doc) {
			return nil
		}

		server := filepath.Base(filepath.Dir(path))
		if len(filterSet) > 0 {
			if _, ok := filterSet[strings.ToLower(server)]; !ok {
				return nil
			}
		}

		project := strings.TrimSpace(doc.Source.Project)
		commit := strings.TrimSpace(doc.Source.Commit)
		if project == "" || commit == "" {
			return nil
		}

		targets = append(targets, auditTarget{
			Server:    server,
			Project:   project,
			Commit:    commit,
			Directory: strings.TrimSpace(doc.Source.Directory),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return targets, nil
}

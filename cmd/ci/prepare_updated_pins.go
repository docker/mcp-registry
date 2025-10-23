package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// runPrepareUpdatedPins fetches upstream repositories and prepares diff
// artifacts for each updated pin listed in the context file. It consumes
// --context-file and --output-dir flags and writes diffs, logs, and metadata
// for downstream analysis.
func runPrepareUpdatedPins(args []string) error {
	flags := flag.NewFlagSet("prepare-updated-pins", flag.ContinueOnError)
	contextFile := flags.String("context-file", "", "path to JSON context file")
	outputDir := flags.String("output-dir", "", "directory to receive prepared artifacts")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *contextFile == "" || *outputDir == "" {
		return errors.New("context-file and output-dir are required")
	}

	var targets []pinTarget
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
		if err := preparePinTarget(*outputDir, target); err != nil {
			return fmt.Errorf("prepare pin target %s: %w", target.Server, err)
		}
	}

	return nil
}

// preparePinTarget materializes git diffs, commit logs, and metadata for a
// single commit pin update, storing the results under the provided output
// directory.
func preparePinTarget(outputDir string, target pinTarget) error {
	serverDir := filepath.Join(outputDir, target.Server)
	repoDir := filepath.Join(serverDir, "repo")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		return err
	}

	if err := initGitRepository(repoDir, target.Project); err != nil {
		return err
	}

	for _, commit := range []string{target.OldCommit, target.NewCommit} {
		if err := fetchCommit(repoDir, commit); err != nil {
			return err
		}
	}

	diffArgs := []string{"diff", target.OldCommit, target.NewCommit}
	if target.Directory != "" && target.Directory != "." {
		diffArgs = append(diffArgs, "--", target.Directory)
	}
	diffOut, err := runGitCommand(repoDir, diffArgs...)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(serverDir, "diff.patch"), []byte(diffOut), 0o644); err != nil {
		return err
	}

	logOut, err := runGitCommand(repoDir, "log", "--oneline", "--stat", fmt.Sprintf("%s..%s", target.OldCommit, target.NewCommit))
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(serverDir, "changes.log"), []byte(logOut), 0o644); err != nil {
		return err
	}

	metadata := map[string]string{
		"server":     target.Server,
		"repository": target.Project,
		"old_commit": target.OldCommit,
		"new_commit": target.NewCommit,
		"directory":  target.Directory,
	}
	return writeJSONFile(filepath.Join(serverDir, "metadata.json"), metadata)
}

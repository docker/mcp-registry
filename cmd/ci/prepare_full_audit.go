package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// runPrepareFullAudit clones source data for a single audit target specified by
// a JSON descriptor. It requires --target-file and --output-dir flags and
// prepares the repository checkout plus metadata.
func runPrepareFullAudit(args []string) error {
	flags := flag.NewFlagSet("prepare-full-audit", flag.ContinueOnError)
	targetFile := flags.String("target-file", "", "path to JSON target descriptor")
	outputDir := flags.String("output-dir", "", "directory to receive prepared artifacts")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *targetFile == "" || *outputDir == "" {
		return errors.New("target-file and output-dir are required")
	}

	var target auditTarget
	if err := readJSONFile(*targetFile, &target); err != nil {
		return err
	}

	return prepareAuditTarget(*outputDir, target)
}

// prepareAuditTarget materializes repository state and metadata for auditing a
// single server, storing artifacts beneath the provided output directory.
func prepareAuditTarget(outputDir string, target auditTarget) error {
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

	context := buildAuditContext(target, repoDir)
	if err := os.WriteFile(filepath.Join(serverDir, "context.md"), []byte(context), 0o644); err != nil {
		return err
	}

	return writeJSONFile(filepath.Join(serverDir, "metadata.json"), target)
}

// buildAuditContext produces Markdown describing the prepared audit checkout,
// which is used to prime review prompts.
func buildAuditContext(target auditTarget, repoDir string) string {
	builder := strings.Builder{}
	builder.WriteString("# Full Audit Target\n\n")
	builder.WriteString(fmt.Sprintf("- Server: %s\n", target.Server))
	builder.WriteString(fmt.Sprintf("- Repository: %s\n", target.Project))
	builder.WriteString(fmt.Sprintf("- Commit: %s\n", target.Commit))
	if target.Directory != "" {
		builder.WriteString(fmt.Sprintf("- Directory: %s\n", target.Directory))
	} else {
		builder.WriteString("- Directory: (repository root)\n")
	}
	builder.WriteString(fmt.Sprintf("- Checkout path: %s\n", repoDir))
	return builder.String()
}

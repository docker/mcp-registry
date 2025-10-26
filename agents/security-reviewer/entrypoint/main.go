package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	reportTemplatePath        = "/opt/security-reviewer/report-template.md"
	defaultPromptPath         = "/input/prompt.md"
	defaultRepositoryPath     = "/input/repository"
	defaultReportPath         = "/output/report.md"
	defaultLabelsPath         = "/output/labels.txt"
	defaultClaudeAllowedTools = "Read,Write,Bash(git:*),Bash(mkdir),Bash(ls),Bash(cat)"
	defaultReviewAgent        = "claude"
	defaultAgentWorkingDir    = "/workspace"

	envReviewAgent          = "REVIEW_AGENT"
	envReviewPromptPath     = "REVIEW_PROMPT_PATH"
	envReviewRepositoryPath = "REVIEW_REPOSITORY_PATH"
	envReviewReportPath     = "REVIEW_REPORT_PATH"
	envReviewLabelsPath     = "REVIEW_LABELS_PATH"
	envClaudeReviewModel    = "CLAUDE_REVIEW_MODEL"
	envCodexReviewModel     = "CODEX_REVIEW_MODEL"
	envAgentAllowedTools    = "REVIEW_AGENT_ALLOWED_TOOLS"
	envAgentExtraArgs       = "REVIEW_AGENT_EXTRA_ARGS"
	envExtraDirs            = "REVIEW_EXTRA_ALLOWED_DIRS"
	envExtraFiles           = "REVIEW_EXTRA_ALLOWED_FILES"
	envCodexQuiet           = "CODEX_QUIET_MODE"
	envCodexJson            = "CODEX_JSON_MODE"
	envCodexWorkingDir      = "CODEX_WORKDIR"
)

// ReviewMode enumerates supported security review modes.
type ReviewMode string

const (
	ReviewModeFull ReviewMode = "full"
	ReviewModeDiff ReviewMode = "diff"
)

// agentInvocation captures execution hints per reviewer agent.
type agentInvocation struct {
	Prompt       string
	Model        string
	AllowedTools string
	AllowedDirs  []string
	AllowedFiles []string
	ExtraArgs    string
	WorkingDir   string
}

// main configures logging, resolves environment, and runs the selected agent.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Run the review workflow and exit non-zero on failure so the container signals an error.
	if err := run(ctx); err != nil {
		logError(err)
		os.Exit(1)
	}
}

// run orchestrates prompt generation and agent execution.
func run(ctx context.Context) error {
	// Parse review configuration from the environment.
	modeRaw, err := fetchEnv("REVIEW_MODE", true)
	if err != nil {
		return err
	}
	mode, err := normalizeMode(modeRaw)
	if err != nil {
		return err
	}

	requireHead := mode != ReviewModeFull
	headSHA, err := fetchEnv("REVIEW_HEAD_SHA", requireHead)
	if err != nil {
		return err
	}
	baseSHA, err := fetchEnv("REVIEW_BASE_SHA", false)
	if err != nil {
		return err
	}
	if mode == ReviewModeDiff && baseSHA == "" {
		return errors.New("REVIEW_BASE_SHA is required when REVIEW_MODE=diff")
	}

	targetLabel, err := fetchEnv("REVIEW_TARGET_LABEL", false)
	if err != nil {
		return err
	}

	// Resolve concrete paths for prompt, repository, and outputs.
	promptPath := strings.TrimSpace(firstNonEmpty(
		mustFetchOptional(envReviewPromptPath),
		defaultPromptPath,
	))
	repositoryPath := strings.TrimSpace(firstNonEmpty(
		mustFetchOptional(envReviewRepositoryPath),
		defaultRepositoryPath,
	))
	reportPath := strings.TrimSpace(firstNonEmpty(
		mustFetchOptional(envReviewReportPath),
		defaultReportPath,
	))
	labelsPath := strings.TrimSpace(firstNonEmpty(
		mustFetchOptional(envReviewLabelsPath),
		defaultLabelsPath,
	))

	promptPath = filepath.Clean(promptPath)
	repositoryPath = filepath.Clean(repositoryPath)
	reportPath = filepath.Clean(reportPath)
	labelsPath = filepath.Clean(labelsPath)
	reportDir := filepath.Dir(reportPath)
	labelsDir := filepath.Dir(labelsPath)

	// Read the rendered prompt and ensure the repository mount is present.
	if err = ensureDirectory(repositoryPath); err != nil {
		return err
	}

	promptBytes, err := os.ReadFile(promptPath)
	if err != nil {
		return fmt.Errorf("read prompt %s: %w", promptPath, err)
	}
	logInfo(fmt.Sprintf("Loaded prompt from %s.", promptPath))

	// Select the reviewer implementation and build invocation parameters.
	agentName, err := fetchEnv(envReviewAgent, false)
	if err != nil {
		return err
	}
	if agentName == "" {
		agentName = defaultReviewAgent
	}
	agentKey := strings.ToLower(strings.TrimSpace(agentName))
	agent, err := selectAgent(agentKey)
	if err != nil {
		return err
	}

	var model string
	if envName := agent.ModelEnvVar(); envName != "" {
		model = mustFetchOptional(envName)
	}

	allowedTools := firstNonEmpty(
		mustFetchOptional(envAgentAllowedTools),
		mustFetchOptional("CLAUDE_ALLOWED_TOOLS"),
		agent.DefaultAllowedTools(),
	)
	extraArgs := firstNonEmpty(
		mustFetchOptional(envAgentExtraArgs),
		mustFetchOptional("CLAUDE_EXTRA_ARGS"),
	)

	allowedDirs := []string{repositoryPath, defaultAgentWorkingDir, reportDir, labelsDir}
	if extraDirs := mustFetchOptional(envExtraDirs); extraDirs != "" {
		allowedDirs = append(allowedDirs, parseList(extraDirs)...)
	}

	allowedFiles := []string{reportTemplatePath, promptPath}
	if extraFiles := mustFetchOptional(envExtraFiles); extraFiles != "" {
		allowedFiles = append(allowedFiles, parseList(extraFiles)...)
	}

	inv := agentInvocation{
		Prompt:       string(promptBytes),
		Model:        model,
		AllowedTools: allowedTools,
		AllowedDirs:  allowedDirs,
		AllowedFiles: allowedFiles,
		ExtraArgs:    extraArgs,
		WorkingDir:   defaultAgentWorkingDir,
	}

	logInfo(fmt.Sprintf(
		"Starting %s review (agent=%s head=%s base=%s label=%s).",
		mode, agent.Name(), headSHA, baseSHA, targetLabel,
	))

	// Execute the agent command and relay its output streams.
	stdout, stderr, runErr := runAgent(ctx, agent, inv)
	if stderr != "" {
		logError(errors.New(stderr))
	}
	if stdout != "" {
		fmt.Print(stdout)
	}
	if runErr != nil {
		return runErr
	}

	// Persist the report and labels outputs, falling back to stdout when needed.
	if err = ensureParent(reportPath); err != nil {
		return err
	}
	if !fileExists(reportPath) {
		if err = os.WriteFile(reportPath, []byte(stdout), 0o644); err != nil {
			return fmt.Errorf("write fallback report: %w", err)
		}
		logInfo("Report not found, wrote fallback using stdout output.")
	}

	if err = ensureParent(labelsPath); err != nil {
		return err
	}
	if err = ensureLabelsFile(labelsPath); err != nil {
		return err
	}

	logInfo("Security review completed successfully.")
	logInfo(fmt.Sprintf("Report stored at %s.", reportPath))
	logInfo(fmt.Sprintf("Labels stored at %s.", labelsPath))
	return nil
}

// fetchEnv reads an environment variable and validates presence when required.
func fetchEnv(name string, required bool) (string, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" && required {
		return "", fmt.Errorf("missing required environment variable: %s", name)
	}
	return value, nil
}

// mustFetchOptional retrieves an optional environment variable without error returns.
func mustFetchOptional(name string) string {
	value, _ := fetchEnv(name, false)
	return value
}

// firstNonEmpty returns the first non-empty string from the provided list.
func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

// ensureDirectory verifies that the provided path exists and is a directory.
func ensureDirectory(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat %s: %w", path, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("expected directory at %s", path)
	}
	return nil
}

// normalizeMode validates mode strings and returns canonical values.
func normalizeMode(raw string) (ReviewMode, error) {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "diff", "differential":
		return ReviewModeDiff, nil
	case "full":
		return ReviewModeFull, nil
	default:
		return "", fmt.Errorf("invalid REVIEW_MODE: %s", raw)
	}
}

// ensureParent creates directories needed for the provided path.
func ensureParent(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// ensureLabelsFile guarantees the labels file exists as a regular file.
func ensureLabelsFile(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return fmt.Errorf("expected file at %s", path)
		}
		return nil
	}
	if errors.Is(err, os.ErrNotExist) {
		if writeErr := os.WriteFile(path, []byte{}, 0o644); writeErr != nil {
			return fmt.Errorf("create labels file %s: %w", path, writeErr)
		}
		logInfo(fmt.Sprintf("Labels file not found, created empty file at %s.", path))
		return nil
	}
	return fmt.Errorf("stat labels file %s: %w", path, err)
}

// parseList splits whitespace separated values into a slice.
func parseList(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	return strings.Fields(raw)
}

// runAgent executes the reviewer agent command and captures output streams.
func runAgent(ctx context.Context, agent reviewerAgent, inv agentInvocation) (string, string, error) {
	cmd, err := agent.BuildCommand(ctx, inv)
	if err != nil {
		return "", "", err
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	if err = cmd.Run(); err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("%s invocation failed: %w", agent.Name(), err)
	}

	return stdout.String(), stderr.String(), nil
}

// fileExists returns true when a non-zero length file exists at path.
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return info.Size() > 0
}

// logInfo prints informational messages prefixed for clarity.
func logInfo(msg string) {
	fmt.Printf("[security-reviewer] %s\n", msg)
}

// logError prints error messages prefixed for clarity.
func logError(err error) {
	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		fmt.Fprintf(os.Stderr, "[security-reviewer] ERROR: %s (%s)\n", pathErr.Path, pathErr.Err)
		return
	}
	fmt.Fprintf(os.Stderr, "[security-reviewer] ERROR: %s\n", err)
}

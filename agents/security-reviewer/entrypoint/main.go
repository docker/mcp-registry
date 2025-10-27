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
	promptTemplatePath        = "/opt/security-reviewer/prompt-template.md"
	reportTemplatePath        = "/opt/security-reviewer/report-template.md"
	defaultPromptPath         = "/workspace/input/prompt.md"
	defaultRepositoryPath     = "/workspace/input/repository"
	defaultReportPath         = "/workspace/output/report.md"
	defaultLabelsPath         = "/workspace/output/labels.txt"
	defaultClaudeAllowedTools = "Read,Write,Bash(git:*),Bash(mkdir),Bash(ls),Bash(cat)"
	defaultReviewAgent        = "claude"
	defaultAgentWorkingDir    = "/workspace"

	envReviewAgent     = "REVIEW_AGENT"
	envAgentExtraArgs  = "REVIEW_AGENT_EXTRA_ARGS"
	envCodexQuiet      = "CODEX_QUIET_MODE"
	envCodexJson       = "CODEX_JSON_MODE"
	envCodexWorkingDir = "CODEX_WORKDIR"
)

// ReviewMode enumerates supported security review modes.
type ReviewMode string

const (
	// ReviewModeFull requests a full repository audit.
	ReviewModeFull ReviewMode = "full"
	// ReviewModeDiff requests a differential review between two commits.
	ReviewModeDiff ReviewMode = "diff"
)

// agentInvocation captures execution hints per reviewer agent.
type agentInvocation struct {
	// Prompt is the rendered instruction text passed over stdin.
	Prompt string
	// Model identifies the model to invoke, when the agent supports overrides.
	Model string
	// AllowedTools enumerates tool permissions for agents that honor them.
	AllowedTools string
	// AllowedDirs lists directories the agent should be allowed to traverse.
	AllowedDirs []string
	// AllowedFiles lists specific files the agent may read or write.
	AllowedFiles []string
	// ExtraArgs contains caller-supplied CLI arguments for the agent.
	ExtraArgs string
	// WorkingDir specifies the directory where the agent command executes.
	WorkingDir string
}

// promptPlaceholders stores values substituted into the static prompt template.
type promptPlaceholders struct {
	// ModeLabel is the human friendly descriptor for the review mode.
	ModeLabel string
	// ModeSummary highlights the responsibilities for the current mode.
	ModeSummary string
	// TargetLabel is an identifier referencing the repository under review.
	TargetLabel string
	// RepositoryPath points to the checked-out repository mount.
	RepositoryPath string
	// HeadCommit is the commit under audit.
	HeadCommit string
	// BaseCommit is the comparison commit for diff reviews.
	BaseCommit string
	// CommitRange renders the <base>...<head> spec for diff reviews.
	CommitRange string
	// GitDiffHint guides the agent on how to inspect the change set.
	GitDiffHint string
	// ReportPath denotes where the agent should write its report.
	ReportPath string
	// LabelsPath denotes where the agent should write labels for automation.
	LabelsPath string
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
	promptPath := defaultPromptPath
	repositoryPath := defaultRepositoryPath
	reportPath := defaultReportPath
	labelsPath := defaultLabelsPath

	promptPath = filepath.Clean(promptPath)
	repositoryPath = filepath.Clean(repositoryPath)
	reportPath = filepath.Clean(reportPath)
	labelsPath = filepath.Clean(labelsPath)

	// Read the rendered prompt and ensure the repository mount is present.
	if err = ensureDirectory(repositoryPath); err != nil {
		return err
	}

	promptContent, err := buildPromptContent(mode, targetLabel, headSHA, baseSHA)
	if err != nil {
		return err
	}
	if err = ensureParent(promptPath); err != nil {
		return err
	}
	if err = os.WriteFile(promptPath, []byte(promptContent), 0o644); err != nil {
		return fmt.Errorf("write prompt: %w", err)
	}
	logInfo(fmt.Sprintf("Rendered prompt to %s.", promptPath))

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

	allowedTools := agent.DefaultAllowedTools()
	extraArgs := mustFetchOptional(envAgentExtraArgs)

	allowedDirs := []string{defaultAgentWorkingDir}

	allowedFiles := []string{reportTemplatePath, promptPath}

	inv := agentInvocation{
		Prompt:       promptContent,
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

// normalizeMode converts raw user input into a canonical ReviewMode value.
func normalizeMode(raw string) (ReviewMode, error) {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case string(ReviewModeDiff), "differential":
		return ReviewModeDiff, nil
	case string(ReviewModeFull):
		return ReviewModeFull, nil
	default:
		return "", fmt.Errorf("invalid review mode: %s", raw)
	}
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

// buildPromptContent renders a concrete prompt for the selected review mode.
func buildPromptContent(mode ReviewMode, targetLabel, headSHA, baseSHA string) (string, error) {
	displayLabel := strings.TrimSpace(targetLabel)
	if displayLabel == "" {
		displayLabel = "Not provided"
	}
	displayHead := strings.TrimSpace(headSHA)
	if displayHead == "" {
		displayHead = "Not provided"
	}
	displayBase := "Not applicable"
	commitRange := "Not applicable"
	if mode == ReviewModeDiff {
		cleanBase := strings.TrimSpace(baseSHA)
		cleanHead := strings.TrimSpace(headSHA)
		if cleanBase == "" {
			displayBase = "Not provided"
		} else {
			displayBase = cleanBase
		}
		if cleanBase != "" && cleanHead != "" {
			commitRange = fmt.Sprintf("%s...%s", baseSHA, headSHA)
		}
	}

	ph := promptPlaceholders{
		ModeLabel:      modeLabel(mode),
		ModeSummary:    modeSummary(mode),
		TargetLabel:    displayLabel,
		RepositoryPath: defaultRepositoryPath,
		HeadCommit:     displayHead,
		BaseCommit:     displayBase,
		CommitRange:    commitRange,
		GitDiffHint:    gitDiffHint(mode, baseSHA, headSHA),
		ReportPath:     defaultReportPath,
		LabelsPath:     defaultLabelsPath,
	}
	return renderPrompt(ph)
}

// renderPrompt injects placeholder values into the prompt template.
func renderPrompt(ph promptPlaceholders) (string, error) {
	templateBytes, err := os.ReadFile(promptTemplatePath)
	if err != nil {
		return "", fmt.Errorf("read prompt template: %w", err)
	}
	replacer := strings.NewReplacer(
		"$MODE_LABEL", ph.ModeLabel,
		"$MODE_SUMMARY", ph.ModeSummary,
		"$TARGET_LABEL", ph.TargetLabel,
		"$REPOSITORY_PATH", ph.RepositoryPath,
		"$HEAD_COMMIT", ph.HeadCommit,
		"$BASE_COMMIT", ph.BaseCommit,
		"$COMMIT_RANGE", ph.CommitRange,
		"$GIT_DIFF_HINT", ph.GitDiffHint,
		"$REPORT_PATH", ph.ReportPath,
		"$LABELS_PATH", ph.LabelsPath,
	)
	return replacer.Replace(string(templateBytes)), nil
}

// gitDiffHint conveys how the agent should inspect the repository state.
func gitDiffHint(mode ReviewMode, baseSHA, headSHA string) string {
	if mode == ReviewModeDiff {
		cleanBase := strings.TrimSpace(baseSHA)
		cleanHead := strings.TrimSpace(headSHA)
		if cleanBase == "" || cleanHead == "" {
			return fmt.Sprintf("Run `git diff` inside %s to inspect the change set.", defaultRepositoryPath)
		}
		return fmt.Sprintf("Run `git diff %s...%s` (and related commands) inside %s to inspect the change set.", baseSHA, headSHA, defaultRepositoryPath)
	}
	return "Audit the entire working tree at the head commit."
}

// modeLabel converts a review mode to a user friendly label.
func modeLabel(mode ReviewMode) string {
	switch mode {
	case ReviewModeDiff:
		return "Differential"
	case ReviewModeFull:
		return "Full"
	default:
		return "Unknown"
	}
}

// modeSummary explains the responsibilities associated with a review mode.
func modeSummary(mode ReviewMode) string {
	switch mode {
	case ReviewModeDiff:
		return "You are reviewing the changes introduced between the base and head commits. Prioritize spotting deliberately malicious additions alongside accidental vulnerabilities."
	case ReviewModeFull:
		return "You are auditing the repository snapshot at the provided head commit. Assume attackers may have hidden malicious logic and hunt for both intentional and accidental risks."
	default:
		return "The review mode is unknown."
	}
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

// ensureDirectory verifies that the supplied path exists and is a directory.
func ensureDirectory(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat directory %s: %w", path, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}

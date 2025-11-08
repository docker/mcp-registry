package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	// promptTemplatePath is the location of the static prompt template bundled with the image.
	promptTemplatePath = "/opt/security-reviewer/prompt-template.md"
	// reportTemplatePath is the location of the report template referenced in prompts.
	reportTemplatePath = "/opt/security-reviewer/report-template.md"
	// defaultPromptPath is where the rendered prompt is written for the agent to consume.
	defaultPromptPath = "/workspace/input/prompt.md"
	// defaultRepositoryPath is the mount point containing the repository under review.
	defaultRepositoryPath = "/workspace/input/repository"
	// defaultReportPath is the expected location for the agent's security report.
	defaultReportPath = "/workspace/output/report.md"
	// defaultLabelsPath is the expected location for the agent's label output.
	defaultLabelsPath = "/workspace/output/labels.txt"
	// defaultReviewAgent is the reviewer implementation used when none is specified.
	defaultReviewAgent = "claude"
	// defaultAgentWorkingDir is the directory from which the agent executes.
	defaultAgentWorkingDir = "/workspace"
	// defaultTimeout bounds how long the reviewer will wait for the agent to complete.
	defaultTimeout = time.Hour

	// envReviewAgent selects which reviewer agent to run.
	envReviewAgent = "REVIEW_AGENT"
	// envAgentExtraArgs contains optional CLI arguments passed through to the agent.
	envAgentExtraArgs = "REVIEW_AGENT_EXTRA_ARGS"
	// envReviewHeadSHA provides the head commit SHA under review.
	envReviewHeadSHA = "REVIEW_HEAD_SHA"
	// envReviewBaseSHA provides the base commit SHA when performing differential reviews.
	envReviewBaseSHA = "REVIEW_BASE_SHA"
	// envReviewTarget supplies a human-readable label describing the review target.
	envReviewTarget = "REVIEW_TARGET_LABEL"
	// envReviewTimeout allows callers to override the agent execution timeout in seconds.
	envReviewTimeout = "REVIEW_TIMEOUT_SECS"
)

// ReviewMode enumerates supported security review modes.
type ReviewMode string

const (
	// ReviewModeFull requests a full repository audit.
	ReviewModeFull ReviewMode = "full"
	// ReviewModeDifferential requests a differential review between two commits.
	ReviewModeDifferential ReviewMode = "differential"
)

// agentInvocation captures execution hints per reviewer agent.
type agentInvocation struct {
	// Prompt is the rendered instruction text passed over stdin.
	Prompt string
	// Model identifies the model to invoke, when the agent supports overrides.
	Model string
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
	// ReportTemplatePath tells the agent which template to follow exactly.
	ReportTemplatePath string
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
	headSHA, err := fetchEnv(envReviewHeadSHA, true)
	if err != nil {
		return err
	}
	baseSHA, err := fetchEnv(envReviewBaseSHA, false)
	if err != nil {
		return err
	}

	mode := ReviewModeFull
	if baseSHA != "" {
		mode = ReviewModeDifferential
	}

	targetLabel, err := fetchEnv(envReviewTarget, false)
	if err != nil {
		return err
	}

	// Ensure the repository mount is present before processing.
	if err = ensureDirectory(defaultRepositoryPath); err != nil {
		return err
	}

	promptContent, err := buildPromptContent(mode, targetLabel, headSHA, baseSHA)
	if err != nil {
		return err
	}
	if err = ensureParent(defaultPromptPath); err != nil {
		return err
	}
	if err = os.WriteFile(defaultPromptPath, []byte(promptContent), 0o644); err != nil {
		return fmt.Errorf("write prompt: %w", err)
	}
	logInfo(fmt.Sprintf("Rendered prompt to %s.", defaultPromptPath))

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
		model, err = fetchEnv(envName, false)
		if err != nil {
			return err
		}
	}

	extraArgs, _ := fetchEnv(envAgentExtraArgs, false)
	inv := agentInvocation{
		Prompt:     promptContent,
		Model:      model,
		ExtraArgs:  extraArgs,
		WorkingDir: defaultAgentWorkingDir,
	}

	timeout, err := resolveTimeout()
	if err != nil {
		return err
	}

	agentCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if mode == ReviewModeDifferential {
		logInfo(fmt.Sprintf(
			"Starting differential review (agent=%s head=%s base=%s label=%s).",
			agent.Name(), headSHA, baseSHA, targetLabel,
		))
	} else {
		logInfo(fmt.Sprintf(
			"Starting full review (agent=%s head=%s label=%s).",
			agent.Name(), headSHA, targetLabel,
		))
	}

	// Execute the agent command and relay its output streams.
	if err := runAgent(agentCtx, agent, inv); err != nil {
		return err
	}

	// Inspect the outputs and warn if they were not produced.
	if fileExists(defaultReportPath) {
		logInfo(fmt.Sprintf("Report stored at %s.", defaultReportPath))
	} else {
		logWarn(fmt.Sprintf("Report not produced at %s.", defaultReportPath))
	}
	if fileExists(defaultLabelsPath) {
		logInfo(fmt.Sprintf("Labels stored at %s.", defaultLabelsPath))
	} else {
		logWarn(fmt.Sprintf("Labels not produced at %s.", defaultLabelsPath))
	}

	logInfo("Security review completed successfully.")
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

// ensureParent creates directories needed for the provided path.
func ensureParent(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// runAgent executes the reviewer agent command and captures output streams.
func runAgent(ctx context.Context, agent reviewerAgent, inv agentInvocation) error {
	cmd, err := agent.BuildCommand(ctx, inv)
	if err != nil {
		return err
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("%s invocation failed: %w", agent.Name(), err)
	}
	return nil
}

func resolveTimeout() (time.Duration, error) {
	value := strings.TrimSpace(os.Getenv(envReviewTimeout))
	if value == "" {
		return defaultTimeout, nil
	}
	secs, err := strconv.Atoi(value)
	if err != nil || secs <= 0 {
		return 0, fmt.Errorf("invalid %s value %q", envReviewTimeout, value)
	}
	return time.Duration(secs) * time.Second, nil
}

// buildPromptContent renders a concrete prompt for the selected review mode.
func buildPromptContent(mode ReviewMode, targetLabel, headSHA, baseSHA string) (string, error) {
	displayLabel := strings.TrimSpace(targetLabel)
	if displayLabel == "" {
		displayLabel = "[Not provided]"
	}
	baseDisplay := "[Not applicable]"
	commitRange := "[Not applicable]"
	if mode == ReviewModeDifferential {
		baseDisplay = baseSHA
		commitRange = fmt.Sprintf("%s...%s", baseSHA, headSHA)
	}

	ph := promptPlaceholders{
		ModeLabel:          modeLabel(mode),
		ModeSummary:        modeSummary(mode),
		TargetLabel:        displayLabel,
		RepositoryPath:     defaultRepositoryPath,
		HeadCommit:         headSHA,
		BaseCommit:         baseDisplay,
		CommitRange:        commitRange,
		GitDiffHint:        gitDiffHint(mode, baseSHA, headSHA),
		ReportPath:         defaultReportPath,
		LabelsPath:         defaultLabelsPath,
		ReportTemplatePath: reportTemplatePath,
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
		"$REPORT_TEMPLATE_PATH", ph.ReportTemplatePath,
	)
	return replacer.Replace(string(templateBytes)), nil
}

// gitDiffHint conveys how the agent should inspect the repository state.
func gitDiffHint(mode ReviewMode, baseSHA, headSHA string) string {
	if mode == ReviewModeDifferential {
		return fmt.Sprintf("Run `git diff %s...%s` (and related commands) inside %s to inspect the change set.", baseSHA, headSHA, defaultRepositoryPath)
	}
	return "Audit the entire working tree at the head commit."
}

// modeLabel converts a review mode to a user friendly label.
func modeLabel(mode ReviewMode) string {
	switch mode {
	case ReviewModeDifferential:
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
	case ReviewModeDifferential:
		return "You are reviewing the changes introduced in a Git repository between the specified base and head commits. Prioritize spotting deliberately malicious additions alongside accidental vulnerabilities."
	case ReviewModeFull:
		return "You are auditing a Git repository snapshot at the specified head commit. Assume attackers may have hidden malicious logic and hunt for both intentional and accidental risks."
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
	if info.Mode()&os.ModeType != 0 {
		return false
	}
	return info.Size() > 0
}

// logInfo prints informational messages prefixed for clarity.
func logInfo(msg string) {
	fmt.Printf("[security-reviewer] %s\n", msg)
}

// logWarn prints warning messages prefixed for clarity.
func logWarn(msg string) {
	fmt.Printf("[security-reviewer] WARNING: %s\n", msg)
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

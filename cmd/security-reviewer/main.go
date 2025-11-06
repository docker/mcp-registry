package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

const (
	// composeFileName is the compose manifest executed for each review run.
	composeFileName = "compose.yml"
	// reportFileName is the name of the individual report emitted by the agent.
	reportFileName = "report.md"
	// labelsFileName is the name of the label output emitted by the agent.
	labelsFileName = "labels.txt"
	// repositoryDirName is the working directory used to stage repository clones.
	repositoryDirName = "repository"
	// dockerExecutable identifies the docker CLI binary invoked by the tool.
	dockerExecutable = "docker"
	// gitExecutable identifies the git CLI binary used to manage repositories.
	gitExecutable = "git"
	// projectPrefix is applied to compose project names to make them unique yet readable.
	projectPrefix = "security-reviewer"
	// agentService is the compose service name running the security reviewer container.
	agentService = "reviewer"
	// composeRelativePath is the path to the compose project relative to the repository root.
	composeRelativePath = "agents/security-reviewer"
	// defaultTimeoutSeconds bounds agent execution time when not explicitly configured.
	defaultTimeoutSeconds = 3600

	// envAnthropicAPIKey is the environment variable supplying Claude credentials.
	envAnthropicAPIKey = "ANTHROPIC_API_KEY"
	// envOpenAIAPIKey is the environment variable supplying Codex credentials.
	envOpenAIAPIKey = "OPENAI_API_KEY"

	// agentNameClaude identifies the Claude-based reviewer.
	agentNameClaude = "claude"
	// agentNameCodex identifies the Codex-based reviewer.
	agentNameCodex = "codex"
)

// options stores parsed CLI arguments.
type options struct {
	// Agent selects the underlying reviewer agent implementation.
	Agent string
	// Repository is the Git repository URL or filesystem path.
	Repository string
	// HeadSHA is the commit under audit.
	HeadSHA string
	// BaseSHA is the comparison commit for differential reviews.
	BaseSHA string
	// TargetLabel is an optional human friendly descriptor.
	TargetLabel string
	// OutputPath is the destination for the final report.
	OutputPath string
	// LabelsOutput is the destination for the label list produced by the reviewer.
	LabelsOutput string
	// Model optionally overrides the reviewer model selection.
	Model string
	// ExtraArgs optionally appends raw arguments to the agent CLI.
	ExtraArgs string
	// KeepWorkdir preserves the temporary workspace when true.
	KeepWorkdir bool
	// TimeoutSeconds bounds the reviewer runtime; zero uses the default.
	TimeoutSeconds int
}

var cliOpts = options{
	Agent:          agentNameClaude,
	OutputPath:     "security-review.md",
	TimeoutSeconds: defaultTimeoutSeconds,
}

var rootCmd = &cobra.Command{
	Use:   "security-reviewer",
	Short: "Run the security reviewer compose workflow",
	RunE: func(cmd *cobra.Command, args []string) error {
		agent := strings.ToLower(strings.TrimSpace(cliOpts.Agent))
		if agent == "" {
			agent = agentNameClaude
		}
		if agent != agentNameClaude && agent != agentNameCodex {
			return fmt.Errorf("invalid agent %q (supported: %s, %s)", cliOpts.Agent, agentNameClaude, agentNameCodex)
		}

		labelsOutput := strings.TrimSpace(cliOpts.LabelsOutput)
		if labelsOutput == "" {
			labelsOutput = deriveDefaultLabelsPath(cliOpts.OutputPath)
		}
		timeoutSecs := cliOpts.TimeoutSeconds
		if timeoutSecs <= 0 {
			timeoutSecs = defaultTimeoutSeconds
		}

		opts := options{
			Agent:          agent,
			Repository:     strings.TrimSpace(cliOpts.Repository),
			HeadSHA:        strings.TrimSpace(cliOpts.HeadSHA),
			BaseSHA:        strings.TrimSpace(cliOpts.BaseSHA),
			TargetLabel:    strings.TrimSpace(cliOpts.TargetLabel),
			OutputPath:     strings.TrimSpace(cliOpts.OutputPath),
			LabelsOutput:   labelsOutput,
			Model:          strings.TrimSpace(cliOpts.Model),
			ExtraArgs:      strings.TrimSpace(cliOpts.ExtraArgs),
			KeepWorkdir:    cliOpts.KeepWorkdir,
			TimeoutSeconds: timeoutSecs,
		}

		if opts.Repository == "" {
			return errors.New("--repo is required")
		}
		if opts.HeadSHA == "" {
			return errors.New("--head is required")
		}
		ctx := cmd.Context()
		return run(ctx, opts)
	},
}

func init() {
	rootCmd.Flags().StringVar(&cliOpts.Agent, "agent", cliOpts.Agent, "Reviewer agent to use (claude or codex).")
	rootCmd.Flags().StringVar(&cliOpts.Repository, "repo", cliOpts.Repository, "Git repository URL or local path to review.")
	rootCmd.Flags().StringVar(&cliOpts.HeadSHA, "head", cliOpts.HeadSHA, "Head commit SHA to review.")
	rootCmd.Flags().StringVar(&cliOpts.BaseSHA, "base", cliOpts.BaseSHA, "Base commit SHA for differential reviews.")
	rootCmd.Flags().StringVar(&cliOpts.TargetLabel, "target-label", cliOpts.TargetLabel, "Human readable identifier for the target.")
	rootCmd.Flags().StringVar(&cliOpts.OutputPath, "output", cliOpts.OutputPath, "Destination for the rendered report.")
	rootCmd.Flags().StringVar(&cliOpts.LabelsOutput, "labels-output", cliOpts.LabelsOutput, "Destination for the labels file (defaults alongside the report).")
	rootCmd.Flags().IntVar(&cliOpts.TimeoutSeconds, "timeout", cliOpts.TimeoutSeconds, "Maximum runtime for the review in seconds (defaults to 3600 seconds).")
	rootCmd.Flags().StringVar(&cliOpts.Model, "model", cliOpts.Model, "Override the reviewer model for the selected agent.")
	rootCmd.Flags().StringVar(&cliOpts.ExtraArgs, "extra-args", cliOpts.ExtraArgs, "Additional arguments passed to the reviewer agent.")
	rootCmd.Flags().BoolVar(&cliOpts.KeepWorkdir, "keep-workdir", cliOpts.KeepWorkdir, "Keep the temporary workspace after completion.")

	_ = rootCmd.MarkFlagRequired("repo")
	_ = rootCmd.MarkFlagRequired("head")
}

// main is the entry point for the security reviewer CLI.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	rootCmd.SilenceUsage = true
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		exitWithError(err)
	}
}

// run coordinates workspace preparation, compose execution, and cleanup.
func run(ctx context.Context, opts options) error {
	if opts.BaseSHA != "" {
		fmt.Printf(
			"Starting differential security review (agent=%s head=%s base=%s target=%s)\n",
			opts.Agent,
			opts.HeadSHA,
			opts.BaseSHA,
			opts.TargetLabel,
		)
	} else {
		fmt.Printf(
			"Starting full security review (agent=%s head=%s target=%s)\n",
			opts.Agent,
			opts.HeadSHA,
			opts.TargetLabel,
		)
	}

	// Ensure the credential proxy has the API keys it needs before staging work.
	switch opts.Agent {
	case "claude":
		if _, ok := os.LookupEnv(envAnthropicAPIKey); !ok {
			return errors.New("ANTHROPIC_API_KEY environment variable is required for the Claude agent")
		}
	case "codex":
		if _, ok := os.LookupEnv(envOpenAIAPIKey); !ok {
			return errors.New("OPENAI_API_KEY environment variable is required for the Codex agent")
		}
	}

	// Prepare a temporary workspace to stage inputs and outputs.
	workdir, err := os.MkdirTemp("", fmt.Sprintf("security-reviewer-%s-", opts.Agent))
	if err != nil {
		return fmt.Errorf("create temporary directory: %w", err)
	}

	if !opts.KeepWorkdir {
		defer os.RemoveAll(workdir)
	} else {
		fmt.Printf("Temporary workspace preserved at %s\n", workdir)
	}

	// Materialize the repository commits required for the review.
	repositoryDir := filepath.Join(workdir, repositoryDirName)
	if err = prepareRepository(ctx, opts, repositoryDir); err != nil {
		return err
	}

	outputDir := filepath.Join(workdir, "output")
	if err = os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	// Launch the compose project and wait for the reviewer to finish.
	if err = runCompose(ctx, opts, workdir, repositoryDir, outputDir); err != nil {
		return err
	}

	// Copy the generated artifacts back to the requested destinations.
	reportPath := filepath.Join(outputDir, reportFileName)
	labelsPath := filepath.Join(outputDir, labelsFileName)
	if _, err = os.Stat(reportPath); err != nil {
		return fmt.Errorf("review report not produced: %w", err)
	}
	if _, err = os.Stat(labelsPath); err != nil {
		return fmt.Errorf("labels file not produced: %w", err)
	}

	if err = copyFile(reportPath, opts.OutputPath); err != nil {
		return err
	}
	if err = copyFile(labelsPath, opts.LabelsOutput); err != nil {
		return err
	}

	fmt.Printf("Security review report copied to %s\n", opts.OutputPath)
	fmt.Printf("Security review labels copied to %s\n", opts.LabelsOutput)
	return nil
}

// deriveDefaultLabelsPath produces a labels output path near the report path.
func deriveDefaultLabelsPath(reportPath string) string {
	reportPath = strings.TrimSpace(reportPath)
	if reportPath == "" {
		// Without an explicit report, fall back to a stable default name.
		return "security-review-labels.txt"
	}
	// Place the labels file alongside the report for easier discovery.
	dir := filepath.Dir(reportPath)
	base := filepath.Base(reportPath)
	idx := strings.LastIndex(base, ".")
	if idx > 0 {
		// Drop the extension so the generated labels file mirrors the report name.
		base = base[:idx]
	}
	if strings.TrimSpace(base) == "" {
		base = "security-review"
	}
	// Append a suffix to distinguish the labels artifact from the report.
	return filepath.Join(dir, base+"-labels.txt")
}

// sanitizeName converts arbitrary text into a slug.
func sanitizeName(text string) string {
	lower := strings.ToLower(text)
	pattern := regexp.MustCompile(`[^a-z0-9]+`)
	cleaned := pattern.ReplaceAllString(lower, "-")
	trimmed := strings.Trim(cleaned, "-")
	if trimmed == "" {
		return "target"
	}
	return trimmed
}

// prepareRepository clones the repository and materializes commits for review.
func prepareRepository(ctx context.Context, opts options, repositoryDir string) error {
	parentDir := filepath.Dir(repositoryDir)
	if err := os.MkdirAll(parentDir, 0o755); err != nil {
		return fmt.Errorf("create repository parent directory: %w", err)
	}
	if err := os.RemoveAll(repositoryDir); err != nil {
		return fmt.Errorf("reset repository directory: %w", err)
	}

	if err := runCommand(ctx, "", gitExecutable, "clone", opts.Repository, repositoryDir); err != nil {
		return fmt.Errorf("clone repository: %w", err)
	}

	if err := ensureCommit(ctx, repositoryDir, opts.HeadSHA); err != nil {
		return err
	}
	if err := runCommand(ctx, repositoryDir, gitExecutable, "checkout", "--detach", opts.HeadSHA); err != nil {
		return fmt.Errorf("checkout head commit: %w", err)
	}

	if opts.BaseSHA != "" {
		if err := ensureCommit(ctx, repositoryDir, opts.BaseSHA); err != nil {
			return err
		}
	}

	return nil
}

// ensureCommit verifies that a commit exists locally, fetching if needed.
func ensureCommit(ctx context.Context, repoDir, sha string) error {
	if sha == "" {
		return nil
	}
	if err := runCommand(ctx, repoDir, gitExecutable, "rev-parse", "--verify", sha); err == nil {
		return nil
	}
	if err := runCommand(ctx, repoDir, gitExecutable, "fetch", "origin", sha); err != nil {
		return fmt.Errorf("fetch commit %s: %w", sha, err)
	}
	if err := runCommand(ctx, repoDir, gitExecutable, "rev-parse", "--verify", sha); err != nil {
		return fmt.Errorf("verify commit %s: %w", sha, err)
	}
	return nil
}

// copyFile copies a file from src to dst, creating parent directories.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open file %s: %w", src, err)
	}
	defer in.Close()
	if err = os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("create directory for %s: %w", dst, err)
	}
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("open destination %s: %w", dst, err)
	}
	defer out.Close()
	if _, err = io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s to %s: %w", src, dst, err)
	}
	return nil
}

// runCompose executes the docker compose workflow for the review.
func runCompose(ctx context.Context, opts options, workdir, repositoryDir, outputDir string) error {
	// Compose assumes relative paths, so stage a copy inside the temp workspace.
	composeDir := filepath.Join(workdir, composeRelativePath)
	if err := copyDir(composeRelativePath, composeDir); err != nil {
		return err
	}

	env := buildComposeEnv(opts, repositoryDir, outputDir)
	up := exec.CommandContext(ctx, dockerExecutable, "compose", "-f", composeFileName, "up", "--build", "--abort-on-container-exit", "--exit-code-from", agentService)
	up.Dir = composeDir
	up.Env = env
	up.Stdout = os.Stdout
	up.Stderr = os.Stderr

	down := exec.CommandContext(context.Background(), dockerExecutable, "compose", "-f", composeFileName, "down", "--volumes", "--remove-orphans")
	down.Dir = composeDir
	down.Env = env

	if err := up.Run(); err != nil {
		_ = down.Run()
		return fmt.Errorf("docker compose up: %w", err)
	}
	if err := down.Run(); err != nil {
		return fmt.Errorf("docker compose down: %w", err)
	}
	return nil
}

// buildComposeEnv prepares environment variables for docker compose.
func buildComposeEnv(opts options, repositoryDir, outputDir string) []string {
	env := os.Environ()
	// Generate a stable slug to keep compose project names readable.
	slug := sanitizeName(opts.TargetLabel)
	if slug == "target" {
		repoBase := filepath.Base(strings.TrimSuffix(opts.Repository, ".git"))
		slug = sanitizeName(repoBase)
	}
	if slug == "" {
		slug = "target"
	}
	projectName := fmt.Sprintf("%s-%s-%d", projectPrefix, slug, time.Now().Unix())

	// Get current UID/GID to match container user with host user.
	// This avoids permission issues with bind mounts.
	uid, gid := getCurrentUIDGID()

	env = append(env,
		fmt.Sprintf("COMPOSE_PROJECT_NAME=%s", projectName),
		fmt.Sprintf("REVIEW_AGENT=%s", opts.Agent),
		fmt.Sprintf("REVIEW_HEAD_SHA=%s", opts.HeadSHA),
		fmt.Sprintf("REVIEW_BASE_SHA=%s", opts.BaseSHA),
		fmt.Sprintf("REVIEW_TARGET_LABEL=%s", opts.TargetLabel),
		fmt.Sprintf("REVIEW_REPOSITORY_PATH=%s", repositoryDir),
		fmt.Sprintf("REVIEW_OUTPUT_PATH=%s", outputDir),
		fmt.Sprintf("REVIEW_TIMEOUT_SECS=%d", opts.TimeoutSeconds),
		fmt.Sprintf("AGENT_UID=%d", uid),
		fmt.Sprintf("AGENT_GID=%d", gid),
	)
	if opts.Model != "" {
		// Route custom models to the right environment variable per agent.
		switch strings.ToLower(opts.Agent) {
		case agentNameClaude:
			env = append(env, fmt.Sprintf("CLAUDE_REVIEW_MODEL=%s", opts.Model))
		case agentNameCodex:
			env = append(env, fmt.Sprintf("CODEX_REVIEW_MODEL=%s", opts.Model))
		}
	}
	if opts.ExtraArgs != "" {
		env = append(env, fmt.Sprintf("REVIEW_AGENT_EXTRA_ARGS=%s", opts.ExtraArgs))
	}
	if key := strings.TrimSpace(os.Getenv(envAnthropicAPIKey)); key != "" {
		env = append(env, fmt.Sprintf("%s=%s", envAnthropicAPIKey, key))
	}
	if key := strings.TrimSpace(os.Getenv(envOpenAIAPIKey)); key != "" {
		env = append(env, fmt.Sprintf("%s=%s", envOpenAIAPIKey, key))
	}
	return env
}

// copyDir performs a recursive directory copy.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, info.Mode())
	})
}

// runCommand executes a command within an optional directory.
func runCommand(ctx context.Context, dir, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// exitWithError prints an error and terminates the process.
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

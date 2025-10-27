package main

import (
	"context"
	"os/exec"
	"strings"
)

// claudeAgent implements reviewerAgent for Claude Code.
type claudeAgent struct{}

// Name returns the stable identifier for the Claude agent implementation.
func (claudeAgent) Name() string {
	return agentNameClaude
}

// ModelEnvVar exposes the environment variable used to override the model.
func (claudeAgent) ModelEnvVar() string {
	// Claude Code reads its target model from CLAUDE_REVIEW_MODEL.
	return "CLAUDE_REVIEW_MODEL"
}

// DefaultAllowedTools returns the default Claude tool allowlist.
func (claudeAgent) DefaultAllowedTools() string {
	// Mirror the default permissions granted in prior workflows.
	return defaultClaudeAllowedTools
}

// BuildCommand constructs the Claude CLI invocation for a review run.
func (claudeAgent) BuildCommand(ctx context.Context, inv agentInvocation) (*exec.Cmd, error) {
	// When running Claude Code in non-interactive mode, the only output format
	// that gives regular progress updates is stream-json - anything else waits
	// for the full analysis to complete and then provides all the output at
	// once. It would be nice if Claude Code had something like a stream-text
	// mode, and there's a request for that here:
	//   https://github.com/anthropics/claude-code/issues/4346
	// In the meantime, I think we'll just live with the JSON output, since at
	// least that gives some indication of progress and what's happening.
	args := []string{"--print", "--verbose", "--output-format", "stream-json"}
	if strings.TrimSpace(inv.AllowedTools) != "" {
		args = append(args, "--allowed-tools", inv.AllowedTools)
	}
	if strings.TrimSpace(inv.Model) != "" {
		args = append(args, "--model", inv.Model)
	}
	for _, dir := range inv.AllowedDirs {
		if strings.TrimSpace(dir) == "" {
			continue
		}
		args = append(args, "--add-dir", dir)
	}
	if strings.TrimSpace(inv.ExtraArgs) != "" {
		args = append(args, strings.Fields(inv.ExtraArgs)...)
	}

	cmd := exec.CommandContext(ctx, "claude", args...)
	cmd.Stdin = strings.NewReader(inv.Prompt)
	if inv.WorkingDir != "" {
		cmd.Dir = inv.WorkingDir
	}

	return cmd, nil
}

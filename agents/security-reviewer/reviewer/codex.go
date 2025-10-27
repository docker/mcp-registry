package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattn/go-shellwords"
)

// codexAgent implements reviewerAgent for the OpenAI Codex CLI.
type codexAgent struct{}

// Name returns the stable identifier for the Codex agent implementation.
func (codexAgent) Name() string {
	return agentNameCodex
}

// ModelEnvVar exposes the environment variable used to override Codex models.
func (codexAgent) ModelEnvVar() string {
	// Codex shells read from CODEX_REVIEW_MODEL when provided.
	return "CODEX_REVIEW_MODEL"
}

// BuildCommand constructs the Codex CLI invocation for a review run.
func (codexAgent) BuildCommand(ctx context.Context, inv agentInvocation) (*exec.Cmd, error) {
	args := []string{
		"exec",
		"--skip-git-repo-check",
		"--dangerously-bypass-approvals-and-sandbox",
	}
	if strings.TrimSpace(inv.Model) != "" {
		args = append(args, "--model", inv.Model)
	}
	if strings.TrimSpace(inv.ExtraArgs) != "" {
		parsed, err := shellwords.Parse(inv.ExtraArgs)
		if err != nil {
			return nil, fmt.Errorf("parse extra args: %w", err)
		}
		args = append(args, parsed...)
	}
	args = append(args, "-")

	cmd := exec.CommandContext(ctx, "codex", args...)
	cmd.Stdin = strings.NewReader(inv.Prompt)
	if inv.WorkingDir != "" {
		cmd.Dir = inv.WorkingDir
	}

	return cmd, nil
}

package main

import (
	"context"
	"os/exec"
	"strings"
)

// claudeAgent implements reviewerAgent for Claude Code.
type claudeAgent struct{}

func (claudeAgent) Name() string {
	return agentNameClaude
}

func (claudeAgent) ModelEnvVar() string {
	// Claude Code reads its target model from CLAUDE_REVIEW_MODEL.
	return "CLAUDE_REVIEW_MODEL"
}

func (claudeAgent) DefaultAllowedTools() string {
	// Mirror the default permissions granted in prior workflows.
	return defaultClaudeAllowedTools
}

func (claudeAgent) BuildCommand(ctx context.Context, inv agentInvocation) (*exec.Cmd, error) {
	args := []string{"--print", "--output-format", "text"}
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

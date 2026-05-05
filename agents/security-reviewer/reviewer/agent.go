package main

import (
	"context"
	"fmt"
	"os/exec"
)

const (
	// agentNameClaude identifies the Claude Code based reviewer.
	agentNameClaude = "claude"
	// agentNameCodex identifies the Codex based reviewer.
	agentNameCodex = "codex"
)

// reviewerAgent defines the behavior required by each agent implementation.
type reviewerAgent interface {
	Name() string
	// ModelEnvVar returns the environment variable that overrides the agent's model, or empty when not applicable.
	ModelEnvVar() string
	// BuildCommand returns the configured command used to invoke the agent.
	BuildCommand(ctx context.Context, inv agentInvocation) (*exec.Cmd, error)
}

// selectAgent resolves an agent by name.
func selectAgent(name string) (reviewerAgent, error) {
	switch name {
	case agentNameClaude:
		return claudeAgent{}, nil
	case agentNameCodex:
		return codexAgent{}, nil
	default:
		return nil, fmt.Errorf("unsupported review agent: %s", name)
	}
}

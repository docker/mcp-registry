/*
Copyright © 2025 Docker, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package mcp

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func runRequirement(ctx context.Context, requirement string) (func(), string, []string, error) {
	if requirement != "neo4j" {
		return nil, "", nil, fmt.Errorf("unsupported requirement: %s", requirement)
	}

	// Pull first to not count the pull duration in the timeout.
	cmdPull := exec.CommandContext(ctx, "docker", "pull", "neo4j")
	if err := cmdPull.Run(); err != nil {
		return nil, "", nil, fmt.Errorf("failed to pull Neo4j: %w", err)
	}

	// Run neo4j as a sidecar.
	ctxRequirement, cancel := context.WithCancel(ctx)

	var stdout bytes.Buffer

	containerName := fmt.Sprintf("neo4j-%s", randString(8))
	cmd := exec.CommandContext(ctxRequirement, "docker", "run", "--name", containerName, "--rm", "--init", "-e", "NEO4J_AUTH=none", "neo4j")
	cmd.Stdout = &stdout
	cmd.Cancel = func() error {
		return cmd.Process.Signal(syscall.SIGTERM)
	}

	if err := cmd.Start(); err != nil {
		cancel()
		return nil, "", nil, err
	}

	start := time.Now()
	started := false
waitStarted:
	for {
		select {
		case <-ctx.Done():
			cancel()
			return nil, "", nil, ctx.Err()
		case <-time.After(100 * time.Millisecond):
			if strings.Contains(stdout.String(), "Started.") {
				started = true
				break waitStarted
			}
			if time.Since(start) > 30*time.Second {
				break waitStarted
			}
		}
	}

	if !started {
		cancel()
		return nil, "", nil, fmt.Errorf("failed to start Neo4j: [%s]", stdout.String())
	}

	return cancel, containerName, []string{"NEO4J_URL=bolt://localhost:7687"}, nil
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

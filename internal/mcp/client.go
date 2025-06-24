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
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/docker/mcp-registry/pkg/servers"
	"github.com/mark3labs/mcp-go/mcp"
)

type client struct {
	image   string
	pull    bool
	env     []servers.Env
	secrets []servers.Secret
	args    []string
	command []string

	c *stdioMCPClient
}

func newClient(image string, pull bool, env []servers.Env, secrets []servers.Secret, args []string, command []string) *client {
	return &client{
		image:   image,
		pull:    pull,
		env:     env,
		secrets: secrets,
		args:    args,
		command: command,
	}
}

func (cl *client) Start(ctx context.Context, debug bool) error {
	if cl.c != nil {
		return fmt.Errorf("already started %s", cl.image)
	}

	if cl.pull {
		output, err := exec.CommandContext(ctx, "docker", "pull", cl.image).CombinedOutput()
		if err != nil {
			return fmt.Errorf("pulling image %s: %w (%s)", cl.image, err, string(output))
		}
	}

	args := []string{"run", "--rm", "-i", "--init", "--cap-drop=ALL"}
	args = append(args, cl.args...)
	for _, env := range cl.env {
		args = append(args, "-e", env.Name)
	}
	for _, secret := range cl.secrets {
		args = append(args, "-e", secret.Env)
	}
	args = append(args, cl.image)
	for _, arg := range cl.command {
		args = append(args, replacePlaceholders(arg, cl.env, cl.secrets))
	}
	c := newMCPClient("docker", toEnviron(cl.env, cl.secrets), args...)
	cl.c = c

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "docker",
		Version: "1.0.0",
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	if _, err := c.Initialize(ctx, initRequest, debug); err != nil {
		return fmt.Errorf("initializing %s: %w", cl.image, err)
	}
	return nil
}

func (cl *client) ListTools(ctx context.Context) ([]mcp.Tool, error) {
	if cl.c == nil {
		return nil, fmt.Errorf("listing tools %s: not started", cl.image)
	}

	response, err := cl.c.ListTools(ctx, mcp.ListToolsRequest{})
	if err != nil {
		return nil, fmt.Errorf("listing tools %s: %w", cl.image, err)
	}

	return response.Tools, nil
}

func (cl *client) ListPrompts(ctx context.Context) ([]mcp.Prompt, error) {
	if cl.c == nil {
		return nil, fmt.Errorf("listing tools %s: not started", cl.image)
	}

	response, err := cl.c.ListPrompts(ctx, mcp.ListPromptsRequest{})
	if err != nil {
		return nil, fmt.Errorf("listing tools %s: %w", cl.image, err)
	}

	return response.Prompts, nil
}

func (cl *client) CallTool(ctx context.Context, name string, args map[string]any) (*mcp.CallToolResult, error) {
	if cl.c == nil {
		return nil, fmt.Errorf("calling tool %s: not started", name)
	}

	request := mcp.CallToolRequest{}
	request.Params.Name = name
	request.Params.Arguments = args
	if request.Params.Arguments == nil {
		request.Params.Arguments = map[string]any{}
	}
	// MCP servers return an error if the args are empty so we make sure
	// there is at least one argument
	if len(request.Params.Arguments) == 0 {
		request.Params.Arguments["args"] = "..."
	}

	result, err := cl.c.CallTool(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("calling tool %s on %s: %w", name, cl.image, err)
	}

	return result, nil
}

func (cl *client) Close(deleteImage bool) error {
	if cl.c == nil {
		return fmt.Errorf("closing %s: not started", cl.image)
	}
	if err := cl.c.Close(); err != nil {
		return err
	}

	if deleteImage {
		output, err := exec.Command("docker", "rmi", "-f", cl.image).CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed removing image %s: %w (%s)", cl.image, err, string(output))
		}
	}

	return nil
}

func replacePlaceholders(arg string, env []servers.Env, secrets []servers.Secret) string {
	// TODO(dga): Temporary fix
	if arg == "{{filesystem.paths|volume-target|into}}" {
		return "."
	}

	for _, env := range env {
		if arg == "$"+env.Name {
			return fmt.Sprintf("%v", env.Example)
		}
	}
	for _, secret := range secrets {
		if arg == "$"+secret.Env {
			return secret.Example
		}
	}

	return arg
}

func toEnviron(env []servers.Env, secrets []servers.Secret) []string {
	var environ []string
	for _, env := range env {
		environ = append(environ, fmt.Sprintf("%s=%s", env.Name, env.Example))
	}
	for _, secret := range secrets {
		environ = append(environ, secret.Env+"="+secret.Example)
	}
	return environ
}

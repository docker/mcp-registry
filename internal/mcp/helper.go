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
	"slices"
	"sort"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/docker/mcp-registry/pkg/servers"
)

func Tools(ctx context.Context, server servers.Server, pull, cleanup, debug bool) ([]Tool, error) {
	var (
		args     []string
		extraEnv []servers.Env
	)
	if len(server.Requirement) > 0 {
		cancel, sidecarID, env, err := runRequirement(ctx, server.Requirement)
		if err != nil {
			return nil, err
		}
		defer cancel()

		for _, e := range env {
			parts := strings.SplitN(e, "=", 2)

			extraEnv = append(extraEnv, servers.Env{
				Name:    parts[0],
				Example: parts[1],
			})
		}

		args = append(args, "--network", "container:"+sidecarID)
	}

	env := append(server.Config.Env, extraEnv...)
	for name, value := range server.Run.Env {
		env = append(env, servers.Env{
			Name:  name,
			Value: value,
		})
	}

	c := newClient(server.Image, pull, env, server.Config.Secrets, args, server.Run.Command)
	if err := c.Start(ctx, debug); err != nil {
		return nil, err
	}

	tools, err := c.ListTools(ctx)
	if err != nil {
		c.Close(cleanup)
		return nil, err
	}

	err = c.Close(cleanup)
	if err != nil {
		return nil, err
	}

	sort.Slice(tools, func(i, j int) bool { return tools[i].Name < tools[j].Name })

	var list []Tool
	for _, tool := range tools {
		var arguments []ToolArgument
		var requiredPropertyNames []string
		var optionalPropertyNames []string
		for name := range tool.InputSchema.Properties {
			if slices.Contains(tool.InputSchema.Required, name) {
				requiredPropertyNames = append(requiredPropertyNames, name)
			} else {
				optionalPropertyNames = append(optionalPropertyNames, name)
			}
		}
		sort.Strings(requiredPropertyNames)
		sort.Strings(optionalPropertyNames)

		propertyNames := append(requiredPropertyNames, optionalPropertyNames...)

		for _, name := range propertyNames {
			v := tool.InputSchema.Properties[name]

			// Type
			argumentType := "string"
			rawType := v.(map[string]any)["type"]
			if rawType != "" && rawType != nil {
				if str, ok := rawType.(string); ok {
					argumentType = str
				}
			}

			// Item types
			var items *Items
			if argumentType == "array" {
				itemsType := "string"
				if rawItems, found := v.(map[string]any)["items"]; found {
					if kv, ok := rawItems.(map[string]any); ok {
						if rawItemsType, found := kv["type"]; found {
							if str, ok := rawItemsType.(string); ok {
								itemsType = str
							}
						}
					}
				}
				items = &Items{
					Type: itemsType,
				}
			}

			// Description
			desc := v.(map[string]any)["description"]

			// Properties
			arguments = append(arguments, ToolArgument{
				Name:        name,
				Type:        argumentType,
				Items:       items,
				Optional:    !slices.Contains(tool.InputSchema.Required, name),
				Description: argumentDescription(name, desc, tool.Description),
			})
		}

		// Annotations
		var annotations *ToolAnnotations
		if tool.Annotations != (mcp.ToolAnnotation{}) {
			annotations = &ToolAnnotations{
				Title:           tool.Annotations.Title,
				ReadOnlyHint:    tool.Annotations.ReadOnlyHint,
				DestructiveHint: tool.Annotations.DestructiveHint,
				IdempotentHint:  tool.Annotations.IdempotentHint,
				OpenWorldHint:   tool.Annotations.OpenWorldHint,
			}
		}

		list = append(list, Tool{
			Name:        tool.Name,
			Description: removeArgs(tool.Description),
			Arguments:   arguments,
			Annotations: annotations,
		})
	}

	return list, nil
}

func removeArgs(input string) string {
	var result []string

	for line := range strings.SplitSeq(input, "\n") {
		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(line)), "args:") {
			break
		}
		if strings.TrimSpace(line) == "" {
			result = append(result, "")
		} else {
			result = append(result, line)
		}
	}

	return strings.TrimSpace(strings.Join(result, "\n"))
}

func argumentDescription(name string, description any, toolDescription string) string {
	if description != nil && description != "" {
		return fmt.Sprintf("%s", description)
	}
	return extractDescription(toolDescription, name)
}

func extractDescription(input string, name string) string {
	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(strings.ToLower(line), name+":") {
			return strings.TrimSpace(strings.TrimPrefix(line, name+":"))
		}
	}

	return ""
}

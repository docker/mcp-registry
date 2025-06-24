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

import "encoding/json"

type RPCResponse struct {
	Error    *string
	Response *json.RawMessage
}

type BaseMessage struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      *int64          `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type Tool struct {
	Name        string           `json:"name" yaml:"name"`
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	Arguments   []ToolArgument   `json:"arguments,omitempty" yaml:"arguments,omitempty"`
	Annotations *ToolAnnotations `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

type ToolArgument struct {
	Name        string `json:"name" yaml:"name"`
	Type        string `json:"type" yaml:"type"`
	Items       *Items `json:"items,omitempty" yaml:"items,omitempty"`
	Description string `json:"desc" yaml:"desc"`
	Optional    bool   `json:"optional,omitempty" yaml:"optional,omitempty"`
}

type ToolAnnotations struct {
	Title           string `json:"title,omitempty" yaml:"title,omitempty"`
	ReadOnlyHint    bool   `json:"readOnlyHint,omitempty" yaml:"readOnlyHint,omitempty"`
	DestructiveHint bool   `json:"destructiveHint,omitempty" yaml:"destructiveHint,omitempty"`
	IdempotentHint  bool   `json:"idempotentHint,omitempty" yaml:"idempotentHint,omitempty"`
	OpenWorldHint   bool   `json:"openWorldHint,omitempty" yaml:"openWorldHint,omitempty"`
}

type Items struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

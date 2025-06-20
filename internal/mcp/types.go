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

package main

import (
	"testing"
)

func Test_isNameValid(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "valid name",
			args: args{
				name: "my-server",
			},
			wantError: false,
		},
		{
			name: "invalid name",
			args: args{
				name: "My-Server",
			},
			wantError: true,
		},
		{
			name: "valid name with numbers",
			args: args{
				name: "my-server-1",
			},
			wantError: false,
		},
		{
			name: "invalid name with symbol",
			args: args{
				name: "my-server-$",
			},
			wantError: true,
		},
		{
			name: "invalid name with space",
			args: args{
				name: "my server",
			},
			wantError: true,
		},
		{
			name: "invalid name with slash",
			args: args{
				name: "my-server/1",
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNameValid(tt.args.name); (got != nil) != tt.wantError {
				t.Errorf("isNameValid() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func Test_areSecretsValid(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "valid secrets",
			args: args{
				name: "astra-db",
			},
			wantError: false,
		},
		{
			name: "no secrets",
			args: args{
				name: "arxiv-mcp-server",
			},
			wantError: false,
		},
		{
			name: "invalid secrets",
			args: args{
				name: "bad-server",
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSecretsValid(tt.args.name); (got != nil) != tt.wantError {
				t.Errorf("areSecretsValid() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func TestValidateDockerfilePath(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
		desc    string
	}{
		{"Dockerfile", false, "simple Dockerfile"},
		{"src/postgres/Dockerfile", false, "relative path with subdirectory"},
		{"Dockerfile.local", false, "Dockerfile with suffix"},
		{"/Dockerfile", true, "absolute path should be rejected"},
		{"../Dockerfile", true, "directory traversal should be rejected"},
		{"src/../Dockerfile", true, "directory traversal in middle should be rejected"},
		{"Dockerfile\x00", true, "null byte should be rejected"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := validateDockerfilePath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDockerfilePath(%q) error = %v, wantErr %v", tt.path, err, tt.wantErr)
			}
		})
	}
}

func TestValidateDockerfileContent(t *testing.T) {
	tests := []struct {
		content string
		wantErr bool
		desc    string
	}{
		{
			content: "FROM node:18-alpine\nWORKDIR /app\nCMD [\"node\", \"index.js\"]",
			wantErr: false,
			desc:    "valid simple Dockerfile",
		},
		{
			content: "FROM alpine\nRUN apk add --no-cache curl\nENTRYPOINT [\"curl\"]",
			wantErr: false,
			desc:    "valid Dockerfile with ENTRYPOINT",
		},
		{
			content: "# Comment only Dockerfile\n# Another comment",
			wantErr: true,
			desc:    "Dockerfile with only comments should fail",
		},
		{
			content: "",
			wantErr: true,
			desc:    "empty Dockerfile should fail",
		},
		{
			content: "RUN echo hello\nFROM alpine",
			wantErr: true,
			desc:    "Dockerfile not starting with FROM should fail",
		},
		{
			content: "FROM node:18\nWORKDIR /app",
			wantErr: true,
			desc:    "Dockerfile without CMD or ENTRYPOINT should fail",
		},
		{
			content: "FROM alpine\nENV PASSWORD=secret123\nCMD [\"sh\"]",
			wantErr: true,
			desc:    "Dockerfile with hardcoded password should fail",
		},
		{
			content: "FROM alpine\nENV API_KEY=abc123\nCMD [\"sh\"]",
			wantErr: true,
			desc:    "Dockerfile with hardcoded API key should fail",
		},
		{
			content: "FROM alpine\nENV SECRET=mysecret\nENTRYPOINT [\"sh\"]",
			wantErr: true,
			desc:    "Dockerfile with hardcoded secret should fail",
		},
		{
			content: "FROM alpine\nENV TOKEN=mytoken\nCMD [\"sh\"]",
			wantErr: true,
			desc:    "Dockerfile with hardcoded token should fail",
		},
		{
			content: "FROM node:18\n# This is a comment\nWORKDIR /app\n# Another comment\nCMD [\"node\", \"app.js\"]",
			wantErr: false,
			desc:    "valid Dockerfile with comments should pass",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := validateDockerfileContent(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDockerfileContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

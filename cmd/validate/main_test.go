package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// Compute the path to this source code file.
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintln(os.Stderr, "mcp-registry/cmd/validate: unable to resolve caller path")
		os.Exit(1)
	}

	// Switch to the repository root so that readServerYaml calls from tests can
	// access YAML files.
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", ".."))
	if err := os.Chdir(repoRoot); err != nil {
		fmt.Fprintln(os.Stderr, "mcp-registry/cmd/validate: chdir:", err)
		os.Exit(1)
	}

	// Run the tests in this package.
	code := m.Run()

	// Restore the working directory.
	originalWD := filepath.Clean(filepath.Join(repoRoot, "cmd", "validate"))
	if err := os.Chdir(originalWD); err != nil {
		fmt.Fprintln(os.Stderr, "mcp-registry/cmd/validate: restore chdir:", err)
		os.Exit(1)
	}

	os.Exit(code)
}

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
		{
			name: "legacy uppercase name",
			args: args{
				name: "SQLite",
			},
			wantError: false,
		},
		{
			name: "legacy underscore name",
			args: args{
				name: "youtube_transcript",
			},
			wantError: false,
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

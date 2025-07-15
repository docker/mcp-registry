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

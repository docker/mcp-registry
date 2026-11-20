package mcp

import (
	"testing"

	"github.com/docker/mcp-registry/pkg/servers"
)

func TestReplacePlaceholders(t *testing.T) {
	params := servers.Schema{
		Properties: servers.SchemaList{
			{Name: "access_level", Schema: servers.Schema{Type: "string", Default: "readonly"}},
			{Name: "retries", Schema: servers.Schema{Type: "integer", Default: 3}},
			{Name: "empty", Schema: servers.Schema{Type: "string", Default: ""}},
			{Name: "optional", Schema: servers.Schema{Type: "string"}},
		},
	}
	env := []servers.Env{{Name: "ENV_NAME", Example: "example-value"}}
	secrets := []servers.Secret{{Env: "SECRET_NAME", Example: "secret-example"}}

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"parameter in flag", "--access-level={{aks.access_level}}", "--access-level=readonly"},
		{"multiple parameters", "{{aks.access_level}}-{{aks.retries}}", "readonly-3"},
		{"empty default", "--filter={{aks.empty}}", "--filter="},
		{"unconfigured parameter", "--optional={{aks.optional}}", "--optional={{aks.optional}}"},
		{"other server parameter", "{{other.access_level}}", "{{other.access_level}}"},
		{"environment example", "$ENV_NAME", "example-value"},
		{"secret example", "$SECRET_NAME", "secret-example"},
		{"filesystem fallback", "{{filesystem.paths|volume-target|into}}", "."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replacePlaceholders(tt.arg, env, secrets, "aks", params)
			if got != tt.want {
				t.Errorf("replacePlaceholders(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}

func TestReplacePlaceholdersFromServerYAML(t *testing.T) {
	server, err := servers.Read("../../servers/aks/server.yaml")
	if err != nil {
		t.Fatal(err)
	}

	got := replacePlaceholders(server.Run.Command[1], server.Config.Env, server.Config.Secrets, server.Name, server.Config.Parameters)
	if got != "--access-level=readonly" {
		t.Errorf("startup argument = %q, want %q", got, "--access-level=readonly")
	}
}

package catalog

import (
	"context"
	"testing"

	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

func TestHostedServiceAttribution(t *testing.T) {
	tests := []struct {
		name, kind, url, metadata, owner, license string
	}{
		{"explicit hosted attribution", "remote", "https://example.com/mcp", "  owner: Example Company\n  license: Proprietary (hosted service)\n", "Example Company", "Proprietary (hosted service)"},
		{"legacy remote defaults", "remote", "https://example.com/mcp", "", "docker", "Apache License 2.0"},
		{"owner only", "remote", "https://example.com/mcp", "  owner: Example Company\n", "Example Company", "Apache License 2.0"},
		{"license only", "remote", "https://example.com/mcp", "  license: Proprietary\n", "docker", "Proprietary"},
		{"blank values preserve defaults", "remote", "https://example.com/mcp", "  owner: '  '\n  license: '  '\n", "docker", "Apache License 2.0"},
		{"legacy hosted server type", "server", "https://example.com/mcp", "  owner: Example Company\n  license: Proprietary\n", "Example Company", "Proprietary"},
		{"nonremote ignores overrides", "poci", "", "  owner: Example Company\n  license: Proprietary\n", "docker", "Apache License 2.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the same metadata format as a contributed server.yaml.
			input := "name: example\ntype: " + tt.kind + "\nabout:\n  title: Example\n  description: Example service\nremote:\n  url: " + tt.url + "\nmeta:\n  category: productivity\n" + tt.metadata
			var server servers.Server
			if err := yaml.Unmarshal([]byte(input), &server); err != nil {
				t.Fatal(err)
			}
			tile, err := ToTile(context.Background(), server)
			if err != nil {
				t.Fatal(err)
			}
			if tile.Metadata.Owner != tt.owner || tile.Metadata.License != tt.license {
				t.Fatalf("attribution = %q / %q, want %q / %q", tile.Metadata.Owner, tile.Metadata.License, tt.owner, tt.license)
			}
		})
	}
}

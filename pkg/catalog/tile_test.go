package catalog

import (
	"context"
	"testing"

	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

func TestRemoteServiceMetadataFromYAML(t *testing.T) {
	cases := []struct {
		name, metadata, owner, license string
	}{
		{"explicit hosted service", "  owner: ForUsApps, LLC\n  license: LicenseRef-Proprietary\n", "ForUsApps, LLC", "LicenseRef-Proprietary"},
		{"existing entry without overrides", "", "docker", "Apache License 2.0"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := "name: hosted-example\ntype: remote\nmeta:\n  category: productivity\n" + tc.metadata + "about:\n  title: Hosted Example\n  description: A hosted service\nremote:\n  transport_type: streamable-http\n  url: https://example.invalid/mcp\n"
			var server servers.Server
			if err := yaml.Unmarshal([]byte(input), &server); err != nil {
				t.Fatal(err)
			}
			tile, err := ToTile(context.Background(), server)
			if err != nil {
				t.Fatal(err)
			}
			if tile.Metadata.Owner != tc.owner || tile.Metadata.License != tc.license {
				t.Fatalf("publisher/license = %q / %q, want %q / %q", tile.Metadata.Owner, tile.Metadata.License, tc.owner, tc.license)
			}
			if tile.Remote.URL != "https://example.invalid/mcp" || tile.Remote.TransportType != "streamable-http" {
				t.Fatal("metadata must preserve the remote connection")
			}
		})
	}
}

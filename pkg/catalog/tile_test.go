package catalog

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/docker/mcp-registry/pkg/servers"
)

func TestToTilePreservesLongLived(t *testing.T) {
	tile, err := ToTile(context.Background(), servers.Server{
		Name:      "long-lived-server",
		Type:      "server",
		LongLived: true,
		About: servers.About{
			Title:       "Long-lived server",
			Description: "A long-lived server",
		},
		Remote: servers.Remote{
			TransportType: "streamable-http",
			URL:           "https://example.com/mcp",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !tile.LongLived {
		t.Error("ToTile() did not preserve LongLived")
	}
}

func TestWriteYamlPreservesLongLived(t *testing.T) {
	catalogFile := filepath.Join(t.TempDir(), "catalog.yaml")
	if err := WriteYaml(catalogFile, TopLevel{
		Version:     Version,
		Name:        Name,
		DisplayName: DisplayName,
		Registry: TileList{
			{
				Name: "long-lived-server",
				Tile: Tile{LongLived: true},
			},
		},
	}); err != nil {
		t.Fatal(err)
	}

	contents, err := os.ReadFile(catalogFile)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(contents), "longLived: true") {
		t.Errorf("WriteYaml() output does not contain longLived: true:\n%s", contents)
	}
}

package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

type tileRoundTripper func(*http.Request) (*http.Response, error)

func (f tileRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// All HTTP in these tests is handled in process; unexpected requests fail closed.
func mockTileHTTP(t *testing.T, f tileRoundTripper) {
	t.Helper()
	t.Setenv("GITHUB_TOKEN", "")
	previous := http.DefaultTransport
	http.DefaultTransport = f
	t.Cleanup(func() { http.DefaultTransport = previous })
}

func TestToTileRemoteMetadata(t *testing.T) {
	cases := []struct {
		name, kind, transport, project string
		oauth                          bool
	}{
		{name: "anonymous-http", kind: "remote", transport: "streamable-http"},
		{name: "oauth-sse", kind: "remote", transport: "sse", oauth: true},
		{name: "legacy-server-type", kind: "server", transport: "streamable-http"},
		{name: "remote-with-project", kind: "remote", transport: "streamable-http", project: "https://github.com/example/remote"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			requests := 0
			mockTileHTTP(t, func(r *http.Request) (*http.Response, error) {
				requests++
				return nil, fmt.Errorf("unexpected HTTP request: %s", r.URL)
			})
			server := servers.Server{
				Name: "remote-fixture", Type: tc.kind,
				About:  servers.About{Title: "Remote fixture", Description: "Public documentation"},
				Meta:   servers.Meta{Category: "documentation", Tags: []string{"remote", "documentation"}},
				Source: servers.Source{Project: tc.project},
				Remote: servers.Remote{TransportType: tc.transport, URL: "https://example.invalid/mcp", Headers: map[string]string{"X-Project": "{{remote-fixture.project}}"}},
			}
			if tc.oauth {
				server.OAuth = []servers.OAuthProvider{{Provider: "fixture", Secret: "fixture.token", Env: "FIXTURE_TOKEN"}}
			}
			tile, err := ToTile(context.Background(), server)
			if err != nil {
				t.Fatal(err)
			}
			if tile.Metadata.Owner != "" || tile.Metadata.License != "" {
				t.Errorf("remote endpoint acquired unverified metadata: owner=%q license=%q", tile.Metadata.Owner, tile.Metadata.License)
			}
			if requests != 0 {
				t.Errorf("remote catalog generation made %d HTTP requests", requests)
			}
			if tile.Remote.URL != server.Remote.URL || tile.Remote.TransportType != tc.transport || !reflect.DeepEqual(tile.Remote.Headers, server.Remote.Headers) {
				t.Errorf("remote configuration changed: %#v", tile.Remote)
			}
			if tile.Metadata.Category != server.Meta.Category || !reflect.DeepEqual(tile.Metadata.Tags, server.Meta.Tags) {
				t.Errorf("category/tags changed: %#v", tile.Metadata)
			}
			if tile.Source != "" || tile.Upstream != tc.project {
				t.Errorf("source behavior changed: source=%q upstream=%q", tile.Source, tile.Upstream)
			}
			if len(tile.OAuth.Providers) != len(server.OAuth) {
				t.Fatalf("OAuth provider count changed")
			}
			if tc.oauth && tile.OAuth.Providers[0] != (OAuthProvider{Provider: "fixture", Secret: "fixture.token", Env: "FIXTURE_TOKEN"}) {
				t.Errorf("OAuth configuration changed: %#v", tile.OAuth)
			}

			// Exercise the real catalog serializers, including TileList's custom encoding.
			catalog := TopLevel{Version: Version, Name: Name, DisplayName: DisplayName, Registry: TileList{{Name: server.Name, Tile: tile}}}
			encoded, err := json.Marshal(catalog)
			if err != nil {
				t.Fatal(err)
			}
			var jsonDocument map[string]any
			if err := json.Unmarshal(encoded, &jsonDocument); err != nil {
				t.Fatal(err)
			}
			assertRemoteSerializedMetadata(t, jsonDocument, server.Name)
			path := filepath.Join(t.TempDir(), "catalog.yaml")
			if err := WriteYaml(path, catalog); err != nil {
				t.Fatal(err)
			}
			encoded, err = os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var yamlDocument map[string]any
			if err := yaml.Unmarshal(encoded, &yamlDocument); err != nil {
				t.Fatal(err)
			}
			assertRemoteSerializedMetadata(t, yamlDocument, server.Name)
		})
	}
}

func assertRemoteSerializedMetadata(t *testing.T, document map[string]any, name string) {
	t.Helper()
	entry := document["registry"].(map[string]any)[name].(map[string]any)
	metadata := entry["metadata"].(map[string]any)
	for _, field := range []string{"owner", "license"} {
		if value, present := metadata[field]; present {
			t.Errorf("serialized remote metadata still has %s=%v", field, value)
		}
	}
	if metadata["category"] != "documentation" {
		t.Errorf("serialized category changed: %v", metadata)
	}
	if entry["remote"].(map[string]any)["url"] != "https://example.invalid/mcp" {
		t.Error("serialized remote URL changed")
	}
}

func TestToTileLocalProjectMetadata(t *testing.T) {
	for _, license := range []string{"MIT License", "Apache License 2.0"} {
		t.Run(license, func(t *testing.T) {
			requests := 0
			mockTileHTTP(t, func(r *http.Request) (*http.Response, error) {
				requests++
				if r.Method != "GET" || r.URL.String() != "https://api.github.com/repos/example/local" || r.Header.Get("Authorization") != "" {
					return nil, fmt.Errorf("unexpected request: %s %s", r.Method, r.URL)
				}
				payload, err := json.Marshal(map[string]any{"owner": map[string]string{"login": "fixture-owner"}, "license": map[string]string{"key": "fixture", "name": license}, "stargazers_count": 7, "description": "Repository description"})
				if err != nil {
					return nil, err
				}
				return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(string(payload))), Request: r}, nil
			})
			server := servers.Server{Name: "local-fixture", Type: "server", Image: "example/local:latest", About: servers.About{Title: "Local fixture"}, Source: servers.Source{Project: "https://github.com/example/local", Commit: "abc123"}}
			tile, err := ToTile(context.Background(), server)
			if err != nil {
				t.Fatal(err)
			}
			if requests != 1 {
				t.Errorf("expected one mocked repository lookup, got %d", requests)
			}
			if tile.Metadata.Owner != "fixture-owner" || tile.Metadata.License != license || tile.Metadata.GitHubStars != 7 {
				t.Errorf("project metadata changed: %#v", tile.Metadata)
			}
			if tile.Source != "https://github.com/example/local/tree/abc123" || tile.Description != "Repository description." || tile.Image != server.Image {
				t.Errorf("local source/description/image changed: %#v", tile)
			}
		})
	}
}

func TestToTilePOCIDefaultMetadata(t *testing.T) {
	mockTileHTTP(t, func(r *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("unexpected HTTP request: %s", r.URL)
	})
	tile, err := ToTile(context.Background(), servers.Server{Name: "poci-fixture", Type: "poci", About: servers.About{Title: "POCI fixture", Description: "Local fixture"}})
	if err != nil {
		t.Fatal(err)
	}
	if tile.Metadata.Owner != "docker" || tile.Metadata.License != "Apache License 2.0" {
		t.Errorf("non-remote default metadata changed: %#v", tile.Metadata)
	}
}

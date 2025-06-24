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

package catalog

import (
	"bytes"
	"encoding/json"

	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

const (
	Version     = 2
	Name        = "docker-mcp"
	DisplayName = "Docker MCP Catalog"
)

type TileWithOrder struct {
	Tile  `json:",inline" yaml:",inline"`
	Order int `json:"order" yaml:"order"`
}

type TopLevel struct {
	Version     int      `json:"version" yaml:"version"`
	Name        string   `json:"name" yaml:"name"`
	DisplayName string   `json:"displayName" yaml:"displayName"`
	Registry    TileList `json:"registry" yaml:"registry"`
}

type TileList []TileEntry

func (tl *TileList) UnmarshalYAML(value *yaml.Node) error {
	for i := 0; i < len(value.Content); i += 2 {
		keyNode := value.Content[i]
		valNode := value.Content[i+1]

		var name string
		if err := keyNode.Decode(&name); err != nil {
			return err
		}

		var tile Tile
		if err := valNode.Decode(&tile); err != nil {
			return err
		}

		*tl = append(*tl, TileEntry{
			Name: name,
			Tile: tile,
		})
	}
	return nil
}

func (tl TileList) MarshalYAML() (interface{}, error) {
	mapNode := &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: []*yaml.Node{},
	}

	for _, entry := range tl {
		// Key node: the tile name
		keyNode := &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: entry.Name,
		}

		// Value node: marshal the Tile
		valNode := &yaml.Node{}
		if err := valNode.Encode(entry.Tile); err != nil {
			return nil, err
		}

		mapNode.Content = append(mapNode.Content, keyNode, valNode)
	}

	return mapNode, nil
}

func (tl TileList) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, entry := range tl {
		if i > 0 {
			buf.WriteByte(',')
		}
		keyBytes, err := json.Marshal(entry.Name)
		if err != nil {
			return nil, err
		}
		valBytes, err := json.Marshal(entry.Tile)
		if err != nil {
			return nil, err
		}
		buf.Write(keyBytes)
		buf.WriteByte(':')
		buf.Write(valBytes)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

type TileEntry struct {
	Name string `json:"name" yaml:"name"`
	Tile Tile   `json:",inline" yaml:",inline"`
}

type Tile struct {
	Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
	Description string  `json:"description" yaml:"description"`
	Title       string  `json:"title" yaml:"title"`
	Type        string  `json:"type" yaml:"type"`
	DateAdded   *string `json:"dateAdded,omitempty" yaml:"dateAdded,omitempty"`
	Image       string  `json:"image,omitempty" yaml:"image,omitempty"`
	// TODO(dga): Remove it when the UI is ready. It's not used but it's still validated.
	Ref       string `json:"ref" yaml:"ref"`
	ReadmeURL string `json:"readme,omitempty" yaml:"readme,omitempty"`
	ToolsURL  string `json:"toolsUrl,omitempty" yaml:"toolsUrl,omitempty"`
	// TODO(dga): The UI ignores tiles without a source. An empty one is ok. Put back omitempty when this is fixed
	// Source         string         `json:"source,omitempty" yaml:"source,omitempty"`
	Source         string         `json:"source" yaml:"source"`
	Upstream       string         `json:"upstream,omitempty" yaml:"upstream,omitempty"`
	Icon           string         `json:"icon" yaml:"icon"`
	Tools          []servers.Tool `json:"tools" yaml:"tools"`
	Secrets        []Secret       `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	Env            []Env          `json:"env,omitempty" yaml:"env,omitempty"`
	Command        []string       `json:"command,omitempty" yaml:"command,omitempty"`
	Volumes        []string       `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	DisableNetwork bool           `json:"disableNetwork,omitempty" yaml:"disableNetwork,omitempty"`
	AllowHosts     []string       `json:"allowHosts,omitempty" yaml:"allowHosts,omitempty"`
	Prompts        int            `json:"prompts" yaml:"prompts"`
	Resources      map[string]any `json:"resources" yaml:"resources"`
	Config         []Config       `json:"config,omitempty" yaml:"config,omitempty"`
	Metadata       Metadata       `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	OAuth          OAuth          `json:"oauth,omitempty" yaml:"oauth,omitempty"`
}

type Metadata struct {
	Pulls       int      `json:"pulls,omitempty" yaml:"pulls,omitempty"`
	Stars       int      `json:"stars,omitempty" yaml:"stars,omitempty"`
	GitHubStars int      `json:"githubStars,omitempty" yaml:"githubStars,omitempty"`
	Category    string   `json:"category,omitempty" yaml:"category,omitempty"`
	Tags        []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	License     string   `json:"license,omitempty" yaml:"license,omitempty"`
	Owner       string   `json:"owner,omitempty" yaml:"owner,omitempty"`
}

type OAuth struct {
	Providers []OAuthProvider `json:"providers,omitempty" yaml:"providers,omitempty"`
}

type OAuthProvider struct {
	Provider string `json:"provider" yaml:"provider"`
	Secret   string `json:"secret,omitempty" yaml:"secret,omitempty"`
	Env      string `json:"env,omitempty" yaml:"env,omitempty"`
}

type Config struct {
	Name        string             `json:"name" yaml:"name"`
	Description string             `json:"description" yaml:"description"`
	Type        string             `json:"type" yaml:"type"`
	Properties  servers.SchemaList `json:"properties,omitempty" yaml:"properties,omitempty"`
	AnyOf       []servers.AnyOf    `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Required    []string           `json:"required,omitempty" yaml:"required,omitempty"`
}

type Secret struct {
	Name     string `json:"name" yaml:"name"`
	Env      string `json:"env" yaml:"env"`
	Example  string `json:"example" yaml:"example"`
	Required bool   `json:"required,omitempty" yaml:"required,omitempty"`
}

type Env struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

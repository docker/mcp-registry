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

package servers

import (
	"gopkg.in/yaml.v3"
)

type Server struct {
	Name        string          `yaml:"name" json:"name"`
	Image       string          `yaml:"image,omitempty" json:"image,omitempty"`
	Type        string          `yaml:"type" json:"type"`
	LongLived   bool            `yaml:"longLived,omitempty" json:"longLived,omitempty"`
	Meta        Meta            `yaml:"meta,omitempty" json:"meta,omitempty"`
	About       About           `yaml:"about,omitempty" json:"about,omitempty"`
	Source      Source          `yaml:"source,omitempty" json:"source,omitempty"`
	Run         Run             `yaml:"run,omitempty" json:"run,omitempty"`
	Config      Config          `yaml:"config,omitempty" json:"config,omitempty"`
	OAuth       []OAuthProvider `yaml:"oauth,omitempty" json:"oauth,omitempty"`
	Tools       []Tool          `yaml:"tools,omitempty" json:"tools,omitempty"`
	Requirement string          `yaml:"requirement,omitempty" json:"requirement,omitempty"`
}

type Secret struct {
	Name     string `yaml:"name" json:"name"`
	Env      string `yaml:"env" json:"env"`
	Example  string `yaml:"example,omitempty" json:"example,omitempty"`
	Required *bool  `yaml:"required,omitempty" json:"required,omitempty"`
}

type Env struct {
	Name    string `yaml:"name" json:"name"`
	Example any    `yaml:"example,omitempty" json:"example,omitempty"`
	Value   string `yaml:"value,omitempty" json:"value,omitempty"`
}

type AnyOf struct {
	Required []string `yaml:"required,omitempty" json:"required,omitempty"`
}

type Schema struct {
	Type        string     `yaml:"type" json:"type"`
	Description string     `yaml:"description,omitempty" json:"description,omitempty"`
	Properties  SchemaList `yaml:"properties,omitempty" json:"properties,omitempty"`
	Required    []string   `yaml:"required,omitempty" json:"required,omitempty"`
	Items       Items      `yaml:"items,omitempty" json:"items,omitempty"`
	AnyOf       []AnyOf    `yaml:"anyOf,omitempty" json:"anyOf,omitempty"`
	Default     any        `yaml:"default,omitempty" json:"default,omitempty"`
}

type Items struct {
	Type string `yaml:"type" json:"type"`
}

type About struct {
	Title       string `yaml:"title,omitempty" json:"title,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	Icon        string `yaml:"icon,omitempty" json:"icon,omitempty"`
}

type Source struct {
	Project    string `yaml:"project,omitempty" json:"project,omitempty"`
	Upstream   string `yaml:"upstream,omitempty" json:"upstream,omitempty"`
	Branch     string `yaml:"branch,omitempty" json:"branch,omitempty"`
	Directory  string `yaml:"directory,omitempty" json:"directory,omitempty"`
	Dockerfile string `yaml:"dockerfile,omitempty" json:"dockerfile,omitempty"`
}

type Run struct {
	Command        []string          `yaml:"command,omitempty" json:"command,omitempty"`
	Volumes        []string          `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	Env            map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
	AllowHosts     []string          `yaml:"allowHosts,omitempty" json:"allowHosts,omitempty"`
	DisableNetwork bool              `yaml:"disableNetwork,omitempty" json:"disableNetwork,omitempty"`
}

type Config struct {
	Description string   `yaml:"description,omitempty" json:"description,omitempty"`
	Secrets     []Secret `yaml:"secrets,omitempty" json:"secrets,omitempty"`
	Env         []Env    `yaml:"env,omitempty" json:"env,omitempty"`
	Parameters  Schema   `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	AnyOf       []AnyOf  `yaml:"anyOf,omitempty" json:"anyOf,omitempty"`
}

type Tool struct {
	Name        string     `yaml:"name" json:"name"`
	Description string     `yaml:"description,omitempty" json:"description,omitempty"`
	Parameters  Parameters `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	Container   Container  `yaml:"container,omitempty" json:"container,omitempty"`
}

type Parameters struct {
	Type       string     `yaml:"type" json:"type"`
	Properties Properties `yaml:"properties" json:"properties"`
	Required   []string   `yaml:"required" json:"required"`
}

type Properties map[string]Property

type Property struct {
	Type        string `yaml:"type" json:"type"`
	Description string `yaml:"description" json:"description"`
	Items       *Items `yaml:"items,omitempty" json:"items,omitempty"`
}

type Container struct {
	Image   string   `yaml:"image,omitempty" json:"image,omitempty"`
	Command []string `yaml:"command,omitempty" json:"command,omitempty"`
	Volumes []string `yaml:"volumes,omitempty" json:"volumes,omitempty"`
}

type Meta struct {
	Category    string   `yaml:"category,omitempty" json:"category,omitempty"`
	Tags        []string `yaml:"tags,omitempty" json:"tags,omitempty"`
	Highlighted bool     `yaml:"highlighted,omitempty" json:"highlighted,omitempty"`
}

type OAuthProvider struct {
	Provider string `yaml:"provider,omitempty" json:"provider,omitempty"`
	Secret   string `yaml:"secret,omitempty" json:"secret,omitempty"`
	Env      string `yaml:"env,omitempty" json:"env,omitempty"`
}

type SchemaEntry struct {
	Schema Schema `yaml:",inline"`
	Name   string `yaml:"name"`
}

type SchemaList []SchemaEntry

func (tl *SchemaList) UnmarshalYAML(value *yaml.Node) error {
	for i := 0; i < len(value.Content); i += 2 {
		keyNode := value.Content[i]
		valNode := value.Content[i+1]

		var name string
		if err := keyNode.Decode(&name); err != nil {
			return err
		}

		var schema Schema
		if err := valNode.Decode(&schema); err != nil {
			return err
		}

		*tl = append(*tl, SchemaEntry{
			Name:   name,
			Schema: schema,
		})
	}
	return nil
}

func (tl SchemaList) MarshalYAML() (interface{}, error) {
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

		// Value node: marshal the Schema
		valNode := &yaml.Node{}
		if err := valNode.Encode(entry.Schema); err != nil {
			return nil, err
		}

		mapNode.Content = append(mapNode.Content, keyNode, valNode)
	}

	return mapNode, nil
}

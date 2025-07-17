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
	"strings"
)

func (s *Server) GetContext() string {
	base := s.Source.Project + ".git"

	if s.GetBranch() != "main" {
		base += "#" + s.Source.Branch
	} else {
		base += "#"
	}

	if s.Source.Directory != "" && s.Source.Directory != "." {
		base += ":" + s.Source.Directory
	}

	return strings.TrimSuffix(base, "#")
}

func (s *Server) GetSourceURL() string {
	source := s.Source.Project + "/tree/" + s.GetBranch()
	if s.Source.Directory != "" {
		source += "/" + s.Source.Directory
	}
	return source
}

func (s *Server) GetUpstream() string {
	if s.Source.Upstream != "" {
		return s.Source.Upstream
	}
	return s.Source.Project
}

func (s *Server) GetBranch() string {
	if s.Source.Branch == "" {
		return "main"
	}
	return s.Source.Branch
}

func (s *Server) GetDockerfileUrl() string {
	base := s.Source.Project + "/blob/" + s.GetBranch()
	if s.Source.Directory != "" {
		base += "/" + s.Source.Directory
	}
	return base + "/" + s.GetDockerfile()
}

func (s *Server) GetDockerfile() string {
	if s.Source.Dockerfile == "" {
		return "Dockerfile"
	}
	return s.Source.Dockerfile
}

func CreateSchema(server string, env []Env) ([]Env, Schema) {
	schema := Schema{}
	if len(env) == 0 {
		return nil, schema
	}

	var updatedEnv []Env
	schema.Type = "object"
	for _, e := range env {

		propertyName := strings.ToLower(e.Name)

		schema.Properties = append(schema.Properties, SchemaEntry{
			Name: propertyName,
			Schema: Schema{
				Type: "string",
			},
		})

		updatedEnv = append(updatedEnv, Env{
			Name:    e.Name,
			Value:   "{{" + server + "." + propertyName + "}}",
			Example: e.Example,
		})
	}

	return updatedEnv, schema
}

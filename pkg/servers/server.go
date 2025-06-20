package servers

import "strings"

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

		name := strings.TrimPrefix(strings.ToLower(e.Name), strings.ToLower(server)+"_")

		schema.Properties = append(schema.Properties, SchemaEntry{
			Name: name,
			Schema: Schema{
				Type: "string (TODO)",
			},
		})

		updatedEnv = append(updatedEnv, Env{
			Name:    e.Name,
			Value:   "{{" + name + "}}",
			Example: e.Example,
		})
	}

	return updatedEnv, schema
}

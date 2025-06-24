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
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/docker/mcp-registry/internal/licenses"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"

	"github.com/docker/mcp-registry/pkg/hub"
)

func ToTile(ctx context.Context, server servers.Server) (Tile, error) {
	description := server.About.Description
	source := ""
	upstream := server.GetUpstream()
	owner := "docker"
	license := "Apache License 2.0"
	githubStars := 0

	if server.Type == "server" {
		client := github.NewFromServer(server)
		repository, err := client.GetProjectRepository(ctx, upstream)
		if err != nil {
			return Tile{}, err
		}
		if !licenses.IsValid(repository.License) {
			panic(fmt.Sprintf("Project %s is licensed under %s which may be incompatible with some tools", upstream, repository.License))
		}

		if description == "" {
			description = repository.GetDescription()
		}
		source = server.GetSourceURL()
		owner = repository.Owner.GetLogin()
		license = repository.License.GetName()
		githubStars = repository.GetStargazersCount()
	}

	if description == "" {
		return Tile{}, fmt.Errorf("no description found for: %s", server.Name)
	}

	var secrets []Secret
	for _, s := range server.Config.Secrets {
		required := false
		if s.Required != nil {
			required = *s.Required
		}

		secrets = append(secrets, Secret{
			Name:     s.Name,
			Env:      s.Env,
			Example:  s.Example,
			Required: required,
		})
	}

	var env []Env
	for _, e := range server.Config.Env {
		env = append(env, Env{
			Name:  e.Name,
			Value: e.Value,
		})
	}
	for name, value := range server.Run.Env {
		env = append(env, Env{
			Name:  name,
			Value: value,
		})
	}

	var config []Config
	if len(server.Config.Parameters.Properties) > 0 {
		if server.Config.Description == "" {
			return Tile{}, fmt.Errorf("no config description found for: %s", server.Name)
		}

		catalogConfig := Config{
			Name:        server.Name,
			Description: server.Config.Description,
			Type:        server.Config.Parameters.Type,
			Properties:  server.Config.Parameters.Properties,
			Required:    server.Config.Parameters.Required,
			AnyOf:       server.Config.Parameters.AnyOf,
		}
		for i, property := range server.Config.Parameters.Properties {
			property.Schema.Default = nil
			server.Config.Parameters.Properties[i] = property
			if property.Schema.Type == "" {
				panic("no type found for: " + property.Name + " in " + server.Name)
			}
		}

		config = append(config, catalogConfig)
	}

	if server.About.Title == "" {
		return Tile{}, fmt.Errorf("no title found for: %s", server.Name)
	}

	image := server.Image

	if server.Type == "server" && image == "" {
		return Tile{}, fmt.Errorf("no image for server: %s", server.Name)
	}
	if server.Type == "poci" && image != "" {
		return Tile{}, fmt.Errorf("pocis don't have images: %s", server.Name)
	}

	pullCount := 0
	starCount := 0
	if strings.HasPrefix(image, "mcp/") {
		repoInfo, err := hub.GetRepositoryInfo(ctx, server.Image)
		if err != nil {
			return Tile{}, err
		}
		pullCount = repoInfo.PullCount
		starCount = repoInfo.StarCount
	}

	meta := Metadata{
		Category:    server.Meta.Category,
		Tags:        server.Meta.Tags,
		Owner:       owner,
		License:     license,
		GitHubStars: githubStars,
		Pulls:       pullCount,
		Stars:       starCount,
	}

	var oauth OAuth
	if len(server.OAuth) > 0 {
		for _, provider := range server.OAuth {
			oauth.Providers = append(oauth.Providers, OAuthProvider{
				Provider: provider.Provider,
				Secret:   provider.Secret,
				Env:      provider.Env,
			})
		}
	}

	dateAdded := time.Now().Format(time.RFC3339)

	return Tile{
		Description:    addDot(strings.TrimSpace(strings.ReplaceAll(description, "\n", " "))),
		Title:          server.About.Title,
		Type:           server.Type,
		Image:          image,
		DateAdded:      &dateAdded,
		ReadmeURL:      "http://desktop.docker.com/mcp/catalog/v" + strconv.Itoa(Version) + "/readme/" + server.Name + ".md",
		ToolsURL:       "http://desktop.docker.com/mcp/catalog/v" + strconv.Itoa(Version) + "/tools/" + server.Name + ".json",
		Source:         source,
		Upstream:       upstream,
		Icon:           server.About.Icon,
		Secrets:        secrets,
		Env:            env,
		Command:        server.Run.Command,
		Volumes:        server.Run.Volumes,
		DisableNetwork: server.Run.DisableNetwork,
		AllowHosts:     server.Run.AllowHosts,
		Config:         config,
		Metadata:       meta,
		OAuth:          oauth,
	}, nil
}

func addDot(text string) string {
	if strings.HasSuffix(text, ".") {
		return text
	}
	return text + "."
}

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

package github

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/docker/mcp-registry/pkg/servers"
	"github.com/google/go-github/v70/github"
)

func NewFromServer(server servers.Server) *Client {
	// A couple of public GitHub repos can't be accessed with authentication if running on GHActions...
	// See https://github.com/xaf/omni/issues/670
	if server.Name == "shopify" || server.Name == "heroku" {
		return NewUnauthenticated()
	}
	return New()
}

func New() *Client {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return NewUnauthenticated()
	}

	return &Client{
		gh: github.NewClient(nil).WithAuthToken(token),
	}
}

func NewUnauthenticated() *Client {
	return &Client{
		gh: github.NewClient(nil),
	}
}

type Client struct {
	gh *github.Client
}

func (c *Client) GetProjectRepository(ctx context.Context, project string) (*github.Repository, error) {
	owner, repo, err := extractOrgAndProject(project)
	if err != nil {
		return nil, err
	}

	for {
		repository, _, err := c.gh.Repositories.Get(ctx, owner, repo)
		if sleepOnRateLimitError(ctx, err) {
			continue
		}

		return repository, err
	}
}

func (c *Client) GetCommitSHA1(ctx context.Context, project, branch string) (string, error) {
	owner, repo, err := extractOrgAndProject(project)
	if err != nil {
		return "", err
	}

	for {
		sha, _, err := c.gh.Repositories.GetCommitSHA1(ctx, owner, repo, branch, "")
		if sleepOnRateLimitError(ctx, err) {
			continue
		}

		return sha, err
	}
}

func (c *Client) FindIcon(ctx context.Context, projectURL string) (string, error) {
	repository, err := c.GetProjectRepository(ctx, projectURL)
	if err != nil {
		return "", err
	}

	return repository.Owner.GetAvatarURL(), nil
}

func sleepOnRateLimitError(ctx context.Context, err error) bool {
	var rateLimitErr *github.RateLimitError
	if !errors.As(err, &rateLimitErr) {
		return false
	}

	sleepDelay := time.Until(rateLimitErr.Rate.Reset.Time)
	fmt.Printf("Rate limit exceeded, waiting %d seconds for reset...\n", int64(sleepDelay.Seconds()))

	select {
	case <-ctx.Done():
	case <-time.After(sleepDelay):
	}

	return true
}

func extractOrgAndProject(rawURL string) (string, string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("URL path doesn't contain enough segments: %s", rawURL)
	}

	org := parts[0]
	project := parts[1]

	return org, project, nil
}

func FindTags(topics []string) []string {
	if len(topics) == 0 {
		return []string{"TODO"}
	}

	var tags []string
	for _, topic := range topics {
		if topic != "mcp" && topic != "mcp-server" {
			tags = append(tags, topic)
		}
	}

	return tags
}

type DetectedInfo struct {
	ProjectURL string
	Branch     string
	Directory  string
}

func DetectBranchAndDirectory(projectURL string, repository *github.Repository) (DetectedInfo, error) {
	u, err := url.Parse(projectURL)
	if err != nil {
		return DetectedInfo{}, err
	}

	var branch string
	var directory string
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) >= 4 && parts[2] == "tree" {
		projectURL = u.Scheme + "://" + u.Host + "/" + parts[0] + "/" + parts[1]
		if parts[3] == "main" { // Should match with any valid branch
			branch = parts[3]
			directory = strings.Join(parts[4:], "/")
		} else {
			branch = strings.Join(parts[3:], "/")
		}
	} else if len(parts) >= 4 && parts[2] == "blob" {
		projectURL = u.Scheme + "://" + u.Host + "/" + parts[0] + "/" + parts[1]
		if parts[3] == "main" { // Should match with any valid branch
			branch = parts[3]
			directory = strings.Join(parts[4:], "/")
		} else {
			branch = strings.Join(parts[3:], "/")
		}
	} else if len(parts) == 4 && parts[2] == "pull" {
		projectURL = u.Scheme + "://" + u.Host + "/" + parts[0] + "/" + parts[1]
		branch = "refs/pull/" + parts[3] + "/merge"
	} else if len(parts) == 4 && parts[2] == "commit" {
		projectURL = u.Scheme + "://" + u.Host + "/" + parts[0] + "/" + parts[1]
		branch = parts[3]
	} else {
		branch = repository.GetDefaultBranch()
	}

	return DetectedInfo{
		ProjectURL: projectURL,
		Branch:     branch,
		Directory:  directory,
	}, nil
}

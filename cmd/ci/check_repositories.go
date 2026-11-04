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

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type repositoryResult struct {
	Server       string `json:"server"`
	File         string `json:"file"`
	Project      string `json:"project"`
	CanonicalURL string `json:"canonical_url,omitempty"`
	Status       string `json:"status"`
	Moved        bool   `json:"moved"`
	Archived     bool   `json:"archived"`
	Error        string `json:"error,omitempty"`
}
type repositoryReport struct {
	CheckedAt          time.Time          `json:"checked_at"`
	Commit             string             `json:"commit,omitempty"`
	UniqueRepositories int                `json:"unique_repositories"`
	Counts             map[string]int     `json:"counts"`
	Results            []repositoryResult `json:"results"`
}
type repositoryChecker struct {
	client   *http.Client
	baseURL  string
	token    string
	now      func() time.Time
	sleep    func(context.Context, time.Duration) error
	attempts int
}

func newRepositoryChecker() *repositoryChecker {
	return &repositoryChecker{
		client: &http.Client{Timeout: 30 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return errors.New("too many redirects")
			}
			if req.URL.Scheme != "https" || req.URL.Host != "api.github.com" || req.URL.User != nil {
				return errors.New("unsafe API redirect")
			}
			return nil
		}},
		baseURL: "https://api.github.com", token: os.Getenv("GITHUB_TOKEN"), now: time.Now, attempts: 4,
		sleep: func(ctx context.Context, d time.Duration) error {
			timer := time.NewTimer(d)
			defer timer.Stop()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timer.C:
				return nil
			}
		},
	}
}

// sourceRepository accepts repository URLs and GitHub tree/blob subpaths, but
// never sends credentials to a host obtained from the catalog.
func sourceRepository(raw string) (string, error) {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return "", err
	}
	if strings.Contains(strings.ToLower(u.EscapedPath()), "%2f") {
		return "", errors.New("encoded path separators are not supported")
	}
	if u.Hostname() == "" {
		return "", errors.New("source.project must be an absolute URL")
	}
	if !strings.EqualFold(u.Hostname(), "github.com") {
		return "", nil
	}
	if u.Scheme != "https" || u.User != nil || u.Port() != "" {
		return "", errors.New("expected an HTTPS github.com URL without credentials or port")
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return "", errors.New("missing owner or repository")
	}
	owner, repo := parts[0], strings.TrimSuffix(parts[1], ".git")
	for _, s := range []string{owner, repo} {
		if s == "" || s == "." || s == ".." {
			return "", errors.New("invalid owner or repository")
		}
		for _, c := range s {
			if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-' || c == '_' || c == '.') {
				return "", errors.New("invalid owner or repository")
			}
		}
	}
	return owner + "/" + repo, nil
}

func (c *repositoryChecker) check(ctx context.Context, repository string, anonymous bool) repositoryResult {
	result := repositoryResult{Status: "unknown"}
	for attempt := 0; attempt < c.attempts; attempt++ {
		if err := ctx.Err(); err != nil {
			result.Error = err.Error()
			return result
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/repos/"+repository, nil)
		if err != nil {
			result.Error = err.Error()
			return result
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		if c.token != "" && !anonymous {
			req.Header.Set("Authorization", "Bearer "+c.token)
		}
		resp, err := c.client.Do(req)
		retry := false
		delay := time.Second * time.Duration(1<<attempt)
		if err != nil {
			result.Error = "request failed: " + err.Error()
			retry = true
		} else {
			body, readErr := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
			resp.Body.Close()
			switch {
			case readErr != nil:
				result.Error = "response read failed: " + readErr.Error()
				retry = true
			case resp.StatusCode == http.StatusOK:
				var metadata struct {
					FullName string `json:"full_name"`
					HTMLURL  string `json:"html_url"`
					Archived bool   `json:"archived"`
				}
				if err := json.Unmarshal(body, &metadata); err != nil {
					result.Error = "invalid metadata: " + err.Error()
					return result
				}
				canonical, err := sourceRepository(metadata.HTMLURL)
				if err != nil || canonical == "" || !strings.EqualFold(canonical, metadata.FullName) {
					result.Error = "invalid canonical repository metadata"
					return result
				}
				result.Status = "healthy"
				result.Moved = !strings.EqualFold(repository, metadata.FullName)
				result.Archived = metadata.Archived
				result.CanonicalURL = metadata.HTMLURL
				if result.Moved || result.Archived {
					result.Status = "finding"
				}
				result.Error = ""
				return result
			case resp.StatusCode == http.StatusNotFound:
				result.Status = "unavailable"
				result.Error = "404: deleted, private, or unavailable without a redirect"
				return result
			default:
				result.Error = fmt.Sprintf("GitHub HTTP %d", resp.StatusCode)
				limited := resp.StatusCode == 429 || resp.StatusCode == 403 && (resp.Header.Get("Retry-After") != "" || resp.Header.Get("X-RateLimit-Remaining") == "0" || strings.Contains(strings.ToLower(string(body)), "rate limit"))
				retry = limited || resp.StatusCode >= 500
				if limited {
					delay = time.Minute
				}
				if limited && attempt > 0 {
					delay = time.Minute * time.Duration(1<<attempt)
				}
				if limited {
					if value := resp.Header.Get("Retry-After"); value != "" {
						if n, e := strconv.Atoi(value); e == nil && n >= 0 {
							delay = max(delay, time.Duration(n)*time.Second)
						} else if t, e := http.ParseTime(value); e == nil {
							delay = max(delay, t.Sub(c.now()))
						}
					}
					if resp.Header.Get("X-RateLimit-Remaining") == "0" {
						if n, e := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64); e == nil {
							delay = max(delay, time.Unix(n, 0).Sub(c.now())+time.Second)
						}
					}
				}
			}
		}
		if !retry || attempt+1 == c.attempts {
			return result
		}
		if deadline, ok := ctx.Deadline(); ok && delay >= time.Until(deadline) {
			result.Error += "; required retry wait exceeds scan deadline"
			return result
		}
		if err := c.sleep(ctx, delay); err != nil {
			result.Error += "; retry interrupted: " + err.Error()
			return result
		}
	}
	return result
}

func collectRepositoryReport(ctx context.Context, workspace, commit string, c *repositoryChecker) (repositoryReport, error) {
	report := repositoryReport{CheckedAt: c.now().UTC(), Commit: commit, Counts: map[string]int{}, Results: []repositoryResult{}}
	paths, err := filepath.Glob(filepath.Join(workspace, "servers", "*", "server.yaml"))
	if err != nil {
		return report, err
	}
	if len(paths) == 0 {
		return report, errors.New("no servers/*/server.yaml files found")
	}
	type target struct {
		row  repositoryResult
		repo string
	}
	targets := []target{}
	anonymous := map[string]bool{}
	for _, path := range paths {
		rel, _ := filepath.Rel(workspace, path)
		row := repositoryResult{Server: filepath.Base(filepath.Dir(path)), File: filepath.ToSlash(rel), Status: "skipped"}
		doc, e := loadServerYAMLFromWorkspace(workspace, rel)
		repo := ""
		if e != nil {
			row.Status = "unknown"
			row.Error = e.Error()
		} else {
			row.Project = doc.Source.Project
			if strings.TrimSpace(row.Project) == "" {
				row.Error = "no source.project"
			} else {
				repo, e = sourceRepository(row.Project)
				if e != nil {
					row.Status = "unknown"
					row.Error = e.Error()
				} else if repo == "" {
					row.Error = "non-GitHub source"
				}
			}
		}
		key := strings.ToLower(repo)
		if repo != "" && (row.Server == "shopify" || row.Server == "heroku") {
			anonymous[key] = true
		}
		targets = append(targets, target{row: row, repo: repo})
	}
	cache := map[string]repositoryResult{}
	incomplete := false
	for _, t := range targets {
		row := t.row
		if t.repo != "" {
			key := strings.ToLower(t.repo)
			checked, ok := cache[key]
			if !ok {
				checked = c.check(ctx, t.repo, anonymous[key])
				cache[key] = checked
			}
			row.Status = checked.Status
			row.Moved = checked.Moved
			row.Archived = checked.Archived
			row.CanonicalURL = checked.CanonicalURL
			row.Error = checked.Error
		}
		if row.Status == "unknown" {
			incomplete = true
		}
		report.Results = append(report.Results, row)
		report.Counts[row.Status]++
	}
	report.UniqueRepositories = len(cache)
	if incomplete {
		return report, errors.New("repository scan incomplete; inspect unknown results")
	}
	return report, nil
}

func markdownCell(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "|", "&#124;", "\n", " ", "\r", " ")
	return r.Replace(s)
}
func repositoryMarkdown(r repositoryReport) string {
	var out strings.Builder
	fmt.Fprintf(&out, "# Repository status report\n\nChecked: %s\n\nCommit: %s\n\nUnique repositories: %d\n\n", r.CheckedAt.Format(time.RFC3339), markdownCell(r.Commit), r.UniqueRepositories)
	keys := make([]string, 0, len(r.Counts))
	for k := range r.Counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&out, "- %s: %d\n", k, r.Counts[k])
	}
	out.WriteString("\n404 means unavailable, not proven deleted. Findings require human review.\n\n| Server / file | Configured URL | Canonical URL | Status | Details |\n| --- | --- | --- | --- | --- |\n")
	for _, row := range r.Results {
		if row.Status == "healthy" {
			continue
		}
		details := row.Error
		if row.Moved {
			details += " moved/renamed"
		}
		if row.Archived {
			details += " archived"
		}
		fmt.Fprintf(&out, "| %s (%s) | %s | %s | %s | %s |\n", markdownCell(row.Server), markdownCell(row.File), markdownCell(row.Project), markdownCell(row.CanonicalURL), row.Status, markdownCell(strings.TrimSpace(details)))
	}
	return out.String()
}
func runCheckRepositories(args []string) error {
	flags := flag.NewFlagSet("check-repositories", flag.ContinueOnError)
	workspace := flags.String("workspace", ".", "repository workspace")
	jsonPath := flags.String("output-json", "repository-status.json", "JSON report path")
	markdownPath := flags.String("output-markdown", "repository-status.md", "Markdown report path")
	timeout := flags.Duration("timeout", 20*time.Minute, "overall scan deadline")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *timeout <= 0 {
		return errors.New("timeout must be positive")
	}
	commit, _ := runGitCommand(*workspace, "rev-parse", "HEAD")
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	report, scanErr := collectRepositoryReport(ctx, *workspace, strings.TrimSpace(commit), newRepositoryChecker())
	if err := writeJSONFile(*jsonPath, report); err != nil {
		return err
	}
	if err := os.WriteFile(*markdownPath, []byte(repositoryMarkdown(report)), 0o644); err != nil {
		return err
	}
	fmt.Printf("Checked %d repositories; report: %s\n", report.UniqueRepositories, *markdownPath)
	return scanErr
}

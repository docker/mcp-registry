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
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestSourceRepository(t *testing.T) {
	for _, tc := range []struct {
		input, want string
		bad         bool
	}{
		{"https://github.com/Owner/Repo.git/", "Owner/Repo", false},
		{"https://github.com/owner/repo/tree/main/path", "owner/repo", false},
		{"https://gitlab.com/owner/repo", "", false},
		{"https://github.com/owner", "", true},
		{"https://github.com/owner/../repo", "", true},
		{"https://github.com/owner/re%2Fpo", "", true},
		{"http://github.com/owner/repo", "", true},
		{"https://user:secret@github.com/owner/repo", "", true},
		{"https://github.com:443/owner/repo", "", true},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, err := sourceRepository(tc.input)
			if got != tc.want || (err != nil) != tc.bad {
				t.Fatalf("got %q, %v", got, err)
			}
		})
	}
}
func TestRepositoryStatuses(t *testing.T) {
	for _, tc := range []struct {
		name, fullName  string
		code            int
		archived, moved bool
		status          string
	}{
		{"healthy", "owner/repo", 200, false, false, "healthy"},
		{"case-only", "Owner/Repo", 200, false, false, "healthy"},
		{"moved", "new/repo", 200, false, true, "finding"},
		{"renamed", "owner/new", 200, false, true, "finding"},
		{"archived-moved", "new/repo", 200, true, true, "finding"},
		{"archived", "owner/repo", 200, true, false, "finding"},
		{"unavailable", "", 404, false, false, "unavailable"},
		{"unauthorized", "", 401, false, false, "unknown"},
		{"forbidden", "", 403, false, false, "unknown"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				if r.URL.Path != "/repos/owner/repo" {
					t.Error(r.URL.Path)
				}
				if r.Header.Get("Authorization") != "Bearer token" {
					t.Error("missing auth")
				}
				w.WriteHeader(tc.code)
				fmt.Fprintf(w, `{"full_name":%q,"html_url":%q,"archived":%t}`, tc.fullName, "https://github.com/"+tc.fullName, tc.archived)
			}))
			defer server.Close()
			c := newRepositoryChecker()
			c.baseURL = server.URL
			c.token = "token"
			got := c.check(context.Background(), "owner/repo", false)
			if got.Status != tc.status || got.Moved != tc.moved || got.Archived != tc.archived || calls != 1 {
				t.Fatalf("result %+v calls %d", got, calls)
			}
		})
	}
}
func TestRepositoryRetry(t *testing.T) {
	now := time.Date(2026, 10, 25, 0, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name    string
		code    int
		headers map[string]string
		body    string
		delay   time.Duration
	}{
		{"retry-after", 429, map[string]string{"Retry-After": "120"}, "", 120 * time.Second},
		{"retry-date", 429, map[string]string{"Retry-After": now.Add(2 * time.Minute).Format(http.TimeFormat)}, "", 2 * time.Minute},
		{"reset", 403, map[string]string{"X-RateLimit-Remaining": "0", "X-RateLimit-Reset": fmt.Sprint(now.Add(2 * time.Minute).Unix())}, "", 121 * time.Second},
		{"secondary", 403, nil, `{"message":"secondary rate limit"}`, time.Minute},
		{"server", 503, nil, "", time.Second},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			waits := []time.Duration{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				if calls == 1 {
					for k, v := range tc.headers {
						w.Header().Set(k, v)
					}
					w.WriteHeader(tc.code)
					fmt.Fprint(w, tc.body)
					return
				}
				fmt.Fprint(w, `{"full_name":"owner/repo","html_url":"https://github.com/owner/repo"}`)
			}))
			defer server.Close()
			c := newRepositoryChecker()
			c.baseURL = server.URL
			c.now = func() time.Time { return now }
			c.sleep = func(ctx context.Context, d time.Duration) error { waits = append(waits, d); return nil }
			got := c.check(context.Background(), "owner/repo", false)
			if got.Status != "healthy" || calls != 2 || len(waits) != 1 || waits[0] != tc.delay {
				t.Fatalf("%+v calls=%d waits=%v", got, calls, waits)
			}
		})
	}
}
func TestBoundedRetriesAndCancellation(t *testing.T) {
	calls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++; w.WriteHeader(503) }))
	defer s.Close()
	c := newRepositoryChecker()
	c.baseURL = s.URL
	c.sleep = func(context.Context, time.Duration) error { return nil }
	got := c.check(context.Background(), "owner/repo", false)
	if got.Status != "unknown" || calls != 4 {
		t.Fatalf("%+v calls %d", got, calls)
	}
	calls = 0
	c.sleep = func(context.Context, time.Duration) error { return context.Canceled }
	got = c.check(context.Background(), "owner/repo", false)
	if calls != 1 || !strings.Contains(got.Error, "interrupted") {
		t.Fatalf("%+v calls %d", got, calls)
	}
}
func TestRedirectPolicy(t *testing.T) {
	c := newRepositoryChecker()
	for _, raw := range []string{"http://api.github.com/repos/x/y", "https://example.com/repos/x/y", "https://api.github.com.evil.test/repos/x/y"} {
		req, _ := http.NewRequest("GET", raw, nil)
		if c.client.CheckRedirect(req, nil) == nil {
			t.Fatal("accepted unsafe redirect", raw)
		}
	}
	req, _ := http.NewRequest("GET", "https://api.github.com/repositories/123", nil)
	if err := c.client.CheckRedirect(req, nil); err != nil {
		t.Fatal(err)
	}
	if c.client.CheckRedirect(req, make([]*http.Request, 5)) == nil {
		t.Fatal("unbounded redirects")
	}
	// Exercise redirect handling with a local transport; production policy stays intact.
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "https://example.com/steal", 301) }))
	defer s.Close()
	c.baseURL = s.URL
	c.sleep = func(context.Context, time.Duration) error { return nil }
	if got := c.check(context.Background(), "owner/repo", false); got.Status != "unknown" {
		t.Fatal(got)
	}
}
func TestMetadataAndNetworkErrors(t *testing.T) {
	for _, body := range []string{"bad json", `{"full_name":"owner/repo","html_url":"https://evil.test/owner/repo"}`, `{}`} {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, body) }))
		c := newRepositoryChecker()
		c.baseURL = s.URL
		if got := c.check(context.Background(), "owner/repo", false); got.Status != "unknown" {
			t.Fatal(got)
		}
		s.Close()
	}
	c := newRepositoryChecker()
	c.baseURL = "http://127.0.0.1:1"
	calls := 0
	c.sleep = func(context.Context, time.Duration) error { calls++; return nil }
	if got := c.check(context.Background(), "owner/repo", false); got.Status != "unknown" || calls != 3 {
		t.Fatalf("%+v waits=%d", got, calls)
	}
}
func writeServerFixture(t *testing.T, dir, name, content string) {
	t.Helper()
	path := filepath.Join(dir, "servers", name)
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, "server.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
func TestRepositoryReport(t *testing.T) {
	dir := t.TempDir()
	writeServerFixture(t, dir, "a", "type: remote\nsource:\n  project: https://github.com/Owner/Repo.git\n")
	writeServerFixture(t, dir, "shopify", "source:\n  project: https://github.com/owner/repo\n")
	writeServerFixture(t, dir, "skip", "source:\n  project: https://gitlab.com/a/b\n")
	writeServerFixture(t, dir, "empty", "name: empty\n")
	writeServerFixture(t, dir, "invalid", "source: [\n")
	calls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Header.Get("Authorization") != "" {
			t.Error("exception must be anonymous")
		}
		fmt.Fprint(w, `{"full_name":"new/repo","html_url":"https://github.com/new/repo","archived":true}`)
	}))
	defer s.Close()
	c := newRepositoryChecker()
	c.baseURL = s.URL
	c.token = "test"
	report, err := collectRepositoryReport(context.Background(), dir, "commit", c)
	if err == nil || calls != 1 || report.UniqueRepositories != 1 || report.Counts["finding"] != 2 || report.Counts["unknown"] != 1 || report.Counts["skipped"] != 2 {
		t.Fatalf("%+v calls=%d err=%v", report, calls, err)
	}
	for i := 1; i < len(report.Results); i++ {
		if report.Results[i-1].File > report.Results[i].File {
			t.Fatal("not sorted")
		}
	}
	if !strings.Contains(repositoryMarkdown(report), "new/repo") {
		t.Fatal("missing canonical URL")
	}
	if got := markdownCell("<x>|a\nb"); got != "&lt;x&gt;&#124;a b" {
		t.Fatal(got)
	}
}
func TestReportExitSemantics(t *testing.T) {
	dir := t.TempDir()
	writeServerFixture(t, dir, "a", "source:\n  project: https://github.com/owner/repo\n")
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(404) }))
	defer s.Close()
	c := newRepositoryChecker()
	c.baseURL = s.URL
	r, err := collectRepositoryReport(context.Background(), dir, "", c)
	if err != nil || r.Counts["unavailable"] != 1 {
		t.Fatalf("%+v %v", r, err)
	}
	_, err = collectRepositoryReport(context.Background(), t.TempDir(), "", c)
	if err == nil {
		t.Fatal("empty scan succeeded")
	}
}
func TestSleepHonorsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(newRepositoryChecker().sleep(ctx, time.Hour), context.Canceled) {
		t.Fatal("sleep ignored context")
	}
}

func TestCLIReportsIncompleteScan(t *testing.T) {
	dir := t.TempDir()
	writeServerFixture(t, dir, "invalid", "source: [\n")
	jsonPath, mdPath := filepath.Join(dir, "report.json"), filepath.Join(dir, "report.md")
	if err := runCheckRepositories([]string{"--workspace", dir, "--output-json", jsonPath, "--output-markdown", mdPath}); err == nil {
		t.Fatal("expected incomplete scan")
	}
	for _, path := range []string{jsonPath, mdPath} {
		if info, err := os.Stat(path); err != nil || info.Size() == 0 {
			t.Fatalf("missing partial report: %s", path)
		}
	}
	if err := runCheckRepositories([]string{"--timeout", "0s"}); err == nil {
		t.Fatal("invalid timeout accepted")
	}
}
func TestCanceledScanMakesNoRequest(t *testing.T) {
	calls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++ }))
	defer s.Close()
	c := newRepositoryChecker()
	c.baseURL = s.URL
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	got := c.check(ctx, "owner/repo", false)
	if calls != 0 || got.Status != "unknown" {
		t.Fatalf("%+v calls=%d", got, calls)
	}
}
func TestRedirectResolution(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/repos/old/repo" {
			http.Redirect(w, r, server.URL+"/repositories/123", 301)
			return
		}
		if r.URL.Path != "/repositories/123" {
			t.Error(r.URL.Path)
		}
		fmt.Fprint(w, `{"full_name":"new/repo","html_url":"https://github.com/new/repo"}`)
	}))
	defer server.Close()
	c := newRepositoryChecker()
	c.baseURL = server.URL
	// Local test-only redirect policy; production HTTPS-host policy is tested separately.
	c.client.CheckRedirect = func(r *http.Request, via []*http.Request) error {
		if len(via) > 5 {
			return errors.New("too many redirects")
		}
		return nil
	}
	got := c.check(context.Background(), "old/repo", false)
	if !got.Moved || got.Status != "finding" {
		t.Fatal(got)
	}
}

func TestLongRateLimitDoesNotConsumeScanDeadline(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "3600")
		w.WriteHeader(429)
	}))
	defer s.Close()
	c := newRepositoryChecker()
	c.baseURL = s.URL
	c.sleep = func(context.Context, time.Duration) error { t.Fatal("must not sleep beyond deadline"); return nil }
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	got := c.check(ctx, "owner/repo", false)
	if got.Status != "unknown" || !strings.Contains(got.Error, "required retry wait") || ctx.Err() != nil {
		t.Fatal(got)
	}
}

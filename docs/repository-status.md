# Source repository status checks

The **Check source repository status** workflow checks GitHub repositories named
by `source.project` in `servers/*/server.yaml`. It runs weekly on Monday at
06:23 UTC and can also be started with **Run workflow**. GitHub schedules are
best-effort and run from the default branch after the workflow is merged.

This is a metadata check, not an assessment of server quality or availability.
It does not execute upstream code, edit the catalog, delete entries, open pull
requests, or post comments/issues. Changes to catalog entries require review.

## Run locally

With Go 1.24 or newer, from the repository root:

```sh
# Optional: set GITHUB_TOKEN through your normal secure environment setup.
# A token raises the API rate limit; never put one in a command or report.
task ci -- check-repositories
# Equivalent, without Task:
go run ./cmd/ci check-repositories
```

Options:

- `--workspace .`: checkout containing `servers/*/server.yaml`.
- `--output-json repository-status.json`: complete machine-readable results.
- `--output-markdown repository-status.md`: human-readable report.
- `--timeout 20m`: overall scan deadline, including retries.

A single sequential request is made per distinct configured repository name,
ignoring case. All referencing server entries remain in the report. GitHub
repository URLs with `.git`, trailing slashes, and tree/blob subpaths are
normalized to owner/repository. Sources outside GitHub and absent source URLs
are listed as skipped. Malformed GitHub URLs or YAML are unknown, not skipped.
The existing Shopify/Heroku anonymous-access exceptions also apply; if another
entry shares that repository, its request uses the same anonymous result.

## Interpret results

| Status | Meaning |
| --- | --- |
| `healthy` | GitHub returned metadata; the name matches and the repo is not archived. |
| `finding` | The canonical owner/name differs or `archived` is true, or both. |
| `unavailable` | GitHub returned 404. It may be deleted, private, or unavailable without a redirect. |
| `unknown` | The check could not establish status, including invalid input, auth errors, exhausted retries, or deadline. |
| `skipped` | No GitHub source repository to check. The reason is included. |

The `moved` and `archived` flags are independent. Case-only name differences are
not moves. Suggested URLs come from GitHub's returned metadata, not guesses.
Every JSON row includes the catalog file, entry name, and configured URL;
metadata results also include the canonical URL. The Markdown report lists
non-healthy entries with counts and scan time/commit. Reports can contain
configuration text, but never intentionally include authentication tokens.

Requests have a 30-second timeout, a maximum of five redirect responses, and
at most four attempts. Redirects must stay on HTTPS `api.github.com`. Network
failures and 5xx responses use exponential backoff. Rate-limit responses respect
`Retry-After` and `X-RateLimit-Reset` with at least a minute between attempts and
exponentially increasing waits. Other authentication/permission errors are not
retried. If a required wait exceeds the remaining scan deadline, that repository is
reported as unknown immediately so other repositories can still be checked;
the checker never retries before the server permits it. The overall deadline
can also end a wait rather than retry early.

## Exit status and workflow output

Known findings (including 404s) do **not** fail the scan. Unknown results,
invalid input, or report-writing errors return a nonzero exit status. Partial
reports are written even when a scan is incomplete. A missing/empty server
collection is an error, not a successful zero-result scan.

The workflow publishes Markdown to the run summary and uploads JSON/Markdown
artifacts for 30 days, also after scan failure when files exist. It needs only
`contents: read`, uses `GITHUB_TOKEN`, and does not need a bot app or write
permission. Reports require a person to review them; there is no separate
notification channel. API checks do not run on ordinary pull requests; normal
Go unit tests cover the checker with mock HTTP responses instead.

Before proposing catalog corrections, re-check each finding. In particular,
a 404 is never evidence for automatic deletion. If a source URL contains a
subpath, updating it requires preserving that path deliberately.

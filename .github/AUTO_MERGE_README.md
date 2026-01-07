# Auto-Merge for Pin Upgrade PRs

## Overview

This repository automatically merges pin upgrade PRs for allowlisted servers when all security and quality checks pass. This reduces manual review burden while maintaining security standards.

## How It Works

1. **Pin Update Created**: The `mcp-registry-bot` creates a PR to update a server's pinned version
2. **CI Runs**: All checks run (CodeQL, security review, validation)
3. **Auto-Merge Triggers**: If all checks pass and the server is allowlisted, the PR is automatically merged
4. **Manual Override**: Add the `skip-auto-merge` label to prevent auto-merge on specific PRs

## Workflow Files

- `.github/workflows/auto-merge-pins.yaml` - Main auto-merge workflow
- `.github/auto-merge-allowlist.yaml` - Configuration for allowed servers

## Checks Verified Before Merge

The workflow verifies ALL GitHub checks have passed, including:

- ✅ CodeQL / Analyze (actions)
- ✅ CodeQL / Analyze (go)
- ✅ Code and Server Validation
- ✅ Code scanning results / CodeQL
- ✅ Security Review Trigger
- ✅ security-review/claude/pin/{server}

## Adding Servers to Allowlist

Before adding a server to the allowlist, verify:

1. **Stable Build Process**: Server has consistent successful builds
2. **Recent History**: Check last 5-10 pin update PRs for issues
3. **Test Coverage**: Server has adequate testing
4. **Low Risk**: Updates typically don't introduce breaking changes

### Steps to Add

1. Edit `.github/auto-merge-allowlist.yaml`
2. Add the server name under the `servers:` section
3. Commit and push to `main`
4. Monitor the first few auto-merges closely

Example:
```yaml
servers:
  - sonarqube
  - redis
  - your-new-server
```

## Blocking Auto-Merge

To prevent auto-merge on a specific PR:

1. Add the `skip-auto-merge` label to the PR
2. The workflow will skip auto-merge and wait for manual review

## Currently Allowlisted Servers

See `.github/auto-merge-allowlist.yaml` for the current list. Initial allowlist includes:
- sonarqube
- redis
- postgres
- mongodb
- playwright
- firecrawl
- browserbase
- temporal
- stripe
- buildkite
- grafana

## Monitoring

Check the workflow runs at:
`https://github.com/docker/mcp-registry/actions/workflows/auto-merge-pins.yaml`

When auto-merge runs:
- ✅ Success: PR is merged with a comment
- ⚠️ Failure: PR gets a comment with error details
- ⏭️ Skipped: PR doesn't meet criteria (wrong author, not allowlisted, checks failed)

## Troubleshooting

### PR Not Auto-Merging

Check:
1. Is the server in `.github/auto-merge-allowlist.yaml`?
2. Did all checks pass (including security review)?
3. Is the PR from `mcp-registry-bot[bot]`?
4. Does the branch match `automation/update-pin-{SERVER}`?
5. Is there a `skip-auto-merge` label?

### Disabling Auto-Merge

To temporarily disable:
1. Remove servers from `.github/auto-merge-allowlist.yaml`
2. Or add `skip-auto-merge` label to individual PRs

To permanently disable:
1. Disable the workflow in GitHub Actions settings
2. Or delete `.github/workflows/auto-merge-pins.yaml`

## Benefits

- **Reduced Manual Work**: No need to review routine pin updates
- **Faster Updates**: Dependencies get updated as soon as checks pass
- **Security Maintained**: All security checks must pass before merge
- **Lower PR Backlog**: Reduces the 500+ pin update PRs per month

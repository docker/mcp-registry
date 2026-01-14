#!/bin/bash
set -e

pr_number="$1"

# Extract server name from changed files (more reliable than branch name)
# Pin update PRs typically modify servers/{server-name}/ files
changed_files=$(gh pr view "$pr_number" --json files --jq '.files[].path')

# Find unique server directories that were modified
servers=$(echo "$changed_files" | grep '^servers/' | cut -d'/' -f2 | sort -u)
server_count=$(echo "$servers" | grep -v '^$' | wc -l | tr -d ' ')

if [ "$server_count" -eq 1 ]; then
  server=$(echo "$servers" | grep -v '^$')
  echo "server=$server" >> "$GITHUB_OUTPUT"
  echo "Detected server from changed files: $server"
elif [ "$server_count" -gt 1 ]; then
  echo "Multiple servers detected in PR changes: $servers"
  echo "Pin update PRs should only modify one server. Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
else
  echo "No server directories found in changed files."
  echo "Could not determine server name. Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

echo "PR validation passed"

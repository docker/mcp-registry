#!/bin/bash
set -e

author="$1"
state="$2"
is_draft="$3"
labels="$4"

# Check if PR is from mcp-registry-bot
if [ "$author" != "mcp-registry-bot[bot]" ]; then
  echo "PR is not from mcp-registry-bot. Author: $author. Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

# Check if PR is still open
if [ "$state" != "OPEN" ]; then
  echo "PR is not open (state: $state). Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

# Check if PR is a draft
if [ "$is_draft" == "true" ]; then
  echo "PR is a draft. Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

# Check for skip-auto-merge label
if echo ",$labels," | grep -q ",skip-auto-merge,"; then
  echo "PR has 'skip-auto-merge' label. Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

echo "PR validation passed"

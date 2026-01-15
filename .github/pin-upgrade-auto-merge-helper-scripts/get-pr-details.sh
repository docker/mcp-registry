#!/bin/bash
set -e

# Get PR number from workflow_dispatch input or check_run event
if [ "$GITHUB_EVENT_NAME" = "workflow_dispatch" ]; then
  pr_number="$INPUT_PR_NUMBER"
  echo "Using PR number from manual trigger: $pr_number"
else
  # Get the PR number from the check_run event payload
  pr_number="$CHECK_RUN_PR_NUMBER"

  if [ -z "$pr_number" ] || [ "$pr_number" = "null" ]; then
    echo "No PR associated with this check run. Exiting."
    echo "skip=true" >> "$GITHUB_OUTPUT"
    exit 0
  fi

  echo "Using PR number from check run: $pr_number"
  echo "Check run: $CHECK_RUN_NAME - $CHECK_RUN_CONCLUSION"
fi

echo "pr_number=$pr_number" >> "$GITHUB_OUTPUT"

# Get PR details
pr_json=$(gh pr view "$pr_number" --json author,headRefName,title,state,isDraft,labels)

author=$(echo "$pr_json" | jq -r '.author.login')
branch=$(echo "$pr_json" | jq -r '.headRefName')
title=$(echo "$pr_json" | jq -r '.title')
state=$(echo "$pr_json" | jq -r '.state')
is_draft=$(echo "$pr_json" | jq -r '.isDraft')
labels=$(echo "$pr_json" | jq -r '.labels[].name' | tr '\n' ',' | sed 's/,$//')

echo "author=$author" >> "$GITHUB_OUTPUT"
echo "branch=$branch" >> "$GITHUB_OUTPUT"
echo "title=$title" >> "$GITHUB_OUTPUT"
echo "state=$state" >> "$GITHUB_OUTPUT"
echo "is_draft=$is_draft" >> "$GITHUB_OUTPUT"
echo "labels=$labels" >> "$GITHUB_OUTPUT"

echo "PR #$pr_number: $title"
echo "Author: $author"
echo "Branch: $branch"
echo "State: $state"
echo "Draft: $is_draft"
echo "Labels: $labels"

#!/bin/bash
set -e

comment_type="$1"
pr_number="$2"
server="$3"

if [ "$comment_type" = "success" ]; then
  gh pr comment "$pr_number" --body "✅ **Auto-merging pin upgrade PR**

Server \`$server\` is in the allowlist and all required checks have passed.

This PR will be automatically merged."

elif [ "$comment_type" = "failure" ]; then
  workflow_url="$4"
  gh pr comment "$pr_number" --body "⚠️ **Auto-merge failed**

An error occurred while attempting to automatically merge this PR. Manual review and merge may be required.

Check the [workflow run logs]($workflow_url) for details."

else
  echo "Invalid comment type: $comment_type"
  echo "Usage: $0 {success|failure} <pr_number> <server> [workflow_url]"
  exit 1
fi

echo "Posted $comment_type comment to PR #$pr_number"

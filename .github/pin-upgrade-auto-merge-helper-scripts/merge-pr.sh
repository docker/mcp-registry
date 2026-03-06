#!/bin/bash
set -e

pr_number="$1"

if [ -z "$pr_number" ]; then
  echo "Usage: merge-pr.sh <pr_number>"
  exit 1
fi

MAX_ATTEMPTS=5
RETRY_DELAY=10

for ((i = 1; i <= MAX_ATTEMPTS; i++)); do
  echo "Attempt $i/$MAX_ATTEMPTS: merging PR #$pr_number..."

  if gh pr merge "$pr_number" --squash 2>&1; then
    echo "Successfully merged PR #$pr_number"
    exit 0
  fi

  if [ "$i" -lt "$MAX_ATTEMPTS" ]; then
    echo "Merge failed, retrying in ${RETRY_DELAY}s..."
    sleep "$RETRY_DELAY"

    # Bail out if the PR is no longer open (e.g. closed or already merged)
    state=$(gh pr view "$pr_number" --json state --jq '.state')
    if [ "$state" = "MERGED" ]; then
      echo "PR #$pr_number was merged by another process"
      exit 0
    fi
    if [ "$state" != "OPEN" ]; then
      echo "PR #$pr_number is in unexpected state '$state', aborting."
      exit 1
    fi
  fi
done

echo "Failed to merge PR #$pr_number after $MAX_ATTEMPTS attempts."
exit 1

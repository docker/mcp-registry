#!/bin/bash
set -e

pr_number="$1"

if [ -z "$pr_number" ]; then
  echo "Usage: merge-pr.sh <pr_number>"
  exit 1
fi

POLL_INTERVAL=15
MAX_ATTEMPTS=9  # 9 * ~20s = ~3 minutes total

echo "Enabling auto-merge for PR #$pr_number..."
gh pr merge "$pr_number" --squash --auto

for ((i = 1; i <= MAX_ATTEMPTS; i++)); do
  sleep "$POLL_INTERVAL"

  state=$(gh pr view "$pr_number" --json state --jq '.state')
  echo "Attempt $i/$MAX_ATTEMPTS: PR #$pr_number state is '$state'"

  if [ "$state" = "MERGED" ]; then
    echo "Successfully merged PR #$pr_number"
    exit 0
  fi

  if [ "$state" != "OPEN" ]; then
    echo "PR #$pr_number is in unexpected state '$state', aborting."
    exit 1
  fi

  # Cycle auto-merge off/on to unstick it
  echo "Cycling auto-merge for PR #$pr_number..."
  gh pr merge "$pr_number" --disable-auto || true
  sleep 5

  # Re-check state before re-enabling — PR may have merged during the gap
  state=$(gh pr view "$pr_number" --json state --jq '.state')
  if [ "$state" = "MERGED" ]; then
    echo "Successfully merged PR #$pr_number"
    exit 0
  fi

  if [ "$state" = "OPEN" ]; then
    gh pr merge "$pr_number" --squash --auto
  fi
done

# Final check — the last cycle may have unstuck the merge
sleep "$POLL_INTERVAL"
state=$(gh pr view "$pr_number" --json state --jq '.state')
if [ "$state" = "MERGED" ]; then
  echo "Successfully merged PR #$pr_number"
  exit 0
fi

echo "PR #$pr_number was not merged after ~$((MAX_ATTEMPTS * 20))s. Giving up."
exit 1

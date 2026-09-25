#!/bin/bash
set -e

pr_number="$1"
repository="$2"

# Required check prefixes that must exist and pass before auto-merge.
# The security review check is posted by a private repo after a dispatch delay
# (typically 60-90s). Without this gate, the script could see "all checks green"
# before the security review check is even created.
REQUIRED_CHECK_PREFIXES=("security-review/")

# Get all check runs for the PR's head SHA
head_sha=$(gh pr view "$pr_number" --json headRefOid --jq '.headRefOid')

echo "Checking status for commit: $head_sha"

# Get all check runs and their conclusions
check_runs=$(gh api "/repos/$repository/commits/$head_sha/check-runs" --jq '.check_runs')

total_checks=$(echo "$check_runs" | jq 'length')
echo "Total check runs: $total_checks"

if [ "$total_checks" -eq 0 ]; then
  echo "No check runs found for this commit yet. Will retry when checks start."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

# Verify all required checks exist
for prefix in "${REQUIRED_CHECK_PREFIXES[@]}"; do
  match=$(echo "$check_runs" | jq -r --arg p "$prefix" '[.[] | select(.name | startswith($p))] | length')
  if [ "$match" -eq 0 ]; then
    echo "Required check with prefix '$prefix' has not been created yet."
    echo "Will retry when it appears."
    echo "skip=true" >> "$GITHUB_OUTPUT"
    exit 0
  fi
  echo "Required check prefix '$prefix' found ($match match(es))"
done

# Check if any checks failed
failed_checks=$(echo "$check_runs" | jq -r '.[] | select(.conclusion != "success" and .conclusion != "skipped" and .conclusion != null) | .name')

if [ -n "$failed_checks" ]; then
  echo "Some checks have failed:"
  echo "$failed_checks"
  echo "Skipping auto-merge."
  exit 1
fi

# Check if any checks are still pending
pending_checks=$(echo "$check_runs" | jq -r '.[] | select(.status != "completed") | .name')

if [ -z "$pending_checks" ]; then
  echo "All checks have passed successfully!"
  exit 0
fi

# Checks are pending - might be a race condition if only 1-2 checks left
pending_count=$(echo "$pending_checks" | wc -l | tr -d ' ')
echo "Found $pending_count pending checks:"
echo "$pending_checks"

# Race condition mitigation: If only 1-2 checks pending, wait and recheck
# GitHub's API can lag behind webhooks by a few seconds (eventual consistency)
if [ "$pending_count" -le 2 ]; then
  echo ""
  echo "Only $pending_count checks pending - waiting 10s to handle potential GitHub API lag..."
  sleep 10

  # Recheck status
  check_runs=$(gh api "/repos/$repository/commits/$head_sha/check-runs" --jq '.check_runs')

  # Check for failures again
  failed_checks=$(echo "$check_runs" | jq -r '.[] | select(.conclusion != "success" and .conclusion != "skipped" and .conclusion != null) | .name')

  if [ -n "$failed_checks" ]; then
    echo "Some checks have failed:"
    echo "$failed_checks"
    echo "Skipping auto-merge."
    exit 1
  fi

  # Check pending again
  pending_checks=$(echo "$check_runs" | jq -r '.[] | select(.status != "completed") | .name')

  if [ -z "$pending_checks" ]; then
    echo "All checks have passed successfully! (Caught after eventual consistency delay)"
    exit 0
  fi

  pending_count=$(echo "$pending_checks" | wc -l | tr -d ' ')
  echo "Still $pending_count checks pending after recheck:"
  echo "$pending_checks"
fi

echo ""
echo "Checks still pending. Workflow will re-trigger when next check completes."
echo "skip=true" >> "$GITHUB_OUTPUT"
exit 0

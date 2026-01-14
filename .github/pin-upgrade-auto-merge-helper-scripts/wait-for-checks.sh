#!/bin/bash
set -e

pr_number="$1"
repository="$2"

# Get all check runs for the PR's head SHA
head_sha=$(gh pr view "$pr_number" --json headRefOid --jq '.headRefOid')

echo "Checking status for commit: $head_sha"

# Get all check runs and their conclusions
check_runs=$(gh api "/repos/$repository/commits/$head_sha/check-runs" --jq '.check_runs')

total_checks=$(echo "$check_runs" | jq 'length')
echo "Total check runs: $total_checks"

if [ "$total_checks" -eq 0 ]; then
  echo "No check runs found for this commit yet. Will retry when checks start."
  exit 0
fi

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
exit 0

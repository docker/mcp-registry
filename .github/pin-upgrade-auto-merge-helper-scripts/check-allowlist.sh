#!/bin/bash
set -e

server="$1"
allowlist_file=".github/auto-merge-allowlist.yaml"

if [ ! -f "$allowlist_file" ]; then
  echo "Allowlist file not found: $allowlist_file. Skipping auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

# Extract server list from YAML (handle comments and whitespace)
servers=$(grep -E "^[[:space:]]+-" "$allowlist_file" | sed 's/^[[:space:]]*-[[:space:]]*//' | sed 's/#.*//' | tr -d ' ' | tr '\n' ',' | sed 's/,$//')

if [ -z "$servers" ]; then
  echo "Allowlist is empty. No servers configured for auto-merge."
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

echo "servers=$servers" >> "$GITHUB_OUTPUT"
echo "Allowlisted servers: $servers"

# Check if server is in the allowlist (case-insensitive)
server_lc=$(echo "$server" | tr '[:upper:]' '[:lower:]')
servers_lc=$(echo "$servers" | tr '[:upper:]' '[:lower:]')

if ! echo ",$servers_lc," | grep -q ",$server_lc,"; then
  echo "Server '$server' is not in the allowlist. Skipping auto-merge."
  echo "Allowlisted servers: $servers"
  echo "skip=true" >> "$GITHUB_OUTPUT"
  exit 0
fi

echo "Server '$server' is allowlisted for auto-merge."

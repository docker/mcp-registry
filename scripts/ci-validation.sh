#!/bin/bash

set -o pipefail

# Track overall success/failure and processed servers
overall_success=true
processed_servers=""

# Function to process a single server
process_server() {
  local file="$1"
  local dir=$(dirname "$file")
  local name=$(basename "$dir")
  
  echo "Processing server: $name"
  echo "================================"
  
  # Run each command and check for failures
  if ! task validate -- --name "$name"; then
    echo "ERROR: Validation failed for $name"
    return 1
  fi
  
  if ! task build -- --tools --pull-community "$name"; then
    echo "ERROR: Build failed for $name"
    return 1
  fi
  
  echo "--------------------------------"
  
  if ! task catalog -- "$name"; then
    echo "ERROR: Catalog generation failed for $name"
    return 1
  fi
  
  echo "--------------------------------"
  
  cat "catalogs/$name/catalog.yaml"
  
  echo "--------------------------------"
  echo "Successfully processed: $name"
  echo ""
  
  return 0
}

# Main loop - process each file but skip duplicate servers
while IFS= read -r file; do
  dir=$(dirname "$file")
  name=$(basename "$dir")

  # Skip if we've already processed this server (can happen when more than one file is changed for the same server)
  if [[ "$processed_servers" == *"|$name|"* ]]; then
    echo "Skipping already processed server: $name (from file: $file)"
    continue
  fi
  
  # Mark this server as processed
  processed_servers="${processed_servers}|$name|"
  
  if ! process_server "$file"; then
    echo "FAILED: Processing server from file: $file"
    overall_success=false
  fi
done < changed-servers.txt

# Exit with appropriate status code
if [ "$overall_success" = true ]; then
  echo "All servers processed successfully!"
  exit 0
else
  echo "One or more servers failed to process!"
  exit 1
fi
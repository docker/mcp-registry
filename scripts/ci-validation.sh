#!/bin/bash

set -o pipefail

# Track overall success/failure
overall_success=true

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

# Main loop
while IFS= read -r file; do
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
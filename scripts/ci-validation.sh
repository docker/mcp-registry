#!/bin/bash

set -o pipefail

# Parse arguments.
BIN_DIR=""
WORKSPACE=""
CHANGED_SERVERS_FILE="changed-servers.txt"

while [[ $# -gt 0 ]]; do
  case $1 in
    --bin-dir)
      BIN_DIR="$2"
      shift 2
      ;;
    --workspace)
      WORKSPACE="$2"
      shift 2
      ;;
    --changed-servers)
      CHANGED_SERVERS_FILE="$2"
      shift 2
      ;;
    *)
      echo "Unknown option: $1" >&2
      exit 1
      ;;
  esac
done

# If workspace is specified, change to it.
if [ -n "$WORKSPACE" ]; then
  cd "$WORKSPACE" || exit 1
fi

# Determine which binaries to use.
if [ -n "$BIN_DIR" ]; then
  VALIDATE_BIN="$BIN_DIR/validate"
  BUILD_BIN="$BIN_DIR/build"
  CATALOG_BIN="$BIN_DIR/catalog"
  CLEAN_BIN="$BIN_DIR/clean"
else
  VALIDATE_BIN="task validate --"
  BUILD_BIN="task build --"
  CATALOG_BIN="task catalog --"
  CLEAN_BIN="task clean --"
fi

# Track overall success/failure and processed servers.
overall_success=true
processed_servers=""

# Function to process a single server.
process_server() {
  local file="$1"
  local dir=$(dirname "$file")
  local name=$(basename "$dir")

  echo "Processing server: $name"
  echo "================================"

  # Run each command and check for failures.
  if ! $VALIDATE_BIN --name "$name"; then
    echo "ERROR: Validation failed for $name"
    $CLEAN_BIN "$name" >/dev/null 2>&1 || true
    return 1
  fi

  if ! $BUILD_BIN --tools --pull-community "$name"; then
    echo "ERROR: Build failed for $name"
    $CLEAN_BIN "$name" >/dev/null 2>&1 || true
    return 1
  fi

  echo "--------------------------------"

  if ! $CATALOG_BIN "$name"; then
    echo "ERROR: Catalog generation failed for $name"
    $CLEAN_BIN "$name" >/dev/null 2>&1 || true
    return 1
  fi

  echo "--------------------------------"

  cat "catalogs/$name/catalog.yaml"

  echo "--------------------------------"
  echo "Successfully processed: $name"
  echo ""
  if ! $CLEAN_BIN "$name"; then
    echo "WARNING: Cleanup encountered issues for $name"
  fi

  return 0
}

# Main loop - process each file but skip duplicate servers.
while IFS= read -r file; do
  dir=$(dirname "$file")
  name=$(basename "$dir")

  # Skip if we've already processed this server (can happen when more than
  # one file is changed for the same server).
  if [[ "$processed_servers" == *"|$name|"* ]]; then
    echo "Skipping already processed server: $name (from file: $file)"
    continue
  fi

  # Mark this server as processed.
  processed_servers="${processed_servers}|$name|"

  if ! process_server "$file"; then
    echo "FAILED: Processing server from file: $file"
    overall_success=false
  fi
done < "$CHANGED_SERVERS_FILE"

# Exit with appropriate status code.
if [ "$overall_success" = true ]; then
  echo "All servers processed successfully!"
  exit 0
else
  echo "One or more servers failed to process!"
  exit 1
fi

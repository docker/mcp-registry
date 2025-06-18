#!/bin/bash

set -e


NEW_SERVER_NAME=$(echo "$1" | tr '[:upper:]' '[:lower:]')
NEW_SERVER_NAME_UPPER=$(echo "$1" | tr '[:lower:]' '[:upper:]')

PATH_TO_SERVER="./servers/$NEW_SERVER_NAME"
CATEGORY=$(echo "$2" | tr '[:upper:]' '[:lower:]')

# IF IMAGE_NAME IS NOT PROVIDED, USE THE DEFAULT ONE
IMAGE_NAME=${IMAGE_NAME:-"mcp/$NEW_SERVER_NAME"}


if [[ -f "$NEW_SERVER_NAME" ]]; then
  echo "❌ File already exists: $NEW_SERVER_NAME"
  exit 1
fi

echo "Creating new server: $NEW_SERVER_NAME"

mkdir -p "$PATH_TO_SERVER"

echo "Server created successfully"

echo "Creating server.yaml"

echo "server:
  name: $NEW_SERVER_NAME
  image: $IMAGE_NAME
type: server
meta:
  category: $CATEGORY
  tags:
    - $CATEGORY
  highlighted: false
about:
  title: $NEW_SERVER_NAME
  icon: https://avatars.githubusercontent.com/u/182288589?s=200&v=4
source:
  project: $3
  branch: main
# config:
#   description: "TODO"
#   secrets:
#     - name: $NEW_SERVER_NAME.secret_name
#       env: $NEW_SERVER_NAME_UPPER
#       example: "TODO"   
#   env:
#     - name: ENV_VAR_NAME
#       example: "TODO"
#       value: '{{$NEW_SERVER_NAME.env_var_name}}'
#   parameters:
#     type: object
#     properties:
#       param_name:
#         type: string
#     required:
#       - param_name

    " > "$PATH_TO_SERVER/server.yaml"

echo "$NEW_SERVER_NAME created successfully"


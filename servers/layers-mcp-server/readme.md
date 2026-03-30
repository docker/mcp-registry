# Layers MCP Server

[Layers](https://layers.md) is a collaborative workspace platform for teams. This MCP server provides 57 tools to interact with your Layers workspace through AI assistants.

## What you can do

- **Workspaces** — List, resolve, view members and hierarchy
- **Projects** — Create, update, delete, view dashboard statistics and members
- **Tasks** — Full CRUD, smart filtering (by assignee, deadline, status), comments
- **Pages** — Read, create, edit, duplicate, move, get plain text content
- **Sprints** — Manage task lists and iterations
- **Folders & Forms** — Full CRUD operations
- **Search** — Full-text search across workspace

## Setup

1. Sign up at [layers.md](https://layers.md)
2. Go to Workspace Settings → Access Tokens
3. Create an access token
4. Connect via `https://app.layers.md/mcp`

## Documentation

- [Layers Platform](https://layers.md)
- [MCP Server Source](https://github.com/AppLayers/layers-mcp-server)

## Self-Hosted

For self-hosted Layers instances, use the Docker image:

```bash
docker pull sinups/layers-mcp-server:latest
```

Configure with your instance URL instead of `app.layers.md`.

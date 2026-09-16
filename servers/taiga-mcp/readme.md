# Taiga MCP Server

Full documentation: https://github.com/JohnWBlack/taiga-mcp

## Overview
Taiga MCP Server provides Model Context Protocol access to Taiga.io project management platform. Users deploy their own containerized instance connecting to their Taiga account.

## Deployment
Users run their own instance using Docker:

```bash
docker run -d \
  -e TAIGA_BASE_URL=https://api.taiga.io/api/v1 \
  -e TAIGA_USERNAME=your-email@example.com \
  -e TAIGA_PASSWORD=your-password \
  -e ACTION_PROXY_API_KEY=$(uuidgen) \
  -p 8000:8000 \
  ghcr.io/johnwblack/taiga-mcp:latest
```

Or deploy to Azure Container Apps, Kubernetes, or any container platform.

## Features
- **Projects**: List and retrieve project details
- **Epics**: Manage high-level features and initiatives
- **User Stories**: Create, update, list stories with search/filters
- **Tasks**: Break down stories into actionable tasks
- **Milestones**: Track sprints and release cycles
- **Users**: Find team members and assign work

## Authentication
Each deployment requires Taiga.io credentials (username/password) and a generated ACTION_PROXY_API_KEY for MCP client authentication.

## Use Cases
- Agile project management through AI agents
- DoD/SBIR program tracking automation
- Sprint planning and story creation
- Task assignment and status updates
- Integration with ChatGPT, Claude, and other MCP clients

## Links
- GitHub: https://github.com/JohnWBlack/taiga-mcp
- Taiga.io: https://taiga.io
- Contact: john.black@offset3.com

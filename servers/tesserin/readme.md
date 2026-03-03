# Tesserin MCP Server

For full documentation, setup instructions, and configuration options, see the project repository:

https://github.com/AnvinX1/Tesserin-pro

## Prerequisites

Tesserin must be running locally. The MCP server connects to the Tesserin REST API (default: `http://localhost:9960`).

Obtain your API token from **Settings → Connections → API Access** inside the Tesserin app.

## Quick start

```bash
export TESSERIN_API_TOKEN=tsk_your_token_here
docker run --rm -i \
  --add-host=host.docker.internal:host-gateway \
  -e TESSERIN_API_TOKEN \
  -e TESSERIN_API_URL=http://host.docker.internal:9960 \
  mcp/tesserin
```

Or with docker compose (see `docker-compose.yml` in the repository).

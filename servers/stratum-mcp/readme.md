# Stratum MCP (Remote)

Stratum MCP is a cloud-hosted MCP proxy that provides code analysis and developer intelligence tools over Streamable HTTP.

## Server

- **URL:** `https://mcp404-1.onrender.com/mcp`
- **Transport:** `streamable-http`
- **Auth:** `Authorization: Bearer <API_KEY>` (Stratum API key)

## Docker Desktop Setup

1. Open **Docker Desktop** → **MCP** (Servers / Registry).
2. Add a **Remote** MCP server:
   - URL: `https://mcp404-1.onrender.com/mcp`
   - Header: `Authorization: Bearer ${STRATUM_API_KEY}` (Docker secret substitution)
3. Configure the required secret (API key):
   - Secret name: `stratum.api_key`
   - Env: `STRATUM_API_KEY`

CLI alternative:

```bash
docker mcp secret set stratum.api_key=mcp_sk_...
```

## Quick Test (curl)

MCP uses Streamable HTTP sessions: call `initialize`, then forward the returned `MCP-Session-Id` header on subsequent calls.

```bash
URL="https://mcp404-1.onrender.com"
API_KEY="mcp_sk_..."

curl -sS -D /tmp/mcp.headers -o /tmp/mcp.init.json -X POST "$URL/mcp" \
  -H "Authorization: Bearer $API_KEY" \
  -H "Accept: application/json, text/event-stream" \
  -H "Content-Type: application/json" \
  -H "MCP-Protocol-Version: 2025-11-25" \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-11-25","capabilities":{},"clientInfo":{"name":"curl","version":"1.0"}}}'

SESSION_ID="$(awk 'BEGIN{IGNORECASE=1} $1 ~ /^mcp-session-id:$/ {sid=$2} END{print sid}' /tmp/mcp.headers | tr -d '\r')"

curl -sS -X POST "$URL/mcp" \
  -H "Authorization: Bearer $API_KEY" \
  -H "Accept: application/json, text/event-stream" \
  -H "Content-Type: application/json" \
  -H "MCP-Session-Id: $SESSION_ID" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}'
```

## Available Tools

Tools are discovered dynamically via `tools/list`. Common built-ins include:

- `cleanup_thinking_sessions`
- `dev_search`
- `explain_architecture`
- `extract_css`
- `find_code`
- `generate_thinking`
- `get_project_context`
- `get_project_structure`
- `get_thinking`
- `list_thinking_sessions`
- `read_code`
- `update_thinking`
- `web_read`

## Support

- Email: `support@stratum-mcp.com`
- Service status / base URL: `https://mcp404-1.onrender.com/`

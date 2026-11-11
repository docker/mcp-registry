# Instadata (remote MCP)
Documentation: https://instadata.dev/docs/mcp-server/
Endpoint: https://mcp.instadata.dev (Streamable HTTP).
Create a key at https://instadata.dev/app/keys and configure the `instadata.api_key` secret.
Authentication is `Authorization: Bearer <key>`, not `X-API-Key`.
$5 per 1,000 credits; 25 free credits. Data tool calls can consume credits.
This draft uses static header authentication; it does not assert a tested OAuth connection.

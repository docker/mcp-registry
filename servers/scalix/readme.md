# Scalix Cloud MCP Server

Documentation: https://docs.scalix.world/mcp

The hosted MCP server for **Scalix Cloud**, the agent-native cloud platform: 53 tools across Postgres databases, serverless functions, AI inference, object storage, persistent Linux computers, containers, KV, DNS, and events.

- **Discovery (`initialize`, `tools/list`):** Unauthenticated
- **Tool execution (`tools/call`):** Supports OAuth 2.1 (RFC 9728 / RFC 8414) for direct MCP clients or a Bearer API key (`SCALIX_API_KEY`, RFC 6750 challenge on 401)
- **Official MCP Registry ID:** `world.scalix/cloud`

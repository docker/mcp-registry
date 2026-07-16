# Maguyva

Maguyva is a hosted (remote) MCP server that provides code intelligence for coding agents — semantic, AST, graph, and full-text search, symbol lookup, dependency analysis, and one-call task context across 279+ languages. No local indexer is required.

- Website & docs: https://maguyva.ai
- MCP endpoint: https://maguyva.tools/mcp (transport: `streamable-http`)

## Authentication

This is a remote server that uses **API-key** authentication (not OAuth). Provide your key via the `Authorization: Bearer <API_KEY>` or `X-API-Key: <API_KEY>` header. A free tier is available — sign up at https://maguyva.ai.

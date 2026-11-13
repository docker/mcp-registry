# KnowMan

KnowMan is a knowledge management app: save links and articles, then search, read and ask questions over your personal knowledge base with cited answers.

This entry is the hosted remote MCP server at `https://api.knowman.io/mcp` (Streamable HTTP). It exposes six tools: `kb_search`, `kb_ask`, `kb_save`, `kb_get_item`, `kb_recent` and `kb_memories`.

## Authentication

- **API key:** create a key in KnowMan under Settings, then API keys, and set it as `KNOWMAN_API_KEY`. The key is sent as `Authorization: Bearer <key>`.
- **OAuth 2.1:** MCP clients that support OAuth can connect from the endpoint URL alone (protected resource metadata at `https://api.knowman.io/.well-known/oauth-protected-resource/mcp`).

## Links

- Website: https://knowman.io
- Source for the local stdio package (`@knowman/mcp`): https://github.com/Essatt/knowman-mcp
- Documentation: https://github.com/Essatt/knowman-mcp#readme

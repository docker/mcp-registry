# Synap Memory

Persistent memory for AI agents over MCP (Streamable HTTP). Give any MCP-compatible
agent long-term memory with just an MCP URL and a Synap API key — no SDK, no code.

- Homepage: https://www.maximem.ai/synap
- Docs: https://www.maximem.ai/docs
- Source: https://github.com/maximem-ai/maximem_synap_sdk/tree/main/packages/mcps/synap-mcp-server

## Authentication

Send your Synap API key as a Bearer token in the `Authorization` header:

    Authorization: Bearer <YOUR_SYNAP_API_KEY>

Get a key from your dashboard at https://synap.maximem.ai.

## Tools

- `log_exchange` — save a conversation exchange to memory
- `recall_context` — recall the most relevant context for the current turn
- `list_recent_memories` — list recent memories (debug / "test my memory")
- `check_memory_status` — check the ingestion status of a logged exchange

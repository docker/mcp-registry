# TBit MCP Server

TBit (https://tbit.app) is a conversational AI platform where businesses run AI agents on WhatsApp and Instagram. This remote MCP server lets ChatGPT, Claude or any MCP client read, configure and train the user's own TBit agents: product catalog, memories, prompt drafts, recent conversations, forms, pipeline and schedule.

- Server URL: `https://api.tbit.app/mcp` (Streamable HTTP, OAuth 2.1 with Dynamic Client Registration)
- Documentation: https://tbit.app/mcp and https://tbit.app/es/content/conecta-tu-ia-mcp
- Permissions are role-based per agent: read with any role, training only where the user is admin. Prompt changes land as drafts the owner approves in the TBit app.

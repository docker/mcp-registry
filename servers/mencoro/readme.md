# Mencoro MCP Server

Brand visibility in AI answers: rank, mentions, sentiment, Share of Voice, cited
sources and competitor head-to-heads across ChatGPT, Perplexity, Google AI
Overview and AI Mode, plus Google Search and Shopping. Every tool is read-only.

- Docs: https://mencoro.com/features/mcp-server/
- Source: https://github.com/mencoro/mencoro-mcp

Authentication is OAuth 2.1 with PKCE S256. The server advertises both Client ID
Metadata Documents and Dynamic Client Registration. A read-only personal access
token sent as `Authorization: Bearer mcp_pat_...` also works.

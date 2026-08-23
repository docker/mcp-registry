# Keenable Web Search

Remote MCP server for live web search and page-content fetch over the Keenable
web index.

- Endpoint: `https://api.keenable.ai/mcp` (Streamable HTTP)
- Docs: https://docs.keenable.ai/mcp-server
- Tools: `search_web_pages`, `fetch_page_content`
- Keyless by default (1,000 requests/hour); an optional `X-API-Key` lifts the cap.

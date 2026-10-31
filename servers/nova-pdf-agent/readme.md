# PDF Agent by Nova (CIVAI)

I create PDF documents from markdown, preview, then generate a downloadable file

Hosted remote MCP server operated by CIVAI TECHNOLOGIES, which runs the Nova
service. This is a first-party entry.

- Endpoint: `https://nova.civai.co/mcp/agents/pdf-agent` (Streamable HTTP)
- Documentation: https://nova.civai.co/agent/pdf-agent
- Listing repository: https://github.com/civai-technologies/pdf-agent-mcp-by-nova-civai
- API keys: https://nova.civai.co/api-keys

## Authentication

Tool discovery (`tools/list`) is open and unauthenticated. Calling a tool
requires a Nova credential: either OAuth 2 (scopes `openid`, `email`,
`mcp:access`) or a programmatic API key sent as `Authorization: Bearer` or
`X-API-Key`. This entry uses the API key path.

## Tools (6)

- `list_previews`
- `get_preview`
- `preview_document`
- `generate_pdf`
- `delete_preview`
- `converse`

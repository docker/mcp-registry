# Slidegen Agent by Nova (CIVAI)

I do everything related to image and slide asset generation

Hosted remote MCP server operated by CIVAI TECHNOLOGIES, which runs the Nova
service. This is a first-party entry.

- Endpoint: `https://nova.civai.co/mcp/agents/slidegen-agent` (Streamable HTTP)
- Documentation: https://nova.civai.co/agent/slidegen-agent
- Listing repository: https://github.com/civai-technologies/slidegen-agent-mcp-by-nova-civai
- API keys: https://nova.civai.co/api-keys

## Authentication

Tool discovery (`tools/list`) is open and unauthenticated. Calling a tool
requires a Nova credential: either OAuth 2 (scopes `openid`, `email`,
`mcp:access`) or a programmatic API key sent as `Authorization: Bearer` or
`X-API-Key`. This entry uses the API key path.

## Tools (4)

- `generate_single_graphic`
- `generate_image`
- `generate_multiple_graphics`
- `converse`

# MarketPounce

Hosted remote MCP server for the MarketPounce growth desk.

## Documentation

- Docs: https://www.marketpounce.com/docs/mcp
- Marketing: https://www.marketpounce.com/mcp
- Privacy: https://www.marketpounce.com/privacy

## Authentication

MarketPounce uses Clerk OAuth 2.1 for interactive clients. Optional brand API keys (`mpb_`) are available for server to server use. This catalog entry does not invent a personal access token env var because that would misrepresent how auth works.

Connect with any MCP client that supports remote Streamable HTTP and OAuth. After login, call `list_brands`, then `set_brand`.

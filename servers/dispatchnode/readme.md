# DispatchNode

Hosted Streamable HTTP MCP for field service operators.

- Docs: https://www.dispatchnode.com/docs/mcp
- Product: https://www.dispatchnode.com/mcp
- Privacy: https://www.dispatchnode.com/privacy

## Auth

DispatchNode uses Clerk OAuth 2.1 for marketplace clients. Connect with Claude, ChatGPT, Cursor, or any MCP client that supports OAuth against `https://www.dispatchnode.com/api/mcp`. Optional `dnk_` operator API keys exist for S2S automation; they are not personal access tokens for Docker Desktop OAuth.

After login, call `list_orgs`, then `set_organization`. The live OAuth catalog is 17 tenant tools covering calls, jobs, callbacks, and partners.

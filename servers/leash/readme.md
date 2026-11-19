# Leash

[Leash](https://lea.sh) is a free hosted MCP gateway for protected access to the personal accounts each user connects. It provides configurable checks, per-AI permissions, activity history and undo for supported reversible actions. Only activity routed through Leash is covered.

Connect to `https://secure.lea.sh/mcp`. The endpoint supports OAuth 2.1 with dynamic client registration and S256 PKCE. Authorize your own Leash account, then connect the services you want in [your personal space](https://lea.sh/personal). No shared test or user credentials are included in this catalog.

The [public integration repository](https://github.com/open-leash/leash-plugin) contains setup instructions, metadata, rules and skills. The hosted gateway implementation is proprietary. The server requires authentication for account and tool access. A public availability check is at `https://secure.lea.sh/health/ready`.

After authorizing, call `leash_check_connection` with `input: {}`, a concise `reason`, and task `context` to verify the connection without reading email or files. Confirm success only from a successful tool result.

[Privacy](https://lea.sh/privacy) | [Terms](https://lea.sh/terms) | [Support](https://lea.sh/support)

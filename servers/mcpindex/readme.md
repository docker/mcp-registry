# mcpindex

The directory client for [mcpindex.ai](https://mcpindex.ai). Find MCP servers by
natural-language task, and get advisory trust screens for a server or tool before
you connect. Advisory, not a safety verdict (the in-path gate is mcpindex-gate).

## Tools
- `recommend_mcp_for_task` - top MCP servers for a task, with reasoning + install commands.
- `search_mcp_servers` - keyword + semantic search across the registry.
- `get_install_command` - exact install command for a server + client.
- `compare_servers` - side-by-side comparison of 2 to 5 servers.
- `check_tool_trust` - advisory pre-invocation screen for a specific tool.
- `assess_server` - aggregated advisory trust assessment for a whole server.

Zero-config: it queries the public mcpindex.ai API; no credentials required.

More: https://mcpindex.ai

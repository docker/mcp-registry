# Cipher & Row MCP Server

Freight counterparty verification for AI agents. The Cipher & Row MCP server lets agents verify US and Canadian freight carriers and brokers before doing business with them: live FMCSA safety and authority data, multi-signal trust scores, registry search, and continuous partner watchlist monitoring.

## Tools

8 read-only tools, discovered dynamically: `cr_verify_carrier`, `cr_verify_broker`, `cr_search_registry`, `cr_session_status`, `cr_add_partner`, `cr_remove_partner`, `cr_list_partners`, `cr_set_notifications`.

## Authentication

OAuth 2.1 with PKCE and dynamic client registration (RFC 7591), discovered via `https://mcp.cipherandrow.com/.well-known/oauth-authorization-server`. No pre-shared secrets required; MCP clients register and authorize through the standard OAuth connector flow.

## More Information

- Setup guide: https://www.cipherandrow.com/mcp/setup
- Docs and JSON schemas: https://github.com/cipherandrowhq/cipherandrow-mcp
- Server card: https://mcp.cipherandrow.com/.well-known/mcp

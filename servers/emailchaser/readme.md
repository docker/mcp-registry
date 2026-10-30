# Emailchaser

Emailchaser is a cold email platform. This remote MCP server exposes the
Emailchaser public API as tools, so an agent can create and launch campaigns,
add, update and source leads, read and answer replies, check inbox placement
and mailbox health, manage sending accounts, webhooks and workspaces, and run
autopilot campaigns.

- Documentation: https://www.emailchaser.com/mcp-server
- API reference: https://run.emailchaser.com
- Endpoint: `https://app.emailchaser.com/api/mcp` (Streamable HTTP)

## Authentication

Create an API key in Emailchaser under **API & MCP** in the side menu and
send it as `Authorization: Bearer <key>`. Keys carry a read or read & write
scope, enforced by the Emailchaser backend, and can be revoked at any time.

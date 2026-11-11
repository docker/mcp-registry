# smart-me

Remote MCP server for the [smart-me](https://web.smart-me.com/) energy platform — smart metering,
EV charging and ZEV (tenant) billing.

- **Documentation:** https://github.com/eCarUp/smart-me-mcp
- **Endpoint:** `https://mcp.smart-me.com/mcp` (Streamable HTTP)
- **Authentication:** OAuth 2.1 with dynamic client registration (RFC 7591) and PKCE. The client
  discovers the flow from the server's `401` response, so there is no API key to provision.
- **Account:** a [smart-me](https://web.smart-me.com/) account. Quarter-hourly load profiles and
  long series need a Professional licence; everything else works on any account.

## What it does

Reads devices and their current readings, quarter-hourly load profiles and daily series, the
folder tree of a site, charging stations with their sessions, settings and load-management groups,
and the tariffs, invoice positions and consumption of a billing property.

Writes are annotated so the client asks first: renaming meters, creating and moving folders,
switching a relay, sending a command to a charging station, setting current limits, and creating
or changing billing properties, tariffs and VEWA cost periods. Deleting a folder, charging station
or load-management group takes two steps — the first previews, the second needs the exact name.

## Support

[smart-me helpdesk](https://odoo.smart-me.com/helpdesk)

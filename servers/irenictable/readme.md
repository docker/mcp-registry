# IrenicTable

Collaborative digital whiteboard for non-profit organizations in the likeness/spirit of the superstar commercial apps, Mural and Miro.

## What it does

IrenicTable lets AI assistants create and manipulate visual workspaces in real time. The MCP server exposes 38 tools across board management, object CRUD, C4 architecture diagramming, mind mapping, spatial layout, and tagging.

## Documentation

- [Remote Setup Guide](https://irenictable.benevolabs.org/docs/mcp-remote-setup)
- [OAuth Metadata](https://irenictable.benevolabs.org/.well-known/oauth-authorization-server)
- [Health Check](https://irenictable.benevolabs.org/mcp/health)

## Authentication

OAuth 2.0 bearer tokens. The server publishes standard OAuth authorization server metadata at `/.well-known/oauth-authorization-server`. Clients that support MCP OAuth with PKCE can acquire tokens automatically. A device authorization flow (RFC 8628) is also available for headless clients.

## Transport

Streamable HTTP at `https://irenictable.benevolabs.org/mcp`

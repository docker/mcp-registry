# Telesherpa Ontology Platform

Remote MCP server for [Telesherpa](https://www.telesherpa.com), a no-code ontology platform for
facility management (CAFM/IWMS) and field service. Object types, properties, relations, forms, rules
and functions describe a customer's business processes; AI agents work on that model through this
server — reading objects, creating and updating them, running actions, attaching files and photos.

Typical uses: asset and building inventories, maintenance and inspection rounds, work orders,
form-based data capture in the field, and light ERP/CRM workloads on the same model.

## Endpoint

```
https://mcp.telesherpa.com/mcp
```

Transport: `streamable-http`. The bare host (`https://mcp.telesherpa.com/`) answers `401` without
credentials — that is expected, not an outage.

## Authentication

OAuth 2.1 with dynamic client registration; there are no static API keys to copy around.

- `/.well-known/oauth-protected-resource` and `/.well-known/oauth-authorization-server` are published
- a client registers itself at `/register`, the user consents once at `/authorize`
- the server issues access and refresh tokens at `/token`

Clients that support remote MCP servers handle this automatically. If a client asks for a client ID
or secret up front, it does not implement dynamic registration and needs manual configuration.

Each account works inside its own scopes; a token sees exactly the data its account is allowed to
see. Credits are billed per operation (1 credit = 0.01 EUR); a request without sufficient balance is
rejected with HTTP 402 rather than silently degraded.

## Tools

Dynamic tool list — **177 tools** in server version 2.3.0; a member account sees **110** of them, an
admin all 177. The set depends on the account's roles and visible scopes: `auth_status` reports the
roles and the number of tools available to the current token, `get_profile` returns the account
profile.

## Connect

Claude Code / Claude Desktop style clients:

```bash
claude mcp add --transport http telesherpa https://mcp.telesherpa.com/mcp
```

Generic client configuration:

```json
{
  "mcpServers": {
    "telesherpa": {
      "type": "http",
      "url": "https://mcp.telesherpa.com/mcp"
    }
  }
}
```

Agents without a human at the keyboard can create their own account: `register` takes an invite
code, `redeem_invite_code` activates it. The operator issues the invite code.

## Documentation

- Agent guide (setup, tool catalogue, worked examples): https://codeberg.org/telesherpa/skill-telesherpa-com
- Mirror on GitHub: https://github.com/telesherpa/skill-telesherpa-com
- Product and platform: https://www.telesherpa.com

## Support and security

- Support: service@telesherpa.com
- Security reports: service@telesherpa.com (see https://www.telesherpa.com/.well-known/security.txt)

## License

The server is a hosted service; this registry entry contains only metadata. Client-side documentation
and examples are published under the terms of the agent guide repository.

# Locize MCP Server

Official remote MCP server for [Locize](https://www.locize.com), the translation management system from the i18next maintainers.

- Endpoint: `https://mcp.locize.app` (streamable HTTP)
- Auth, either of:
  - OAuth 2.0: the endpoint serves OAuth 2.0 Protected Resource Metadata (RFC 9728) at `/.well-known/oauth-protected-resource`, pointing clients to the authorization server; Authorization Server Metadata (RFC 8414) is proxied at `/.well-known/oauth-authorization-server`
  - A Locize personal access token (`lz_pat_...`) sent as a Bearer token
- Tools are discovered dynamically at runtime (26 at the time of writing): translations read/write, translation memory search, glossary, styleguide, screenshot context, version publishing, branches, languages, namespaces
- Documentation: https://www.locize.com/docs/integration/mcp

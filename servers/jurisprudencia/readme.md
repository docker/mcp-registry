Docs: https://jurisprudencia.io/mcp/

# Jurisprudencia.io

Search Spanish case law and retrieve available judgment passages with links to official sources. Tools distinguish source summaries from full judgment text and preserve document versions. Coverage is progressive, not exhaustive.

## Connection

- Streamable HTTP endpoint: https://jurisprudencia.io/mcp
- A Jurisprudencia.io account and authorization are required.
- The service supports OAuth authorization code with PKCE and dynamic client registration. A valid account API token is also accepted as a Bearer token.
- OAuth authorization metadata: https://jurisprudencia.io/.well-known/oauth-authorization-server
- Protected resource metadata: https://jurisprudencia.io/.well-known/oauth-protected-resource
- Account usage limits apply to searches; protocol initialization and discovery do not count as searches.

Tool definitions are discovered dynamically. They include court and source discovery, judgment search, lookup and passage reading, and separate search/lookup/reading tools for tax administrative documents.

Official MCP Registry name: `io.jurisprudencia/jurisprudencia`.

Technical support: developer@jurisprudencia.io

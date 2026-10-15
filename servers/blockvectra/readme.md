# BlockVectra

Hosted MCP endpoint: https://docs.blockvectra.com/mcp (Streamable HTTP).

This Docker catalog entry provides keyless documentation, chain discovery, and JSON-RPC methods permitted by each chain’s public policy.

For authenticated RPC, Data API access, or signed-transaction broadcasting, consult the [BlockVectra MCP setup guide](https://docs.blockvectra.com/en/guides/mcp-server/) and use a client that explicitly supports configuring a private `x-api-key` header. Those capabilities are not configured by this Docker catalog entry.

A separate local stdio option is the published [`@blockvectra/mcp@0.1.0`](https://www.npmjs.com/package/@blockvectra/mcp) package, run with `npx --yes @blockvectra/mcp@0.1.0`. It requires Node.js 22+ and network access. See the setup guide for its configuration and capabilities.

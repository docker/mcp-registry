# Vivu

[Connection instructions](https://app.vivu.ai/help/connectors)

Vivu lets an MCP client search a video library using natural language and retrieve relevant moments.

The remote endpoint is `https://mcp.vivu.ai/mcp` and uses Streamable HTTP. A Vivu account and OAuth authorization are required.

OAuth authorization-server metadata is available at `https://mcp.vivu.ai/.well-known/oauth-authorization-server`. It advertises dynamic client registration and PKCE with S256. Tools are discovered dynamically after authorization.

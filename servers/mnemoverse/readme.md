# Mnemoverse Memory

Hosted persistent memory for AI agents over MCP (streamable HTTP). One account gives an agent the same memory across Claude, Cursor, VS Code, ChatGPT and other MCP clients.

- Sign-in: OAuth 2.1 with PKCE and dynamic client registration. No API key in the client config.
- Tools are discovered when the client connects (dynamic tools), so this entry ships an empty tools.json.
- Access is scoped to the signed-in account. A client granted only memory:read can recall but not write.

Setup: https://mnemoverse.com/docs/api/remote-mcp-server
Setup by client: https://mnemoverse.com/docs/api/install-by-host
Privacy policy: https://mnemoverse.com/docs/legal/privacy-policy
Security: https://mnemoverse.com/docs/api/security

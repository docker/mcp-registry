Docs: https://maxibestof.one/mcp

Connect with OAuth in an MCP client that supports Streamable HTTP. An active MaxiBestOf Pro account is required.

The public repository contains connection configuration and documentation. The server runs as a hosted service.

After importing the catalog and enabling the server, run `docker mcp oauth authorize maxibestof` to sign in. Docker discovers the OAuth registration and authorization endpoints from the hosted service.

The `personal_access_token` names in the catalog are Docker OAuth credential storage names. This service does not require users to create a static personal access token.

Docs: https://auzzura.com

Sign-up: https://dm.auzzura.com
MCP endpoint (Streamable HTTP): https://api.dm.auzzura.com/mcp

DM is shared context for human + AI teams. Before a task, a connected agent asks DM for the approved Rules, Skills and Decisions that apply. After the task, the agent submits what it learned as a proposal, and a person decides what becomes shared.

Authentication is OAuth 2.0 with authorization-server discovery and dynamic client registration. Users sign in at DM; they do not mint a personal access token by hand. The `oauth` block in `server.yaml` follows the Docker MCP Catalog remote-wizard convention (see servers such as Linear and Notion) so Docker Desktop can run `docker mcp oauth authorize dm-by-auzzura`.

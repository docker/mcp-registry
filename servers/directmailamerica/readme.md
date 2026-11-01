Docs: https://directmailamerica.com/docs/mcp?utm_source=docker&utm_medium=referral&utm_campaign=mcp_directories

Publisher: DirectMailAmerica. The hosted service is proprietary and subject to
https://directmailamerica.com/terms. The registry entry is contributed under this
repository's MIT license; that license does not cover the hosted service or its
private implementation.

Connect an existing workspace through the OAuth consent page on
https://directmailamerica.com. Enter the workspace key only on that page; Docker
receives a separate revocable OAuth token. Choose only the permissions needed.

Use a Gateway version that sends the RFC 8707 `resource` parameter for both
authorization and token exchange. Gateway v0.34.0 bundled with Docker Desktop
4.55.0 omits it and cannot complete this connection. Current standalone Gateway
v0.44.3 supports it with `DOCKER_MCP_USE_CE=true`. An old cached DCR registration
may need to be revoked and recreated after upgrading so discovery uses the Narrow
endpoint. The endpoint is `https://directmailamerica.com/api/mcp-narrow`.

Website: https://directmailamerica.com/?utm_source=docker&utm_medium=referral&utm_campaign=mcp_directories

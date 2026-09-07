# ZopDev MCP

Cloud cost and infrastructure governance across AWS, Azure, GCP, Databricks and Snowflake.
289 tools (165 read, 124 write) covering cost, resources, schedules, recommendations,
budgets, governance and diagnostics over a hosted remote MCP endpoint. Read-only by default, with
optional scoped writes.

One server covers both ZopDev products — ZopNight (cost and FinOps) and ZopDay (build and
deploy). Hosted and remote: nothing to install, no npm package or Docker image required.

- server link - https://api.zop.dev/mcp-server
- learn doc - https://zop.dev/learn/mcp-server?utm_source=docker-mcp-registry&utm_medium=listing&utm_campaign=mcp-directory
- claude learn - https://zop.dev/learn/how-to/set-up-zopnight-mcp-for-claude
- public repo - https://github.com/zopdev/mcp

**Documentation:** https://zop.dev/learn/mcp-server?utm_source=docker-mcp-registry&utm_medium=listing&utm_campaign=mcp-directory

**Authentication:** OAuth 2.1 discovery is live, so most clients sign in with nothing to copy.
CI and non-browser clients send `Authorization: Bearer zn_pat_…`, a Zop personal access token
created under Settings → Organisation → Developer Settings (Admin or Editor role required).

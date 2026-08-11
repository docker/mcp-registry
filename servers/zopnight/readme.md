# Zopnight

Read-only cloud cost and infrastructure governance across AWS, Azure and GCP. 85 tools covering
cost overview and trends, cost by provider/resource/tag/team, budgets, resources and resource
groups, schedules and overrides, recommendations, tagging policies, audit logs, anomalies,
Kubernetes resources and pod logs.

Hosted remote server — nothing to install, no npm package or Docker image required. Read-only:
all write operations are refused with `mcp_write_not_allowed`.

- server link - https://api.zop.dev/mcp-server
- learn doc - https://zop.dev/learn/mcp-server?utm_source=docker-mcp-registry&utm_medium=listing&utm_campaign=mcp-directory
- claude learn - https://zop.dev/learn/how-to/set-up-zopnight-mcp-for-claude

**Documentation:** https://zop.dev/learn/mcp-server?utm_source=docker-mcp-registry&utm_medium=listing&utm_campaign=mcp-directory

**Authentication:** `Authorization: Bearer zn_pat_…` — a Zopnight personal access token, created
under Settings → Organisation → Developer Settings (Admin or Editor role required).

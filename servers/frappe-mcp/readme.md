# Frappe MCP Server

MCP server for [Frappe Framework](https://frappeframework.com) — interact with any Frappe/ERPNext site via its REST API.

## Documentation

- [README](https://github.com/muthanii/frappe_mcp/blob/main/README.md)
- [Docker Hub](https://hub.docker.com/r/muthanii/frappe-mcp)

## Features

- **Document CRUD** — get, create, update, delete Frappe doctypes
- **Search** — full-text and filtered document search
- **Remote method calls** — invoke any whitelisted server-side Python method
- **7 MCP tools** — `frappe_ping`, `frappe_get_doc`, `frappe_search_docs`, `frappe_create_doc`, `frappe_update_doc`, `frappe_delete_doc`, `frappe_run_method`

## Quick start

```bash
docker run -e FRAPPE_URL=https://your-site.com \
           -e FRAPPE_API_KEY=your-api-key \
           -e FRAPPE_API_SECRET=your-api-secret \
           muthanii/frappe-mcp
```

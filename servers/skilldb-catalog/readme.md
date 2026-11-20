# SkillDB Catalog

[Setup documentation](https://skilldb.dev/mcp) | [Browse public skills](https://skilldb.dev/skills) | [Privacy](https://skilldb.dev/privacy)

Search reusable AI agent skills by task, compare short public previews, and follow original source links before deciding what to use.

This hosted MCP connection uses Streamable HTTP at `https://skilldb.dev/api/mcp/catalog` and requires no authentication. It exposes only `skilldb_search` and `skilldb_get_preview`. Tools are discovered from the remote service.

The catalog returns public search results and short excerpts. It does not return complete skill content, install or execute skills, link accounts, or access private/team skills. The separate authenticated full MCP service is outside this listing.

Queries and selected public skill IDs are sent to SkillDB. Hosting infrastructure receives request metadata, including IP addresses. See the linked privacy policy for retention and data handling.

[Public MIT adapters](https://github.com/latentsmurf/skilldb-plugins) contain connection manifests, discovery instructions, documentation, and branding. They are not the hosted server source and do not relicense separately hosted catalog content.

Support and security reports: `dev_chad@skilldb.dev`.

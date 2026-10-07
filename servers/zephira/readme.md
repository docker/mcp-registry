# Zephira Company Intelligence

Connect to Zephira's first-party hosted Streamable HTTP endpoint for six read-only tools: company search, company profiles, officers, shareholders, corporate groups and available financial statements. Available source provenance is preserved; coverage varies by company and jurisdiction.

Docs: https://zephira.ai/developers/mcp/

Client source and setup examples: https://github.com/sens663/zephira-mcp

Endpoint: https://dashboard.zephira.ai/api/mcp

Set the `zephira.api_key` secret to your own active production dashboard MCP key (`zph_live_...`). The catalog sends it only as `Authorization: Bearer <key>` to the canonical endpoint. API v2 Token credentials are separate. This endpoint currently uses API-key authentication rather than OAuth.

Use `company:read`, `ownership:read` and `financials:read` scopes for the respective tools. Discovery and tool listing do not consume data credits; company-data calls use the user's dashboard allowance. Search by jurisdiction and name/identifier, then use the returned numeric company ID as a string. Officers and shareholders are paginated. Preserve source/modelled labels in group results.

Tool metadata is derived from the public discovery card at https://dashboard.zephira.ai/.well-known/mcp/server-card.json; the hosted endpoint is authoritative. The MIT license in the client repository applies to its client code and documentation, while the hosted service and company data have separate service terms.

Support and security reports: support@zephira.ai. Do not include secrets or private account data in public issues.

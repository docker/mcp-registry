# AaaSagent MCP

Remote MCP server indexing open-source agent prompt templates.

- Endpoint: https://aaasagent.xyz/api/v1/mcp
- Transport: streamable-http
- Authentication: none required (read-only, no API key)
- Documentation: https://aaasagent.xyz/for-agents
- Tools: `list_templates`, `get_template`
- Source repo: https://github.com/AaaSagent/aaasagent-mcp

## Licence note

This server **indexes** templates from 8 upstream GitHub repositories and keeps each
template's original licence and source link. 17 of 649 entries are metadata-only
(4 CC BY-NC non-commercial, 13 with no LICENSE file) and their text is not redistributed.
The server code itself is MIT.

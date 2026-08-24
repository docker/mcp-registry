# MCHS NWAU Calculator MCP

MCHS exposes a stdio MCP server for calculator schemas, validation, support
status, and release evidence. The MCP adapter is intentionally thin and does not
duplicate calculator formula logic.

The local Docker path runs `mchs-mcp` inside a container. Discovery and read-only
metadata calls do not require secrets. Do not send private healthcare data,
patient-level records, or institutional costing submissions to public catalog
validation environments.

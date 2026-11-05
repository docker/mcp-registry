# Datadef

Generate, edit, and export data-architecture diagrams from any MCP client. Describe a pipeline ("Kafka → Flink → Iceberg with the marts on top") and the server returns an editable diagram — typed tables with columns, column-level lineage, 2,000+ real tool icons — plus a PNG inline in the chat.

**Authentication:** the server requires `Authorization: Bearer dd_live_...`. Create an API key at [datadef.io/settings/mcp](https://datadef.io/settings/mcp) — the 7-day free trial includes MCP access.

**Tools:** outcome-level (`create_diagram`, `list_diagrams`, `get_diagram`, `edit_diagram`, `export_diagram`, `get_design_guide`) plus 25 atomic `canvas_*` tools the calling model can drive directly. A `datadef_design_guide` prompt teaches any model the design standard before it draws.

Documentation: [datadef.io/guides/en/mcp-diagram-server](https://datadef.io/guides/en/mcp-diagram-server)

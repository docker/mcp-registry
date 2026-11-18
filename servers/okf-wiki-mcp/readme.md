# OneMira OKF & Obsidian Wiki MCP Server

Zero-dependency, 100% offline Model Context Protocol (`MCP` JSON-RPC 2.0) server for querying local **Google OKF v0.2 (Open Knowledge Format)** and **Obsidian** Markdown meeting vaults in Claude Desktop, Claude Code, Cursor, Windsurf, and Docker MCP Toolkit.

## Tools Provided

- `okf_search_meetings` — Search local OKF v0.2 and Obsidian Markdown meeting transcripts and notes by keyword, speaker, tag, and date range (`100%` offline).
- `okf_list_action_items` — Extract open or completed Markdown task items (`- [ ]` / `- [x]`) and `@owner` assignments across local meeting notes.
- `okf_read_note` — Read a specific local OKF/Obsidian Markdown note by path, slug, or `[[wiki-link]]`, including YAML frontmatter and incoming backlinks.
- `okf_topic_graph` — Build the relational `[[wiki-link]]` topic graph and audit unresolved links across the local Markdown vault.
- `okf_vault_stats` — Audit vault health: total notes, YAML frontmatter coverage, unique speakers, wiki-link edges, and open/completed action items.

## Documentation & Source

- MCP Integration Guide: https://onemira.com/docs/mcp-integration
- OneMira Local AI Context System: https://onemira.com
- Source Repository (MIT): https://github.com/onemira-space/okf-wiki-mcp

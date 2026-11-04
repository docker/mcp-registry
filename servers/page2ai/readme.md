# Page2AI — Web to Markdown

Convert web pages or HTML to clean Markdown for LLM context. Preserves code
blocks with language hints, tables, and reading structure; strips ads,
navigation, and cookie banners.

- **Source**: https://github.com/igorsaevets/page2ai-mcp
- **Documentation**: https://igorsaevets.github.io/page2ai-docs/
- **Hosted endpoint**: https://page2ai-mcp-remote.vercel.app/api/mcp
- **npm alternative** (stdio): `npx -y @page2ai/mcp` (Node 22+)
- **License**: MIT
- **Author**: Igor Saevets (ORCID `0009-0006-8636-1377`)

## Tools

`page_to_markdown(url, include_images?, include_frontmatter?, timeout_ms?)` —
fetch a URL and return clean Markdown. Blocks private ranges, loopback, and
cloud metadata endpoints (GHSA-fx68-fhhj-5874 hardened).

## Handshake

Protocol: `2025-11-25` (default) and `2026-07-28` (opt-in). Session ID is
optional. No auth. Anonymous callers get a real response.

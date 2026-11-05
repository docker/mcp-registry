Docs: https://lucydomi.com/mcp-guide

BusyBee exposes public social media data to AI agents through ten uniform MCP tools:

- `list_platforms` — which platforms the API key can reach
- `search_creators`, `search_posts` — find accounts and posts by keyword
- `get_creator`, `list_creator_posts`, `get_post`, `get_comments` — one account, one post, or its comment thread
- `get_comment_insights` — sentiment split, question ratio, top phrases and a one-line summary from a comment sample
- `get_trending` — trending board
- `check_balance` — remaining credit

Platforms: TikTok, Instagram, YouTube, X (Twitter), Douyin, Xiaohongshu. One schema across all six, so an agent can switch platform by changing one field.

Remote endpoint: `https://lucydomi.com/mcp` (streamable HTTP). Authentication is an API key sent as `Authorization: Bearer sk-...`; get a free key at https://lucydomi.com/sign-up. A no-signup demo at https://lucydomi.com/demo returns the same data shape (limited to 3 calls per device per day).

A local stdio package is also published: `uvx busybee-mcp`. Client guides for Claude, Cursor, Cline, Windsurf, VS Code, Dify, n8n, Coze and plain HTTP are in the repository under `docs/`.

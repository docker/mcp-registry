Docs: https://www.mellow.world/hub/guides/social-media-mcp

Mellow Hub is a hosted MCP server: Streamable HTTP at `https://www.mellow.world/mcp`, OAuth 2.1 with dynamic client registration and PKCE. It finds the short videos taking off on TikTok, Instagram and YouTube, and publishes or schedules one post to ten networks: Instagram, TikTok, YouTube, X, LinkedIn, Threads, Bluesky, Pinterest, Facebook and Telegram. `validate_post` checks every network's rules before anything is sent, `create_post` requires an idempotency key, and Instagram auto-replies answer a keyword in a comment with a public reply and a direct message.

On the consent screen the owner chooses the channels a connection may reach, review (prepare only) or autopilot, a daily ceiling and an expiry. Every action is journaled, and revoking takes effect on the next request.

- Website: https://www.mellow.world
- Repository: https://github.com/wco1/mellow-hub-mcp
- Agent skill: https://github.com/wco1/mellow-hub-skill
- Official MCP Registry: `world.mellow.www/hub`

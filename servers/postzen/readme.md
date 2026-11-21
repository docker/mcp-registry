# PostZen MCP Server

Docs: https://docs.postzen.dev/mcp

PostZen is a social media publishing API. This remote MCP server exposes the PostZen API as tools, so an assistant can draft, schedule and publish posts, handle incoming comments and direct messages, and report on performance across ten platforms: X, Instagram, TikTok, LinkedIn, Facebook, YouTube, Threads, Pinterest, Bluesky and Telegram.

## Features

- **Publishing** - Create drafts, schedule posts, and publish immediately, with per-platform options such as threads, carousels, video and Pinterest boards.
- **Queues** - Define a posting schedule once, then drop posts into the next open slot.
- **Inbox** - Read and reply to comments, hide them, and handle Instagram direct messages.
- **Analytics** - Post timelines, daily metrics, follower stats and best-time-to-post suggestions.
- **Accounts** - Hosted connect flows for every platform, plus webhooks for publish and engagement events.

## Authentication

This entry uses a PostZen API key sent as a bearer token. Create one in the [PostZen dashboard](https://app.postzen.dev) under Settings, then paste it when enabling the server.

The same endpoint also supports OAuth 2.1 with dynamic client registration, which interactive clients can use instead of a key.

## More information

- Documentation: [docs.postzen.dev/mcp](https://docs.postzen.dev/mcp)
- Website: [www.postzen.dev](https://www.postzen.dev)

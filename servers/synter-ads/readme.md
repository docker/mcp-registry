# Synter Ads MCP Server

Cross-platform advertising MCP server that gives AI agents the ability to manage real ad campaigns across Google, Meta, LinkedIn, Microsoft, Reddit, TikTok, and more.

## Documentation

- **Setup Guide:** https://syntermedia.ai/mcp
- **API Documentation:** https://docs.syntermedia.ai
- **GitHub:** https://github.com/jshorwitz/synter-mcp-server
- **npm:** https://www.npmjs.com/package/@synterai/mcp-server

## Getting Started

1. Sign up at [syntermedia.ai](https://syntermedia.ai)
2. Get your API key at [syntermedia.ai/developer](https://syntermedia.ai/developer)
3. Connect your ad accounts at [syntermedia.ai/settings/credentials](https://syntermedia.ai/settings/credentials)

## Remote Connection

This server uses Streamable HTTP transport:

```
URL: https://mcp.syntermedia.ai/mcp/
Header: X-Synter-Key: syn_your_api_key_here
```

## Features

- **Campaign Management:** Create, pause, and update campaigns across 7+ platforms
- **Performance Analytics:** Pull impressions, clicks, spend, conversions, ROAS
- **AI Creative Generation:** Generate images (Imagen 4, Flux) and videos (Veo, Runway)
- **Conversion Tracking:** Set up and diagnose tracking across platforms
- **Keyword Management:** Add keywords and negative keywords
- **140+ Tools:** Full access via the `run_tool` command

# YouTube Autocomplete — FlashData

Get YouTube search suggestions for a query prefix, language and region. Build keyword idea sheets from actual autocomplete phrases.

Use this task-specific MCP entry with an AI agent. It exposes only `youtube_suggestions`, `get_pricing`, `get_job` and `get_result`, subject to the user's API Key permissions.

1. Create your own [FlashData API Key](https://flashdata.dev/app/api-keys?utm_source=docker-mcp-catalog&utm_medium=integration&utm_campaign=intent_gtm_2026q4) with `youtube_suggestions` access.
2. Save it in the Toolkit secret `flashdata-youtube-suggestions.api_key` (environment variable `FLASHDATA_API_KEY`).
3. Follow the [task guide with inputs, outputs and Python example](https://flashdata.dev/docs/sources/youtube_suggestions?utm_source=docker-mcp-catalog&utm_medium=integration&utm_campaign=intent_gtm_2026q4&utm_content=youtube_suggestions).

Installation is free. Data calls consume the user's own [FlashData Credits](https://flashdata.dev/pricing?utm_source=docker-mcp-catalog&utm_medium=integration&utm_campaign=intent_gtm_2026q4); Docker, client and storage costs are separate. Pricing/job/result reads do not submit another query. A timed-out query may already be charged; inspect Jobs before retrying.

[Container setup and limitations](https://github.com/rocosen/flashdata-mcp/blob/main/docker/README.md) · [Detailed example and result handling](https://github.com/rocosen/flashdata-mcp/blob/main/registry/guides/youtube_suggestions.md) · [Security](https://github.com/rocosen/flashdata-mcp/blob/main/SECURITY.md)

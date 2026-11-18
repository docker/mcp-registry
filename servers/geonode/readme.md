# Geonode MCP Server

Give MCP-compatible AI assistants web search and content collection tools. Extract pages, discover URLs, crawl sites and run batches for research and downstream analysis. Geonode offers limited keyless page extraction; an API key enables the full toolset, including JavaScript rendering and country targeting.

Supported functions:
• Extract webpage content from a compatible AI assistant or coding tool.
• Try basic synchronous page extraction without an API key, within anonymous limits.
• Use an API key for URL discovery, crawls, web search and batch jobs.
• Access JavaScript rendering and country targeting with authenticated requests.
• Monitor long-running extraction and crawl jobs from your MCP client.

Turn a page into context: Extract a page as Markdown for your assistant to read.

MCP server URL: https://mcp.geonode.com/mcp

Try limited keyless page extraction through a compatible MCP client. Sign up and connect your Geonode API key to access additional features and the usage limits included in your plan.

Try Geonode for free: https://app.geonode.com/sign-up

Learn more: https://geonode.com

## Connection

Endpoint: `https://mcp.geonode.com/mcp`
Transport: Streamable HTTP
Authentication: optional `X-Api-Key` header

Without an API key, only `extract` runs: synchronous page extraction, no JavaScript
rendering, and a small per-client rate and daily volume. An API key lifts these
restrictions and enables the remaining tools; the plan's own limits apply instead.

## Source

Client source and connection docs: https://github.com/geonodecom/mcp (MIT)

## Documentation

https://docs.geonode.com

## Licence

The MCP client integration in https://github.com/geonodecom/mcp is MIT licensed.
The hosted service behind `https://mcp.geonode.com/mcp` is operated by Geonode;
access and use are subject to Geonode's Terms of Service:
https://geonode.com/terms-and-conditions

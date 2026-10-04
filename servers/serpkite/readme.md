Docs: https://serpkite.com/docs/mcp

SerpKite hosts a stateless Streamable HTTP MCP server at `https://api.serpkite.com/v1/mcp`. Set the `serpkite.api_key` secret to your API key from https://app.serpkite.com; the gateway sends it in the `Authorization: Bearer` header. OAuth is not currently supported.

The server exposes fourteen tools: `search`, `news`, `maps`, `scholar`, `patents`, `shopping`, `images`, `videos`, `autocomplete`, `webpage`, `extract`, `map`, `crawl`, and `crawl_result`. Results are Markdown. Initialization and tool discovery do not run searches; tool calls use credits at the same rates as their REST endpoints. Credits never expire, and failed or empty searches are not billed. Use a dedicated API key with a monthly credit limit.

Connection examples and registry metadata: https://github.com/SerpKite/serpkite-mcp (MIT). The hosted API implementation is proprietary. The service is maintained by the SerpKite team. Support and security reports: support@serpkite.com.

Google is a trademark of Google LLC; SerpKite is not affiliated with Google.

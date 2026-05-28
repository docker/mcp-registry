# Parallel Search Remote MCP Server

The **Parallel Search MCP Server** exposes Parallel Web Systems' web search and URL extraction APIs through the Model Context Protocol. It provides two tools designed to slot directly into LLM agent loops:

- **`web_search`** — Pass an `objective` (natural-language description of the underlying goal) plus 2–3 short keyword `search_queries`; returns ranked URLs with LLM-optimized excerpts. One call replaces several keyword searches.
- **`web_fetch`** — Pass a list of URLs and an optional `objective`; returns the most relevant markdown content from each page, bounded for token efficiency.

Both tools return content already compressed and ranked for reasoning utility (rather than human engagement), so excerpts can be fed directly into the model context with minimal post-processing.

## Authentication

The server accepts unauthenticated requests by default for free, lower-rate-limit use — useful for exploration and light traffic.

For higher rate limits or production traffic, pass a Bearer token in the `Authorization` header using a key issued at [platform.parallel.ai](https://platform.parallel.ai):

```
Authorization: Bearer <PARALLEL_API_KEY>
```

When using the Docker MCP Toolkit, set the API key via the standard secret-injection mechanism for the `parallel-search` server.

## Docs and Links

- Search MCP overview: <https://docs.parallel.ai/integrations/mcp/search-mcp>
- Search API reference: <https://docs.parallel.ai/api-reference/search/search>
- Extract API reference: <https://docs.parallel.ai/api-reference/extract/extract>
- Best practices: <https://docs.parallel.ai/search/best-practices>
- Pricing: <https://docs.parallel.ai/getting-started/pricing>

## Example Use Cases

- Research-grade web lookups with inline citations (news, releases, market data).
- Drop-in replacement for multi-step keyword search loops — one `web_search` call with an `objective` plus 2–3 keyword queries typically suffices.
- URL summarization and extraction for any page the agent needs to read.
- Pipeline component for deep research / report generation when paired with Parallel's Task API.

# AgentSearch

stdio MCP server for web search, URL-to-markdown extract, JS page render, and one-payment `search_and_read`.

Default (no env): keyless free search (up to 25/day per IP) plus the extract/render trial.

Paid calls use AgentSearch's direct x402 API on Base at https://api.agentsearchhq.com when you set `AGENTSEARCH_WALLET_PRIVATE_KEY`. Prices: search $0.008, extract $0.007, render $0.010, search_and_read $0.015.

- Repo and setup: https://github.com/agentsearchhq-coder/agentsearch-mcp
- API docs: https://agentsearchhq.com/docs/x402
- npm: `@agentsearchhq/agentsearch-mcp` 1.7.0

Security reports: GitHub private vulnerability reporting on the repository above, or AgentSearchhq@gmail.com.

# aiworker-data

Remote MCP server (streamable HTTP) at `https://aiworker.duckdns.org/mcp`: 18 data tools for agents, each paid per call in USDC over the x402 v2 protocol. No API key and no account — a tool called without payment answers with its exact price and how to pay; a client with an x402 wallet (USDC on Base or Solana) pays per call, $0.005–$1.

Tools: DeFi yields and protocol snapshots (DefiLlama), Base token safety and wallet cards (Blockscout, Honeypot.is, DexScreener-derived), Polymarket odds, resolution, history, screener and rule backtests, crypto news search, Hacker News and Polymarket mentions, page-to-Markdown, fact checks and cited topic briefs.

- Catalogue with prices: https://aiworker.duckdns.org/llms.txt
- Client wrappers (LangChain, Vercel AI SDK, ElizaOS, OpenAI Agents SDK, Mastra): https://github.com/ai-worker227/aiworker-examples
- Health: https://aiworker.duckdns.org/health

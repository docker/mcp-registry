# FinData MCP

Financial data MCP server providing real-time market data via x402 micropayments.

## Tools

| Tool | Description | Cost |
|------|-------------|------|
| `stock_quote` | Real-time stock price, volume, market cap | $0.001/call |
| `company_fundamentals` | P/E, EPS, revenue, margins | $0.002/call |
| `economic_indicator` | GDP, CPI, unemployment, Fed rate | $0.001/call |
| `sec_filing` | Recent 10-K, 10-Q, 8-K filings | $0.002/call |
| `crypto_price` | Crypto price, 24h change, market cap | $0.001/call |

## Authentication

This server uses the [x402 payment protocol](https://x402.org) for pay-per-use access. No API key or subscription is required — each tool call is priced individually and settled via Base network USDC micropayments.

## Source

- Thin client: https://pypi.org/project/findata-mcp/
- GitHub: https://github.com/sapph1re/findata-mcp

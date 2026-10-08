# GMGN Market

Remote, anonymous, read-only MCP server with public on-chain market data from [GMGN](https://gmgn.ai).
No API key is required. It cannot trade, sign transactions or access user assets.

- Endpoint: `https://openapi.gmgn.ai/mcp/public` (Streamable HTTP)
- Chains: `sol`, `bsc`, `base`, `eth`, `arbitrum`, `hyperevm`, `robinhood`, `arc`, `stable`
- Rate limits: per-IP token bucket; rate-limited tool calls return `RATE_LIMITED` with `retry_after_ms`.

## Tools

| Tool                              | Description                                               |
| --------------------------------- | --------------------------------------------------------- |
| `system_get_capabilities`         | Describe the service and list its tools                   |
| `market_search_assets`            | Search tokens and wallets by name, symbol or address      |
| `market_get_token_overview`       | Token price, liquidity, metadata and risk overview        |
| `market_get_token_security`       | Contract and honeypot risk factors                        |
| `market_get_token_pool`           | Primary DEX pool and liquidity                            |
| `market_get_token_top_holders`    | Top holders of a token                                    |
| `market_get_token_top_traders`    | Top traders of a token (historical data)                  |
| `market_get_token_signals`        | Token signal observations                                 |
| `market_get_kline`                | OHLCV candles                                             |
| `market_get_rank`                 | Token ranking by market activity                          |
| `market_get_hot_searches`         | Hot-search ranking                                        |
| `market_get_trenches`             | New-token trenches data                                   |
| `market_get_launchpad_statistics` | Token-creation statistics by launchpad                    |
| `wallet_get_stats`                | Wallet performance statistics for a period                |
| `wallet_get_profits`              | Profit and loss for up to 20 wallets                      |
| `wallet_get_activity`             | Historical on-chain activity of a wallet                  |
| `wallet_get_created_tokens`       | Tokens created by a wallet                                |
| `wallet_list_smart_money`         | Public smart-money wallet activity                        |
| `wallet_list_kol`                 | Public KOL wallet activity                                |

All tools are annotated `readOnlyHint: true`, `destructiveHint: false`.
Data is for reference only and is not investment advice.

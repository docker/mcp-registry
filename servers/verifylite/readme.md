# VerifyLite

Verify before you pay or act: pay-per-call safety checks for sites, transactions and tokens.

## MCP endpoint

Remote server (Streamable HTTP). No Docker image is needed.

- Endpoint: https://verifylite.hydratinghealingtouch.workers.dev/mcp
- MCP Registry name: io.github.SecondShotAI/verifylite
- Homepage: https://verifylite.hydratinghealingtouch.workers.dev
- Public repo (MIT licensed): https://github.com/SecondShotAI/verifylite-mcp

## Tools

VerifyLite is pay-per-call in USDC via x402, on Base, Arc and Polygon. No accounts, no API keys, no login. Failed requests are never charged. `list_prices` is free.

| Tool | Price (USDC) |
| --- | --- |
| chain_token_report | $0.0098 |
| check_site_safety | $0.009 |
| tx_preflight | $0.0095 |
| check_prompt_injection | $0.0044 |
| ping | $0.001 |

More details, examples and [health status](https://github.com/SecondShotAI/verifylite-mcp/tree/main/status) are in the public repo.

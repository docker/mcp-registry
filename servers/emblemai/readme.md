# EmblemAI

EmblemAI is a non-custodial crypto wallet MCP server. The hosted endpoint at `https://emblemvault.ai/api/mcp` exposes 200+ on-chain tools across 7 blockchains (Solana, Ethereum, Base, BSC, Polygon, Hedera, Bitcoin) — wallet balance/transfer, token swaps, NFT operations, DeFi yield, Bitcoin Ordinals/Runes/BRC-20, Polymarket events, and market intelligence from Birdeye, CoinGlass, and Nansen.

**Non-custodial:** private keys live in the user's Emblem vault. The MCP server never sees or stores them.

## Tool surface

| Auth | Tools exposed | Capabilities |
|------|---------------|--------------|
| None (anonymous) | 133 | Read-only: balances, prices, market data, NFT browse, Ordinals research |
| `x-api-key` header | 200+ | All of the above, plus full read + write |
| `x-api-key` + `x-mcp-transactions: enabled` | 200+ | All of the above, plus transaction tools (transfer, swap, mint) |

State-changing tools always carry a `-transaction` suffix in `tools/list` once enabled, so clients can wildcard-filter them.

## Get an API key

Sign in at [emblemvault.ai](https://emblemvault.ai) and grab your Vault Access Key from **Settings → Vault Access Key**.

## Docs

- **MCP setup matrix per client (Claude Code, Cursor, Windsurf, Cline, GitHub Copilot, Claude Desktop):** https://emblemvault.ai/docs/mcp
- **Repository:** https://github.com/EmblemCompany/Agent-skills (MIT, 9 skills + MCP bridge, ships `llms-install.md`)
- **Live in other registries:** [Smithery](https://smithery.ai/server/@emblemai/emblem-mcp) (quality score 84), [Glama](https://glama.ai/mcp/servers/EmblemCompany/Agent-skills), [LobeHub](https://lobehub.com/mcp/emblemcompany-agent-skills), [Cline](https://github.com/cline/mcp-marketplace/issues/1604) (pending review)

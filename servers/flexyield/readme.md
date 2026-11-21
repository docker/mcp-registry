# FlexYield MCP Server

**FlexYield** is a multi-chain RPC gateway for humans and the agents they run: one API key, one meter, one bill across six mainnets (Ethereum, Base, Arbitrum One, OP Mainnet, Polygon PoS, Solana) with automatic multi-provider failover and a publicly measured status page.

The MCP server exposes the same capabilities as the HTTP API — 35 verb-first tools, one job each:

- **Relay** — `rpc_call(chain, method, params)` through the same routing, cache and meter as `POST /rpc/{chain}`
- **Toolbox** — gas recommendations, transaction status, ENS, on-chain prices, verified ABIs, calldata decoding, EVM utilities, signature verification
- **Keys and budgets** — create, configure, revoke keys; a monthly ceiling per key; service fences; statements per key
- **Agent menu, no key needed** — `get_key` (capability key without signup, 10,000 weighted units per day), `what_can_i_do`, `get_quote`, `explain_error`, `recommend`, `migrate`

Every limit answers with a machine-readable doorway that names the next step.

## Connect

- Endpoint: `https://flexyield.io/mcp` (Streamable HTTP)
- Auth: `Authorization: Bearer <api key>` — or call `get_key` first and use the returned key
- Agent skill file: https://flexyield.io/skill.md · Docs: https://flexyield.io/docs · Status: https://flexyield.io/status

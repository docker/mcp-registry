AssetFare is an agent-native, non-custodial cross-chain route service. Its
public product is a curated catalog of 65 whole routes: 3 core, 50 conditional,
and 12 strategic/scarce. Thirty-five additional researched routes are not
currently offered. The catalog is reviewed weekly and does not publish a
competitor-ranking directory or guarantee that an offered route is cheapest.

Use the live capability and route-selection tools to discover currently offered
endpoints. Request an exact fresh quote at the caller's intended amount; do not
hard-code chains, assets, providers, or route availability. Every quote exposes
the ordered `direct_route_summary`, exact 1bp AssetFare research-fee step,
complete priced/unpriced cost disclosure, and quote-bound `continuation_v3`.
Treat the result as an unranked candidate requiring a caller decision.

AssetFare never receives a private key, signs, or submits. Unsigned prepare and
session operations require explicit caller approval and public wallet addresses.
The caller independently verifies every action and signs/submits with its own
wallet. Never infer execution approval from a quote or call both one-shot and
session paths.

The primary hosted endpoint is the v2-only remote MCP at
`https://api.assetfare.dev/mcp`. Legacy compatibility tools are isolated at
`https://api.assetfare.dev/mcp/legacy`. Current package and hosted-adapter
release: `assetfare-mcp@1.21.5`.

Documentation: https://assetfare.dev/agents/

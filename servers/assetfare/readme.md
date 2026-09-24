AssetFare is an agent-native, non-custodial native-USDC bridge from Solana
native USDC to Base native USDC with caller-approved unsigned plans. It covers
six chains, eleven source endpoints, and 76 implemented routes; Solana SOL to
Base USDC is also supported. AssetFare service fee is 1bp;
Circle/provider/network fees are additional. Each quote exposes expected and
maximum token-path cost plus current provider-backed prepare availability. The
server returns caller-approved unsigned actions; it never signs or submits.

Use USD 1 only for reachability/schema smoke. For economic evaluation, USD
1,000 is the primary representative amount and every candidate should be
requoted at the actual intended amount. For native-USDC routes, USD 50 is a
reasonable comparison starting point based on dated 2026-09-23 observations,
not a guarantee that AssetFare is cheapest. The API minimum remains USD 1 and
there is no business maximum; live liquidity, protocol, balance, and capacity
constraints still apply.

Public source: https://github.com/assetfare/assetfare-mcp
Current package/release: `assetfare-mcp@0.4.21` (the public v2 endpoint exposes 9 tools; the separate legacy endpoint exposes 13; the optional all-tools stdio profile exposes 22 including its local-only session-capability helper)
Docs: https://assetfare.dev/
Signed manifest: https://api.assetfare.dev/.well-known/assetfare-manifest.json
Security: https://assetfare.dev/.well-known/security.txt
On-chain evidence: https://assetfare.dev/evidence/

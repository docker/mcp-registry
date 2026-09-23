AssetFare is an agent-native, non-custodial native-USDC bridge from Solana
native USDC to Base native USDC with caller-approved unsigned plans. It covers
six chains, eleven source endpoints, and 76 implemented routes; Solana SOL to
Base USDC is also supported. AssetFare service fee is 1bp;
Circle/provider/network fees are additional. Each quote exposes expected and
maximum token-path cost plus current provider-backed prepare availability. The
server returns caller-approved unsigned actions; it never signs or submits.

Public source: https://github.com/assetfare/assetfare-mcp
Current package/release: `assetfare-mcp@0.4.17` (21 remote tools; the optional stdio package has one additional local-only session-capability helper)
Docs: https://assetfare.dev/
Signed manifest: https://api.assetfare.dev/.well-known/assetfare-manifest.json
Security: https://assetfare.dev/.well-known/security.txt
On-chain evidence: https://assetfare.dev/evidence/

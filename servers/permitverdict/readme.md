# Permit Verdict

Documentation: https://permitverdict.com/developers

Settlement verdicts for crypto price events, and US building-permit intelligence, over a single
remote MCP endpoint.

**No account, no API key, no signup.** Five of the ten tools are free. The paid tools are
charged per call in USDC on Base using [x402](https://permitverdict.com/.well-known/x402), so an
agent pays for exactly what it uses and there is nothing to provision first.

| Tool | Price |
|---|---|
| `list_jurisdictions`, `check_permit_activity`, `find_active_projects`, `list_settlement_assets`, `decode_decision_id` | free |
| `get_crypto_spot` | $0.01 |
| `verify_decision` | $0.02 |
| `resolve_price_event` | $0.05 |
| `attest_price_event` | $0.25 |
| `get_permit_verdict` | $2.00 |

Settlement resolutions carry per-venue evidence, explicit touch-vs-close handling, and a stated
confidence. `attest_price_event` returns a `decision_id` signed by the wallet that receives
payment, so a counterparty can verify its origin offline.

Also published as an npm package (`permitverdict-mcp`) and in the official MCP registry as
`io.github.okwithit9-debug/permitverdict-mcp`.

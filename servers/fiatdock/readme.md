Docs: https://fiatdock.com/docs.html

FiatDock is a remote MCP server (streamable HTTP) at `https://fiatdock.com/mcp`: a marketplace
where AI agents find and buy MCP services published by other agents, paid per call in USDC over
the x402 protocol on Base — plus Base chain data and a cash-out of an agent's USDC to its owner's
bank account.

No account, no API key and no OAuth. Connect and call — five of the eighteen tools return a real
answer with no credentials and no wallet:

- `search_services` — search the marketplace catalog of pay-per-call services
- `get_service` — full detail for one listing, including the argument names its tools take
- `token_price` — price, liquidity and 24h volume for any EVM token
- `get_quote` — live rate and the exact net amount for a USDC cash-out, in any of 18 currencies
- `get_order_status` — status of an on/off-ramp order

The other thirteen are paid per call in USDC over x402 on Base. `call_service` buys a marketplace
listing at the seller's price, and settlement happens only after the seller's server answers: a
call that returns no answer costs nothing, while an answer the buyer merely dislikes is still a
delivered call. The data tools cost $0.001–$0.05 and minting a cash-out checkout costs $0.01. A
paid tool called without payment answers HTTP 402 with the exact price — a 402 is a price, not a
charge.

FiatDock is non-custodial: a seller's share settles on-chain directly to the seller's own wallet,
USDC for a cash-out moves between the agent's wallet and a licensed provider, and fiat moves between
the provider and the owner's bank. FiatDock never holds customer funds.

Machine-readable surfaces: https://fiatdock.com/llms.txt ·
https://fiatdock.com/openapi.json · https://fiatdock.com/.well-known/x402

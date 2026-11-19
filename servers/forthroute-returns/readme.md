Docs: https://app.forthroute.io/docs/ai-connector

# ForthRoute Returns

Remote MCP server for [ForthRoute](https://forthroute.io), a returns and return-label app for Shopify merchants (UK/EU carriers via Shippo, for example Hermes/Evri). It lets an AI agent search, review, approve and decline Shopify returns in ForthRoute.

- Endpoint: `https://app.forthroute.io/mcp` (streamable HTTP)
- Authentication: OAuth 2.1 with dynamic client registration, handled by the server. The merchant signs in with their ForthRoute account; there is no API key to configure.
- Publisher: Forthsuite (Hylke Reitsma), returns@forthroute.io
- Also published in the official MCP registry as `io.forthroute/returns`.

## Tools

- `get_returns_overview`: summary of the store's returns.
- `search_returns`: find returns by order number, RMA, customer name or email.
- `get_return`: full details of one return.
- `approve_return`: approve a return. Changes the return in Shopify and emails the customer.
- `decline_return`: decline a return. Changes the return in Shopify and emails the customer.

All amounts are in the store currency.

# Apiguru Amazon Data

Live Amazon marketplace data for AI agents: product details, reviews, keyword
search, best-sellers, deals, live offers and stock, and seller profiles across
20 country marketplaces. Twelve tools with prices and retry rules in their
descriptions.

Ten of them are read-only data fetches. `list_capabilities` answers offline
from the bundled spec and costs nothing. `send_feedback` is the only tool that
writes: it posts the message you give it to Apiguru's public feedback wall,
and it is free.

No API key is required. Each machine gets a small free daily budget; after
that the gateway answers with an HTTP 402 payment challenge (x402, USDC on
Base). Set `APIGURU_API_KEY` to bill an Apiguru account instead.

- Source and documentation: https://github.com/apiguru-app/agent-kit
- Tool reference: https://github.com/apiguru-app/agent-kit/blob/main/mcp/README.md
- API documentation: https://dash.apiguru.app/docs
- Support: support@apiguru.app

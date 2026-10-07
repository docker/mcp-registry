Docs: https://cubicle.algotrada.com/llms.txt

OpenAPI: https://cubicle.algotrada.com/deck/api/openapi.json

Keyless by default: `initialize`, `tools/list`, `catalogue`, `plans`, `strategies` and `strategy_quote` answer without a key. The other tools act as a Cubicle principal and need a free Cubicle key (https://cubicle.algotrada.com/deck/join, or `POST https://cubicle.algotrada.com/deck/api/join`) sent as the `Authorization: Bearer` header; every refusal says how to get one.

Ausca provides metered document OCR and analysis, media transcription, remote browser sessions with CDP access, and receive-only inboxes. Each service publishes its input contract, limits and price. Compatible clients pay per call in USDC on Base using x402 and receive evidence for completed work.

The hosted MCP exposes offer discovery, invocation preparation and state, and capability-scoped browser and inbox lifecycle tools. Public discovery needs no account or provider API key. Paid invocation requires explicit payment support on the caller's side. For a service without a dedicated paid MCP tool, use the canonical HTTP resource documented in its service skill.

The public client also provides a local MCP server (`npx -y ausca mcp`) for caller-side wallet signing under a configured spending cap. Follow the setup guide before using paid services.

Docs: https://ausca.com/docs
Service instructions: https://ausca.com/SKILL.md
OpenAPI: https://ausca.com/openapi.json
Public client source: https://github.com/auscahq/ausca

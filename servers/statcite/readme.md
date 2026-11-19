# StatCite

Official economic statistics with a full citation attached to every number.

**Endpoint:** `https://statcite.com/mcp` (streamable-http, stateless, no auth, no API key)

## What it does

Twelve tools over World Bank, IMF (WEO and Fiscal Monitor), BIS and ECB data. Every
response carries a citation object naming the source, dataset, series id, canonical URL,
licence, retrieval date and a ready-to-paste citation sentence.

The distinctive one is `verify_stat`: give it a claimed figure and it checks it against the
official series and returns a verdict (match / close / mismatch / cannot_verify) with
diagnostics for the usual failure modes, wrong year, percent-versus-decimal, and
millions-versus-billions. `verify_claims` batches up to 15 claims from a draft in one call.

Coverage is deliberate about its own gaps. Some economies are not World Bank reporting
economies at all, and for those the response says so explicitly with a machine-readable
`no_published_data` flag rather than returning a substitute figure from somewhere else.

- Docs: https://statcite.com/docs
- Sources and licences: https://statcite.com/sources
- Agent-readable reference: https://statcite.com/llms-full.txt
- OpenAPI (REST mirror): https://statcite.com/openapi.json

# CanYouGrab.it MCP Server

The **CanYouGrab.it MCP Server** lets AI agents check domain name availability with real-time DNS and WHOIS lookups — something LLMs can't do natively. Every result includes a confidence score (high/medium/low) so the agent knows how reliable the answer is, plus the data source (dns/whois/cache) and registration details for taken domains.

## Features

- **Confidence Scoring** – every result is rated high/medium/low so the agent can qualify its answer.
- **Bulk Checking** – up to 100 domains per request.
- **WHOIS/RDAP Enrichment** – registrar, creation/expiry dates, nameservers, and EPP status codes.
- **Read-Only and Safe** – all tools annotated `readOnlyHint: true`, `destructiveHint: false`.
- **Fast** – responses typically under 2 seconds; cached results under 100 ms.

## Tools

- `check_domains` – check availability of up to 100 domains per request, with confidence, source, cache age, and registrar details.
- `get_domain_info` – RDAP/WHOIS lookup for a registered domain (registrar, dates, nameservers, status).
- `check_usage` – view current plan, quota, and remaining lookups for the billing period.

## Authentication

OAuth 2.0 (Authorization Code with PKCE) via Auth0 at `login.canyougrab.it`. Free tier available at [portal.canyougrab.it](https://portal.canyougrab.it).

## More Information

- Website: https://canyougrab.it
- API docs: https://portal.canyougrab.it/docs
- Source: https://github.com/einiba/canyougrab-api
- Privacy policy: https://canyougrab.it/privacy

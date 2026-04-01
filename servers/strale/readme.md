Strale is trust and quality infrastructure for AI agents. 250+ verified data capabilities — IBAN validation, company registry lookups across 27 countries, sanctions and PEP screening, web scraping, lead enrichment, Web3 risk analysis — each continuously tested and assigned a dual-profile quality score. Every execution returns structured JSON with data provenance and an audit trail.

## What you can do

- **Validate financial identifiers** — IBAN, VAT, SWIFT/BIC, LEI, EORI
- **Look up companies** — registry data from 27 countries (SE, NO, DK, FI, UK, DE, FR, NL, US, AU, and more)
- **Screen for compliance risk** — sanctions lists, PEP databases, adverse media, beneficial ownership
- **Extract structured data** — invoices, PDFs, web pages, metadata
- **Assess domain reputation** — WHOIS, DNS, SSL, security headers, tech stack
- **Web3 intelligence** — wallet risk scoring, token safety, DeFi protocol data, ENS resolution

## Authentication

1. Sign up at [strale.dev/signup](https://strale.dev/signup) — free €2.00 trial credits, no card required
2. Set your API key as `STRALE_API_KEY`

## Free tier

These tools work without an API key:

- `strale_search` — browse all 250+ capabilities
- `strale_getting_started` — see free capabilities with examples
- `strale_trust_profile` — check quality scores
- `strale_execute` with free slugs: `email-validate`, `iban-validate`, `dns-lookup`, `url-to-markdown`, `json-repair`

## Quick start

1. Call `strale_search` with query "IBAN" to find validation capabilities
2. Call `strale_execute` with slug `iban-validate` and inputs `{"iban": "DE89370400440532013000"}`
3. Call `strale_trust_profile` with slug `iban-validate` to check the quality score

Docs: [strale.dev/docs](https://strale.dev/docs)

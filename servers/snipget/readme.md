# Snipget

Snipget is a hosted MCP server exposing 317 deterministic data utilities for AI agents: validation, normalization, parsing, entity matching, and PII redaction for names, addresses, phones, emails, dates, and identifiers, plus healthcare (NPI/DEA validation, NUCC taxonomy), chemistry, and biotech reference lookups. Every tool is a pure function (no LLM inside) with a consistent envelope and `_batch` variants; catalog discovery is built in.

- Documentation: https://snipget.ai/docs
- API reference: https://api.snipget.ai/redoc
- Authentication: OAuth 2.0 with dynamic client registration (handled by the MCP client; free tier available)

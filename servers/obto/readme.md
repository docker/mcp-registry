# OBTO

OBTO gives your AI assistant a hosted app platform: it scaffolds, deploys, and operates real full-stack apps — pages, API routes, server scripts, a database, scheduled jobs — in a multi-tenant workspace created automatically the first time you sign in.

- **Endpoint:** `https://app.obto.co/ms/mcp` (streamable HTTP)
- **Auth:** OAuth 2.0 with Google sign-in — a first sign-in auto-provisions your workspace (~30s); no API keys or configuration.
- **Reliability for agents:** pre-write validation, content-hash deploy receipts on every write, structured error envelopes, explicit per-call app/workspace scoping (stateless server).
- **Docs & skills:** https://github.com/obto-inc/platform (per-client connect instructions + portable SKILL.md build discipline).
- **Official MCP registry:** `co.obto/obto`.

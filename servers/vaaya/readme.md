# Vaaya

Vaaya ([vaaya.ai](https://vaaya.ai)) is the agent payment system: a remote MCP
server that gives any agent pay-per-call access to paid capabilities without
API keys.

- **Documentation:** https://vaaya.ai
- **Endpoint:** `https://vaaya.ai/mcp` (Streamable HTTP, OAuth 2.1 with RFC 9728
  discovery and RFC 7591 dynamic client registration; anonymous `tools/list`)
- **npm stdio shim (alternative):** [`@vaaya/mcp`](https://www.npmjs.com/package/@vaaya/mcp)
- **Revoke access:** https://vaaya.ai/connected-apps
- **Support:** support@vaaya.ai

## What agents can do through Vaaya

Image & video generation (Nano Banana Pro 2, GPT Image 2, Seedream, Kling,
Seedance), narrated product demo videos from a raw screen capture, web search
(Exa, Parallel), web scraping (Firecrawl), deep/market/competitive research,
GTM & sales (lead finding, contact enrichment, buying signals, LinkedIn and
email outreach with human approval), scheduled monitoring workers, code
sandboxes (Modal), browser automation (Browserbase), agent email (AgentMail),
and persistent memory (mem0/Zep/Letta).

The `consult` tool is the entry point: describe any goal in plain English and
it returns the exact call(s) to run via `use`. Billing is on success only —
failed calls are never charged, and every call carries a `max_cost_cents`
ceiling.

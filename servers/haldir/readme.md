# Haldir

The guardian layer for AI agents — scoped sessions with spend caps (Gate), AES-256-GCM encrypted secrets (Vault), hash-chained tamper-evident audit trail (Watch), and policy enforcement proxy for every tool call (Proxy).

MIT licensed. Use the hosted service (this server) or [self-host with `docker compose up`](https://github.com/ExposureGuard/haldir/blob/main/SELF_HOSTING.md).

- **Documentation:** [haldir.xyz/docs](https://haldir.xyz/docs)
- **OpenAPI spec:** [haldir.xyz/openapi.json](https://haldir.xyz/openapi.json)
- **LLM-readable docs:** [haldir.xyz/llms-full.txt](https://haldir.xyz/llms-full.txt)
- **Source:** [github.com/ExposureGuard/haldir](https://github.com/ExposureGuard/haldir)
- **Security contact:** [haldir.xyz/.well-known/security.txt](https://haldir.xyz/.well-known/security.txt)

## Getting an API key

Free, one click, no signup required:

```bash
curl -X POST https://haldir.xyz/v1/demo/key
```

Or visit [haldir.xyz/quickstart](https://haldir.xyz/quickstart).

## Tools

Haldir exposes these MCP tools:

- `createSession` — create a scoped agent session with spend cap and TTL
- `getSession` — look up a session's current state
- `revokeSession` — kill switch; invalidates the session immediately
- `checkPermission` — verify a session has a required scope before acting
- `storeSecret` — put a credential in the AES-256-GCM vault
- `getSecret` — retrieve a secret, scope-checked
- `authorizePayment` — deduct against the session's spend cap
- `logAction` — record an action to the hash-chained audit trail
- `getAuditTrail` — query the audit log
- `getSpend` — running total of what the agent has spent

## Native framework integrations

- `langchain-haldir` — [PyPI](https://pypi.org/project/langchain-haldir/) · `HaldirCallbackHandler` + `GovernedTool`
- `crewai-haldir` — [PyPI](https://pypi.org/project/crewai-haldir/) · `GovernedTool.wrap()`
- `@haldir/ai-sdk` — [npm](https://www.npmjs.com/package/@haldir/ai-sdk) · `governTool()` for Vercel AI SDK

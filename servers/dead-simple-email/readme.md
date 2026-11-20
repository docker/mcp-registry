# Dead Simple Email

Email for AI agents. Documentation: https://deadsimple.email/docs.html

Authentication is a bearer token sent per request:

```
Authorization: Bearer dse_your_api_key
```

A key can be minted in one unauthenticated call — no signup form, no dashboard:

```bash
curl -X POST https://api.deadsimple.email/v1/auth/agent-signup
```

That returns a trial account (1 inbox, 10 sends/hour, 25/day) plus a live inbox.

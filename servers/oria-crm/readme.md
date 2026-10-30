# Oria CRM

https://realoria.com/crm/mcp

Read-only MCP server for [Oria CRM](https://realoria.com/crm), a real-estate
CRM used by agencies in Romania. Exposes 17 read tools (CRM overview,
properties, contacts, deals, viewings, agents, open-house auctions, gallery,
publish checklist, property/contact matching) plus `search` and `fetch` for
ChatGPT's deep-research contract. No write tools are registered on this
endpoint.

## Authentication

This server implements OAuth 2.1 with dynamic client registration (RFC 7591)
and protected-resource metadata (RFC 9728) — the same mechanism Claude and
ChatGPT custom connectors use for "Sign in" flows. It also accepts a static
personal access token via `Authorization: Bearer <token>` for clients that
don't do interactive sign-in (e.g. Cursor, Claude Code). Tokens/grants are
created and revoked by an agency admin from inside Oria CRM
(`/admin/integrations/mcp`); nothing is provisioned from this catalog entry.

Every call is scoped to the signed-in agency's own data — one organization
per token/grant, enforced in the API layer, never by client-supplied input.

## Note for reviewers

Oria CRM's OAuth server is a from-scratch OAuth 2.1 authorization server
(not a third-party provider), discoverable via
`/.well-known/oauth-protected-resource/api/mcp` and
`/.well-known/oauth-authorization-server` — both respond without a token. If
the gateway's `oauth:` config block expects a single fixed provider/PAT
shape rather than dynamic client registration, the personal-access-token
path above (`Authorization: Bearer`) is the fallback that needs no special
gateway support.

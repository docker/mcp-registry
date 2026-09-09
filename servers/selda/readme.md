# Selda MCP Server

Selda is a go-to-market engine. It finds the companies worth reaching, researches each one, and
drafts the outreach — for anyone who sells to other businesses and has nobody whose job is selling.

The research happens when you ask. Selda reads the prospect's own site and what is currently
published about them, and the reason to write is found then rather than looked up in a database.
That is the difference from a contact list: a database can say a company exists and who works
there, but not that they opened a second location last week.

**Nothing here sends anything.** The one function that sends is in no tool registry and takes no
parameter reachable from this server. Drafts wait for a human in the Selda app. That is a property
of the registry, not a setting.

## Connecting

Endpoint: `https://mcp.selda.ai/api/mcp` (Streamable HTTP)

Authentication is either OAuth 2.1 with dynamic client registration, discovered from
`/.well-known/oauth-protected-resource` so no key is pasted anywhere, or a bearer API key from
**app.selda.ai → Settings → Apps → API Key**.

## What the tools do

Projects and leads, campaigns and runs, drafts and threads, the message structure, the workspace's
knowledge base, your own uploaded material, and inbound events. Plus the `search` and `fetch` pair
that the ChatGPT connector contract expects.

Two habits make the difference: list projects first, since almost everything takes a `projectId`;
and use `selda_add_lead` for companies you already have, `selda_run_pipeline` only for open-ended
discovery like "find SaaS founders in Finland".

## More information

Full tool reference: https://docs.selda.ai/reference/mcp
Getting started: https://docs.selda.ai/ways-to-use/mcp-server
REST API and OpenAPI 3.1: https://docs.selda.ai/reference/rest

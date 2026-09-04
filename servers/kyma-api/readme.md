# Kyma API

Kyma API is an LLM gateway: one OpenAI-compatible endpoint (plus an Anthropic-compatible one) in front of open and frontier models, with uptime measured per model from real traffic and scheduled probes. This MCP server gives your agent the live catalog, prices, rankings, uptime and your own spend as read tools, and one chat tool guarded by a per-connection spend cap.

## What you get

- Hosted, remote, OAuth 2.1 sign-in with PKCE and dynamic client registration. No API key is pasted into the client.
- A dedicated key per connection, separate from your REST API keys, capped at $10 per 30 days by default (adjustable from $1 to $500 at https://kymaapi.com/integrations) and revocable at any time.
- 15 tools. 13 appear on a default connection; `list_keys` and `set_low_balance_alert` appear when their optional scopes are granted. Read tools never charge. `send_message` is the only tool that spends credit; the client must allow it first, and every reply reports the call's cost, the amount spent and the cap.
- Measured uptime per model via `get_model_uptime`, published at https://kymaapi.com/status. If a route fails mid-request, Kyma retries the same model on another route.

## Tools

| Tool | What it does | Approval |
|---|---|---|
| `list_models`, `get_model`, `list_pricing` | Live catalog, one model's details, prices | auto, read-only |
| `list_rankings`, `get_model_uptime`, `recommend_model` | Rankings from real traffic, 30-day uptime, best model for your agent | auto, read-only |
| `get_credits`, `get_spend`, `get_transactions`, `get_topup_link` | Balance, this connection's cap and remaining budget, ledger, billing page link | auto, read-only |
| `search_docs`, `ping` | Docs search, health check | auto, read-only |
| `list_keys` (optional scope `keys.read`) | Your REST keys by name, masked | auto, read-only |
| `set_low_balance_alert` (optional scope `billing.alerts`) | Balance at which Kyma emails you | asks once |
| `send_message` | A chat completion through any model; the only tool that spends | asks once; no charge without your Allow |

## Security and data

The connection cannot create or delete API keys, change billing, buy credits or register accounts; those scopes do not exist on this server. Access tokens live 7 days, refresh tokens 90 days and rotate on use; connection keys are stored hashed. Request logs are kept 90 days for billing reconciliation and abuse handling, then deleted; Kyma does not train on customer data. Privacy: https://kymaapi.com/privacy

## Links

- Server: https://mcp.kymaapi.com/mcp
- Docs: https://docs.kymaapi.com/guides/mcp-server
- Source and manifests: https://github.com/kyma-api/kyma-mcp-plugin
- Support: hello@kymaapi.com

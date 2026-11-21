# ParlayAPI

Use sports odds, player props, public discovery and account tools from a local
MCP client. This stdio server exposes the existing ParlayAPI MCP tool set; it
does not host an HTTP endpoint or a shared odds feed.

## Configure your own account

Public discovery tools work without an API key. For account data, add your own
ParlayAPI key in the catalog's secret field. The secret is supplied as
`PARLAYAPI_KEY` to your local server. Each user should configure their own
account, and account allowances apply to tool calls. Do not share one key across
a community.

[Get an account](https://parlay-api.com/signup?utm_source=docker_mcp&utm_medium=marketplace&utm_campaign=catalog)
and review [current plans and request limits](https://parlay-api.com/pricing).
Availability depends on the requested sport, event, market and source. An odds
quote, timestamp or calculation does not guarantee freshness or profit.

## Actions and permissions

The 22 tools include public pricing/discovery, account usage, sports odds,
player props and calculation endpoints. Four tools can change state or send
something: `parlayapi_signup` creates an account, `parlayapi_checkout_link`
creates a checkout link, `parlayapi_magic_link` sends login email, and
`parlayapi_set_bettable_books` updates account preferences. Approve these actions
explicitly in your client. Starting the process or listing tools makes no API
request and performs none of those actions.

The container uses stdio, a non-root user and no host volumes or published ports.
It contacts ParlayAPI when a tool is called. Keep keys and account responses
private within your own client; this is not a multi-user relay.

## Documentation and data use

- [Source and client configuration](https://github.com/JacobiusMakes/parlay-api-mcp)
- [API documentation](https://parlay-api.com/docs)
- [MIT software license](https://github.com/JacobiusMakes/parlay-api-mcp/blob/main/LICENSE)
- [Applicable API Terms](https://parlay-api.com/terms)

Sharing the MIT code or container grants no API data distribution rights. Terms
and any written agreement govern data use; these local defaults do not amend
existing agreements. Report security concerns privately to support@parlay-api.com
and omit keys from public issues.

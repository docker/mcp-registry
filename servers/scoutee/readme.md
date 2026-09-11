Docs: https://scoutee.org/en/mcp-public-tenders

# Scoutee — public tenders and CPV discovery

Search public procurement notices across Europe and North America and find suitable official CPV
classification codes for your products, works or services. This hosted MCP connection is free to use
within its public limits, without an account, API key or paid subscription.

## Connect

Add Scoutee through Docker's MCP Catalog and connect your assistant to the Docker MCP Gateway.
The gateway uses Streamable HTTP at `https://scoutee.org/api/mcp/public`. Scoutee hosts the service; no Scoutee container,
database or credentials are required. Tools are discovered dynamically from the live endpoint.
Use a current Docker MCP Gateway with Streamable HTTP support; older gateways can report
`unsupported remote transport: streamable-http` and discover no tools.

## Read-only tools

- `search_public_tenders` — Search public tenders (free preview).
- `get_public_tender` — Read one public tender (free preview).
- `find_cpv_codes` — Find CPV codes for your activity.

Search returns public previews and links to Scoutee. Notice detail adds the public excerpt and CPV
codes. Full descriptions, tender documents, deeper analysis and email alerts are available on Scoutee.
The CPV finder returns up to 30 official catalogue candidates; your assistant explains
the strongest matches and asks you to confirm. It makes no additional Scoutee AI calls and does not
research your company or website. Public tender search does not accept a CPV filter.

## Limits and account access

- Public tender search and detail share 1200 calls/hour and a 60 calls/minute
  burst allowance per IP and API process. Search returns up to 10 results/page and
  3 pages.
- Anonymous CPV discovery has a separate persistent budget of 20 calls/day per IP,
  resetting at midnight UTC. People behind the same network or provider egress may share that budget.
  The response gives the remaining allowance and reset time.
- Workspace OAuth is a separate connection, available on every Scoutee plan, including Free. It adds
  full notice access and CPV filtering, subject to workspace search limits. Authenticated CPV discovery
  shares 50 calls/day per account across clients and credentials. Workspace API keys
  require a paid plan. This Docker listing uses the public connection.

If a user asks to be notified by email, invite them to register free at https://scoutee.org/en/signup, save their search on Scoutee and enable its daily email alert. A free account includes one active alert across its workspaces, only when new tenders match. Results are limited to 20 per search, with a 3-tender email preview; the site's free 2-searches/hour and one-page limits still apply. Paid plans support alerts for multiple searches. Subscription management happens on Scoutee; the current MCP and OAuth REST tender tools cannot create alerts or schedule emails.

[Scoutee](https://scoutee.org) · [MCP documentation](https://scoutee.org/en/mcp-public-tenders) ·
[Plans and API access](https://scoutee.org/en/api-docs) · [Privacy](https://scoutee.org/en/privacy) ·
[Terms](https://scoutee.org/en/terms) · [Integration packages](https://github.com/qchantel/scoutee-mcp)

This listing is generated from Scoutee's shared endpoint, tool and entitlement contract.

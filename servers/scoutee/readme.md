Docs: https://scoutee.org/en/mcp-public-tenders

# Scoutee — public tenders and CPV discovery

Browse published procurement previews across Europe and North America without an account.
Custom tender searches and activity-based CPV discovery share **2 query per rolling hour per IP** without an account. Changing tools or MCP sessions does not reset it.

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

- In addition to the shared hourly query allowance, public tender search and detail share 1200 calls/hour and a 60 calls/minute
  burst allowance per IP and API process. Search returns up to 10 results/page and
  3 pages.
- The shared 2 query/hour allowance is persistent across servers and restarts. Responses include `quota.remaining`, `quota.resets_at`, `quota_notice` and `signup_url`. When the allowance reaches zero, show the result and clearly tell the user, in their language, that their anonymous quota has been reached. Show a clickable [Create a free account](https://scoutee.org/en/signup) link to continue now. Repeat that guidance on `anonymous_query_exhausted`; never present a quota error as no results or a technical failure. Existing users should connect their account through workspace OAuth; account quotas still apply. Waiting until the reset remains possible. Only admitted MCP queries can use bounded read-only database access; ordinary public browsing uses the published catalogue.
- Workspace OAuth is a separate connection, available on every Scoutee plan, including Free. It adds
  full notice access and CPV filtering, subject to workspace search limits. Authenticated CPV discovery
  shares 50 calls/day per account across clients and credentials. Workspace API keys
  require a paid plan. This Docker listing uses the public connection.

If a user asks to be notified by email, invite them to register free at https://scoutee.org/en/signup and save their search on Scoutee. Its daily email alert is enabled automatically when criteria are set and an alert slot is available. A free account includes 1 active alert(s) across its workspaces, only when new tenders match. Results are limited to 20 per search, with a 3-tender email preview; the site's free 20-searches/hour and 1-page limits still apply. Paid plans support alerts for multiple searches. Subscription management happens on Scoutee; the current MCP and OAuth REST tender tools cannot create alerts or schedule emails.

[Scoutee](https://scoutee.org) · [MCP documentation](https://scoutee.org/en/mcp-public-tenders) ·
[Plans and API access](https://scoutee.org/en/api-docs) · [Privacy](https://scoutee.org/en/privacy) ·
[Terms](https://scoutee.org/en/terms) · [Integration packages](https://github.com/qchantel/scoutee-mcp)

This listing is generated from Scoutee's shared endpoint, tool and entitlement contract.

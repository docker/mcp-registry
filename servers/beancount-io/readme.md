# Beancount.io

Connect to the hosted Beancount.io service over Streamable HTTP with OAuth.

After enabling the remote entry in your Docker MCP catalog/profile, request the
ledger permissions needed for queries and previews:

```sh
docker mcp oauth authorize beancount-io --scopes 'ledger.read ledger.write'
```

Select the intended ledger on Beancount.io's consent page. A preview uses
`appendLedgerText` with `dry_run: true`; it does not commit accounting changes.

- [Connection, authorization and tool guide](https://github.com/bex-co/beancount-io/blob/main/backend-cluster/backend-v2/docs/mcp.md)
- [Docker Gateway setup and dated native verification](https://github.com/bex-co/beancount-io/blob/main/backend-cluster/backend-v2/docs/marketplace-submissions/docker/README.md)

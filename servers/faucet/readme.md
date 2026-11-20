# Faucet

Faucet is an open-source (MIT) server that turns PostgreSQL, MySQL, MariaDB, SQL Server, Oracle, SQLite or Snowflake into a REST API and an MCP server, with an admin UI, API keys and role-based access control per service, table and verb.

This catalog entry runs `faucet mcp` (stdio) in the `faucetdb/faucet` image with the `faucet-data` volume mounted at `/data`. It serves the databases registered in that volume. Stdio mode runs with local admin rights, so role-based access rules are not applied to these tool calls.

## Setup

Register your databases in the same volume first, either through the admin UI:

```
docker run -p 8080:8080 -v faucet-data:/data faucetdb/faucet
```

then open http://localhost:8080 and add a database, or from the CLI:

```
docker run --rm -v faucet-data:/data faucetdb/faucet db add --name mydb --driver postgres --dsn "postgres://user:pass@host.docker.internal:5432/mydb"
```

The database host must be reachable from inside a container (for a database on your machine, use `host.docker.internal`).

If you want role-based access per API key, use the streamable HTTP endpoint at `http://localhost:8080/mcp` of a running Faucet server with an `X-API-Key` header instead of this stdio entry.

## Tools

- `faucet_list_services`, `faucet_list_tables`, `faucet_describe_table`
- `faucet_query`, `faucet_insert`, `faucet_update`, `faucet_delete`
- `faucet_raw_sql` (only on services with raw SQL enabled)

Docs: https://wiki.faucetdb.ai/mcp-server

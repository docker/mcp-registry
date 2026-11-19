# AlgoRover Research

Hosted Streamable HTTP MCP endpoint: https://research.algorover.ai/mcp

AlgoRover Research discovers published strategies and runs historical backtests on public market data. It does not connect exchange accounts or execute orders. Public discovery is anonymous; personal tasks and reports require OAuth with `research:read research:run`. Never share tokens or passwords in chat.

## Docker Gateway validation

Validated with Docker MCP Gateway 0.44.2 on macOS. Generate the catalog with `go run ./cmd/catalog algorover-research`, then copy `catalogs/algorover-research/catalog.yaml` into `~/.docker/mcp/catalogs/algorover-research.yaml`. Create a local catalog and profile using the resulting absolute file URL:

```sh
docker mcp catalog create algorover/research:local --server file:///ABSOLUTE/PATH/.docker/mcp/catalogs/algorover-research.yaml
docker mcp profile create --id algorover-research --name 'AlgoRover Research' --server file:///ABSOLUTE/PATH/.docker/mcp/catalogs/algorover-research.yaml
DOCKER_MCP_USE_CE=true docker mcp oauth authorize algorover-research --scopes 'research:read research:run'
DOCKER_MCP_USE_CE=true docker mcp gateway run --profile algorover-research
```

Complete Research browser sign-in and consent. Both authorization and gateway commands must use the same OAuth mode. In a client configuration, set `DOCKER_MCP_USE_CE` to `true` in the gateway process environment.

Real gateway calls passed: public `search_strategies`, authenticated `get_profile`, `list_backtests`, and `get_research_usage`. The server exposes 16 Research tools; the profile also exposes Docker's management tools. New backtest execution and cancellation were not tested in this Docker validation.

## Known Docker compatibility limitations

Docker Desktop's default OAuth URL omitted the RFC 8707 `resource` parameter and was rejected by Research. The official standalone OAuth mode above includes the parameter during authorization and token exchange.

Gateway 0.44.2 reported a native-keychain legacy-credential cleanup error after issuing and storing real access/refresh tokens. A subsequent gateway connection successfully used the stored credentials for the authenticated calls listed above. Treat the CLI error as a limitation; verify authenticated access rather than assuming authorization succeeded. Standard Docker Desktop UI OAuth has not passed validation. Maintainer review is requested before general catalog publication.

## Usage and support

Search strategies, inspect an exact returned version, and prepare a supported historical backtest. Review the actual preview before starting. Personal tasks and reports belong to the signed-in account. Historical results do not establish future returns.

- Publisher: Rivendell Holdings Investment Limited
- Website: https://www.algorover.ai/
- Support: https://www.algorover.ai/en/help/
- Privacy: https://www.algorover.ai/privacy/
- Terms: https://www.algorover.ai/terms/

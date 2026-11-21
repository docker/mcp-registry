# SigNoz MCP Server

Query metrics, traces, logs, alerts, dashboards, and services from your [SigNoz](https://signoz.io) observability platform using natural language.

This server runs as a local Docker container (stdio) and connects to your SigNoz instance — SigNoz Cloud (any region) or self-hosted — using an API key.

Docs: https://signoz.io/docs/ai/signoz-mcp-server/

## Prerequisite: create an API key

In your SigNoz UI, go to **Settings → API Keys → New Key** (requires an **Admin** user) and copy the key. This is your `SIGNOZ_API_KEY`.

## Configuration

| Field            | Description                                                                                                                                            |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `signoz_url`     | The URL of your **SigNoz instance** (not an MCP endpoint). Cloud: `https://your-org.us.signoz.cloud`. Self-hosted: `http://host.docker.internal:8080`. |
| `SIGNOZ_API_KEY` | The API key created above (stored as a secret).                                                                                                        |

> `signoz_url` is your SigNoz app/backend URL. Do **not** enter a hosted-MCP endpoint such as `https://mcp.us.signoz.cloud/mcp` here.

## Install: SigNoz Cloud

1. Docker Desktop → **MCP Toolkit → Catalog → SigNoz → Configure**.
2. Set `signoz_url` to your Cloud instance URL, e.g. `https://your-org.us.signoz.cloud` (use your own region: `us`, `eu`, `in`, `us2`, `eu2`, `in2`).
3. Set `SIGNOZ_API_KEY`.
4. Enable the server, then connect your MCP client (Claude Desktop, Cursor, VS Code, …).

CLI equivalent:

```bash
docker mcp server enable signoz-mcp-server
docker mcp config write signoz-mcp-server signoz_url "https://your-org.us.signoz.cloud"
docker mcp secret set signoz-mcp-server.signoz_api_key "<your-api-key>"
```

## Install: self-hosted SigNoz

Same flow, pointing at your own instance:

1. Docker Desktop → **MCP Toolkit → Catalog → SigNoz → Configure**.
2. Set `signoz_url` to your self-hosted URL. Because the MCP server runs inside a container, use `http://host.docker.internal:8080` (not `http://localhost:8080`) to reach SigNoz running on your host.
3. Set `SIGNOZ_API_KEY` from your self-hosted SigNoz's **Settings → API Keys**.
4. Enable the server and connect your MCP client.

CLI equivalent:

```bash
docker mcp server enable signoz-mcp-server
docker mcp config write signoz-mcp-server signoz_url "http://host.docker.internal:8080"
docker mcp secret set signoz-mcp-server.signoz_api_key "<your-api-key>"
```

## Verify

Ask your client something like _"list my SigNoz services"_ — you should see your services with call rates, error rates, and latencies.

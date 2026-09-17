# Datris

The data control plane for AI agents. Agents ask Datris for data; Datris acquires it, validates it, lands it in the stores you already run, and returns it with provenance, without the agent ever holding a credential.

- 74 MCP tools: pipelines, AI-generated taps, plain-English data quality rules, SQL and MongoDB queries, vector search, metadata discovery, secrets, job status and lineage
- Destinations: PostgreSQL, MongoDB, Snowflake, Databricks, Kafka, S3/MinIO, Qdrant, Weaviate, Milvus, Chroma, pgvector
- Credentials are brokered by HashiCorp Vault; agents reference secrets by name
- Self-hosted, open source

## Setup

1. Install the platform: `curl -fsSL https://get.datris.ai/install.sh | sh` (Docker only, no checkout needed).
2. Set `api_url` to your Datris server, for example `http://host.docker.internal:8080` for a local install.
3. Leave the API key blank unless your platform runs with `USE_API_KEYS=true`.

## Links

- Documentation: https://docs.datris.ai/mcp-server
- Repository: https://github.com/datris/datris-platform-oss
- PyPI: https://pypi.org/project/datris-mcp-server/
- Official MCP Registry: https://registry.modelcontextprotocol.io/servers/io.github.datris/datris

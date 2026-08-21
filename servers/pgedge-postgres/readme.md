# pgEdge Postgres MCP Server

Enterprise-grade PostgreSQL integration for AI applications. Query databases with natural language, perform hybrid search (pgvector + BM25), analyze query performance, and explore schemas — all through the Model Context Protocol.

Works with PostgreSQL 14+ and all major cloud providers (AWS RDS, Azure, GCP Cloud SQL, Supabase, Neon, etc.).

## Available Tools

| Tool | Description |
|------|-------------|
| **query_database** | Execute SQL queries with natural language translation. Read-only by default. |
| **get_schema_info** | Discover tables, columns, types, constraints, indexes, and foreign keys. |
| **execute_explain** | Run EXPLAIN ANALYZE to diagnose query performance and suggest optimizations. |
| **count_rows** | Efficiently count rows with optional filtering. |
| **similarity_search** | Hybrid search combining pgvector semantic similarity, BM25 lexical matching, and MMR diversity. |
| **generate_embedding** | Generate vector embeddings via OpenAI, Voyage AI, or Ollama. |
| **search_knowledgebase** | Search a pre-built documentation knowledgebase for PostgreSQL and pgEdge topics. |

## Quick Start with Sample Database

To try the MCP server with a sample Northwind database:

**1. Start a PostgreSQL container with Northwind data:**

```bash
docker run -d --name pgedge-sample-db \
  -e POSTGRES_USER=demo \
  -e POSTGRES_PASSWORD=demo123 \
  -e POSTGRES_DB=northwind \
  -p 5432:5432 \
  ghcr.io/pgedge/pgedge-postgres:17-spock5-standard \
  -c listen_addresses='*'
```

Then load the Northwind dataset:

```bash
docker exec pgedge-sample-db bash -c \
  "curl -fsSL https://downloads.pgedge.com/platform/examples/northwind/northwind.sql | psql -U demo -d northwind"
```

**2. Configure the MCP server in Docker Desktop MCP Toolkit:**

| Setting | Value |
|---------|-------|
| `pgedge_db_host` | `host.docker.internal` |
| `pgedge_db_port` | `5432` |
| `pgedge_db_name` | `northwind` |
| `pgedge_db_user` | `demo` |
| `pgedge_db_password` (secret) | `demo123` |

**3. Connect a client** (Claude, Cursor, etc.) via `docker mcp client connect <client-name>` and start querying:

- "What tables are in this database?"
- "Show me the top 10 customers by order count"
- "Explain the performance of SELECT * FROM orders WHERE ship_country = 'France'"

## Configuration Reference

### Required

| Variable | Description |
|----------|-------------|
| `PGEDGE_DB_HOST` | PostgreSQL hostname. Use `host.docker.internal` for databases on your host machine. |
| `PGEDGE_DB_NAME` | Database name. |
| `PGEDGE_DB_USER` | Database username. |
| `PGEDGE_DB_PASSWORD` | Database password (stored as a secret). |

### Database Options

| Variable | Default | Description |
|----------|---------|-------------|
| `PGEDGE_DB_PORT` | `5432` | PostgreSQL port. |
| `PGEDGE_DB_SSLMODE` | `prefer` | SSL mode: `disable`, `prefer`, `require`, `verify-ca`, `verify-full`. |
| `PGEDGE_DB_ALLOW_WRITES` | `false` | Allow INSERT, UPDATE, DELETE queries. |

### Embedding & Semantic Search

Enable these to use the `similarity_search` and `generate_embedding` tools.

| Variable | Default | Description |
|----------|---------|-------------|
| `PGEDGE_EMBEDDING_ENABLED` | `false` | Enable embedding generation. |
| `PGEDGE_EMBEDDING_PROVIDER` | `ollama` | Provider: `ollama`, `openai`, or `voyage`. |
| `PGEDGE_EMBEDDING_MODEL` | `nomic-embed-text` | Model name (depends on provider). |
| `PGEDGE_OLLAMA_URL` | `http://localhost:11434` | Ollama server URL. Use `http://host.docker.internal:11434` in Docker. |
| `OPENAI_API_KEY` | — | OpenAI API key (secret). Required if provider is `openai`. |
| `VOYAGE_API_KEY` | — | Voyage AI API key (secret). Required if provider is `voyage`. |

### LLM Proxy

Enable a built-in LLM chat proxy for the pgEdge web client interface.

| Variable | Default | Description |
|----------|---------|-------------|
| `PGEDGE_LLM_ENABLED` | `false` | Enable LLM proxy. |
| `PGEDGE_LLM_PROVIDER` | `anthropic` | Provider: `anthropic`, `openai`, or `ollama`. |
| `PGEDGE_LLM_MODEL` | `claude-sonnet-4-5-20250514` | Model name. |
| `PGEDGE_LLM_MAX_TOKENS` | `4096` | Maximum response tokens. |
| `PGEDGE_LLM_TEMPERATURE` | `0.7` | Sampling temperature (0.0–1.0). |
| `ANTHROPIC_API_KEY` | — | Anthropic API key (secret). Required if provider is `anthropic`. |

### Knowledgebase Search

Enable the `search_knowledgebase` tool for PostgreSQL and pgEdge documentation search.

| Variable | Default | Description |
|----------|---------|-------------|
| `PGEDGE_KB_ENABLED` | `false` | Enable knowledgebase search. |

### HTTP Mode

By default the server runs in MCP stdio mode. Enable HTTP mode for API access and the web client.

| Variable | Default | Description |
|----------|---------|-------------|
| `PGEDGE_HTTP_ENABLED` | `false` | Enable HTTP transport mode. |
| `PGEDGE_HTTP_ADDRESS` | `:8080` | HTTP listen address. |
| `PGEDGE_AUTH_ENABLED` | `true` | Enable API authentication (recommended for HTTP mode). |
| `INIT_TOKENS` | — | Comma-separated API tokens to initialize on startup. |
| `INIT_USERS` | — | User credentials to initialize (format: `user1:pass1,user2:pass2`). |

### Advanced

| Variable | Default | Description |
|----------|---------|-------------|
| `PGEDGE_CUSTOM_DEFINITIONS_PATH` | — | Path to YAML file with custom prompts and resources. |

## Documentation

For complete documentation including multi-database configuration, TLS setup, and advanced features:

- [pgEdge MCP Server Documentation](https://docs.pgedge.com/pgedge-postgres-mcp-server/)
- [GitHub Repository](https://github.com/pgEdge/pgedge-postgres-mcp)
- [pgEdge AI Toolkit Quickstart](https://www.pgedge.com/ai-toolkit)

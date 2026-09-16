# SurrealDB

Managed MCP server for SurrealDB Cloud: deploy, pause, resize and upgrade instances; run SurrealQL queries and CRUD operations against them; inspect and create schemas, tables and indexes; manage organisations, members, usage and billing; and read/write persistent agent memories in Spectron.

**Authentication:** OAuth 2.1 with dynamic client registration — `docker mcp oauth authorize surrealdb` opens a browser to sign in with your Surreal ID. A personal access token from [account.surrealdb.com/tokens](https://account.surrealdb.com/tokens) can be used instead via `Authorization: Bearer <token>` for scope-limited access.

Docs: https://surrealdb.com/docs/build/ai-agents/mcp

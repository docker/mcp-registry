# 404 Directory

Connect an AI Agent to a curated, read-only discovery and trust gateway with 12 tools. The flagship `search_official_docs` tool searches current first-party OpenAI, Microsoft Learn, AWS, and Cloudflare documentation in one call.

Before enabling the server, create a unique privacy-safe Agent ID at [404.directory/connect](https://404.directory/connect?source=docker-mcp-catalog) and paste the generated `agent:...` value into the Docker configuration field. The service stores only an irreversible HMAC digest for de-duplicated adoption measurement; it does not store the raw ID, prompts, tool arguments, or results.

- [Install guide](https://404.directory/connect?source=docker-mcp-catalog)
- [Documentation](https://github.com/MM-sheng/404-directory#readme)
- [Public Agent adoption metric](https://404.directory/v1/metrics/agents)

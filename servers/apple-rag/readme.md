Docs: https://github.com/BingoWon/apple-rag-mcp

## Features

- **Semantic Search**: Search Apple Developer Documentation using natural language queries
- **RAG (Retrieval-Augmented Generation)**: Fetch complete documentation content for AI processing
- **AI Reranking**: Results are reranked using AI for better relevance
- **Free Tier**: Available without authentication with rate limits
- **Authenticated Access**: Higher rate limits with Bearer Token authentication

## Authentication (Optional)

Authentication is optional. The server provides a free tier with basic rate limits.

For higher rate limits, you can authenticate using a Bearer Token:

1. Visit https://apple-rag.com/dashboard to get your token
2. Configure the token in Docker Desktop's MCP Toolkit

## Tools

- `search`: Search Apple Developer Documentation with semantic search
- `fetch`: Fetch complete documentation content by URL


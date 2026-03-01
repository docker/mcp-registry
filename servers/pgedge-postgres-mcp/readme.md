# pgEdge PostgreSQL MCP Server

Enterprise-grade PostgreSQL integration for AI applications with natural language query capabilities and advanced features.

## Overview

The pgEdge PostgreSQL MCP server enables AI applications to interact with PostgreSQL databases through a production-ready Model Context Protocol implementation. It combines powerful features like hybrid search, natural language processing, and cost-optimized operations.

## Key Features

### 🔍 Natural Language SQL
Query your PostgreSQL databases using plain English. The server translates natural language requests into optimized SQL queries.

### 🎯 Hybrid Search (pgvector + BM25)
Combines semantic search using pgvector with keyword-based BM25 search for superior search quality and relevance.

### 💰 90% Cost Reduction
Advanced prompt caching technology reduces API costs by up to 90% through intelligent caching of repeated operations.

### 🌐 Modern Web UI
User-friendly web interface for database interactions, query building, and result visualization.

### 🔒 Enterprise Security
Production-ready with comprehensive logging, connection pooling, and security best practices.

### ☁️ Universal Compatibility
Works with standard PostgreSQL 14+ including:
- Amazon RDS
- Google Cloud SQL
- Azure Database for PostgreSQL
- Self-hosted PostgreSQL instances

## Quick Start
https://docs.pgedge.com/pgedge-postgres-mcp-server/quickstart/

### Configuration

The server requires a PostgreSQL connection string set via the `DATABASE_URL` environment variable:

```bash
postgresql://username:password@hostname:5432/database
```

### Optional Configuration

- `LOG_LEVEL`: Set logging verbosity (DEBUG, INFO, WARNING, ERROR) - Default: INFO

## Use Cases

- **AI-Powered Analytics**: Enable LLMs to query and analyze database data
- **RAG Applications**: Build Retrieval-Augmented Generation systems with PostgreSQL backends
- **Natural Language Interfaces**: Allow users to query databases conversationally
- **Hybrid Search Applications**: Combine semantic and keyword search for better results
- **Database Automation**: Automate database operations through AI agents

## Documentation

Comprehensive documentation is available at:
**https://docs.pgedge.com/pgedge-postgres-mcp-server/**

Topics covered:
- Installation and configuration
- Authentication and security
- Query examples and best practices
- Hybrid search setup
- Web UI usage
- CLI client reference
- API documentation

## Source Code & Repository

**GitHub Repository**: https://github.com/pgEdge/pgedge-postgres-mcp

- ⭐ Star the project
- 🐛 Report issues
- 🔧 Contribute via pull requests
- 📖 Browse source code

## Support

- **Documentation**: https://docs.pgedge.com/pgedge-postgres-mcp-server/
- **GitHub Issues**: https://github.com/pgEdge/pgedge-postgres-mcp/issues
- **Website**: https://pgedge.com/ai-toolkit
- **Blog Post**: https://www.pgedge.com/blog/introducing-the-pgedge-postgres-mcp-server

## License

PostgreSQL License

## About pgEdge

pgEdge provides enterprise-grade distributed PostgreSQL solutions with a focus on modern application requirements including AI and edge computing capabilities.

**Learn More**: https://pgedge.com

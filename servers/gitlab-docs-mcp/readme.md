# GitLab Documentation MCP Server

A Model Context Protocol server that provides searchable access to GitLab's official documentation.

## Documentation

For full documentation, installation instructions, and usage examples, see:
https://github.com/nunolima/gitlab-docs-mcp

## Features

- Full-text search across GitLab documentation using SQLite FTS5
- Version-specific docs (18.7.x currently indexed)
- Multiple repositories indexed:
  - GitLab CE/EE (main application)
  - GitLab Runner (CI/CD runner)
  - Omnibus GitLab (installation packages)
  - Gitaly (Git RPC service)
  - GitLab Pages (static sites)
  - GitLab Agent (Kubernetes integration)

## Available Tools

- `search_docs` - Full-text search across all GitLab documentation
- `get_doc` - Retrieve specific documentation by file path
- `list_repositories` - List all indexed GitLab repositories

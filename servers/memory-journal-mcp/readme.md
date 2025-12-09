# Memory Journal MCP Server

Project context management for AI-assisted development. Bridge the gap between fragmented AI threads with persistent knowledge graphs and intelligent context recall.

## Documentation

- **[GitHub Repository](https://github.com/neverinfamous/memory-journal-mcp)**
- **[Wiki Documentation](https://github.com/neverinfamous/memory-journal-mcp/wiki)**
- **[Changelog](https://github.com/neverinfamous/memory-journal-mcp/wiki/CHANGELOG)**
- **[PyPI Package](https://pypi.org/project/memory-journal-mcp/)**
- **[Docker Hub](https://hub.docker.com/r/writenotenow/memory-journal-mcp)**

## Features

### 16 MCP Tools
- **Entry Management**: create_entry, update_entry, delete_entry, get_entry_by_id, create_entry_minimal
- **Search**: search_entries, semantic_search, search_by_date_range, get_recent_entries, list_tags
- **Analytics**: get_statistics, get_cross_project_insights
- **Relationships**: link_entries, visualize_relationships
- **Export**: export_entries
- **Testing**: test_simple

### 14 Workflow Prompts
- Daily standups, retrospectives, weekly digests
- PR workflows (summary, review prep, retrospective)
- GitHub Actions failure digest
- Goal tracking, period analysis
- Context bundle generation

### 13 MCP Resources
- Recent entries and significant milestones
- Relationship graphs (Mermaid visualization)
- Project/Issue/PR timelines
- GitHub Actions CI/CD timelines
- Team-shared entries

### GitHub Integration
- **Projects**: Auto-link entries to GitHub Projects
- **Issues**: Auto-detect from branch names
- **Pull Requests**: Track PR lifecycle
- **Actions**: CI/CD narrative graphs and failure analysis

### Advanced Features
- **Semantic Search**: Vector similarity with sentence-transformers
- **Knowledge Graphs**: 5 relationship types with Mermaid visualization
- **Tool Filtering**: Reduce token usage by up to 69%
- **Team Collaboration**: Git-synced databases for shared context

## Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `GITHUB_TOKEN` | GitHub PAT for Projects/Issues/PRs integration | Optional |
| `GITHUB_ORG_TOKEN` | GitHub token for org project access | Optional |
| `DEFAULT_ORG` | Default GitHub organization name | Optional |
| `MEMORY_JOURNAL_MCP_TOOL_FILTER` | Tool filter expression (e.g., `-test,-analytics`) | Optional |

### Tool Filtering

Reduce exposed tools for token efficiency:

```bash
# Disable test tools
MEMORY_JOURNAL_MCP_TOOL_FILTER="-test"

# Lightweight mode (core only)
MEMORY_JOURNAL_MCP_TOOL_FILTER="-search,-analytics,-relationships,-export,-admin,-test"
```

**Available Groups**: core, search, analytics, relationships, export, admin, test

## Quick Start

### Docker

```bash
docker run -i --rm \
  -v ./data:/app/data \
  writenotenow/memory-journal-mcp:v2.2.0 \
  python src/server.py
```

### PyPI

```bash
pip install memory-journal-mcp
memory-journal-mcp
```

## License

MIT License - See [LICENSE](https://github.com/neverinfamous/memory-journal-mcp/blob/main/LICENSE) for details.

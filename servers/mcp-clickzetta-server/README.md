# ClickZetta Lakehouse MCP Server

A Model Context Protocol (MCP) server for interacting with ClickZetta Lakehouse platform, providing comprehensive data management and analytics capabilities.

## Overview

This server enables AI assistants to interact with ClickZetta Lakehouse through 60+ tools for SQL operations, vector search, data management, and analytics. Perfect for data engineers and analysts working with ClickZetta's cloud-native data lakehouse.

## Features

- **SQL Operations**: Execute SELECT, INSERT, UPDATE, DELETE, and DDL statements
- **Vector & Semantic Search**: Knowledge retrieval and similarity search
- **Schema Management**: List tables, describe schemas, manage database objects
- **Data Import/Export**: Multiple format support (CSV, JSON, Parquet, etc.)
- **Advanced Analytics**: Query insights and data analysis
- **Multi-Connection Support**: Manage multiple ClickZetta environments

## Quick Start

### Prerequisites

1. **ClickZetta Account**: [Sign up here](https://accounts.clickzetta.com/register)
2. **Docker**: Ensure Docker is installed and running
3. **Configuration File**: Create `~/.clickzetta/connections.json`

### Configuration Template

Download and edit the configuration file:

```bash
# Create config directory
mkdir -p ~/.clickzetta

# Download template
curl -o ~/.clickzetta/connections.json \
  https://raw.githubusercontent.com/yunqiqiliang/mcp-clickzetta-server/main/config/connections-template.json

# Edit with your credentials
nano ~/.clickzetta/connections.json
```

Required fields in `connections.json`:
- `username`: Your ClickZetta username
- `password`: Your ClickZetta password
- `service`: API endpoint (e.g., `cn-shanghai-alicloud.api.clickzetta.com`)
- `instance`: Your instance name
- `workspace`: Workspace name
- `schema`: Schema name
- `vcluster`: Virtual cluster name

### Docker Run

```bash
docker run -i --rm \
  --stop-timeout 60 \
  -p 8502:8501 \
  -v ${HOME}/.clickzetta:/app/.clickzetta \
  czqiliang/mcp-clickzetta-server:latest
```

### MCP Client Configuration

For Claude Desktop (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "mcp-clickzetta-stdio": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "--stop-timeout", "60",
        "-p", "8502:8501",
        "-v", "${HOME}/.clickzetta:/app/.clickzetta",
        "czqiliang/mcp-clickzetta-server:latest"
      ]
    }
  }
}
```

## Usage Examples

Once connected, you can ask your AI assistant:

- "List all tables in my ClickZetta workspace"
- "Describe the structure of the customer_data table"
- "Run a SQL query to analyze sales by region"
- "Perform vector search for similar product descriptions"
- "Import CSV data into a new table"
- "Switch to my production ClickZetta connection"

## Configuration Options

The server supports flexible configuration through:

1. **Configuration File** (recommended): `~/.clickzetta/connections.json`
2. **Environment Variables**: For Docker deployments
3. **Command-line Parameters**: Override specific settings

### System Configuration

Optional settings in `connections.json`:

```json
{
  "system_config": {
    "allow_write": false,
    "prefetch": true,
    "log_level": "INFO",
    "exclude_tools": []
  },
  "connections": [...]
}
```

### Security

- **Read-only by default**: Set `allow_write: true` to enable write operations
- **SQL injection prevention**: Automatic input validation
- **Credential security**: Never commit credentials to version control
- **Volume mount**: Configuration file stays on your local machine

## Tools Overview

The server provides 60+ tools across categories:

- **Query Tools**: `read_query`, `write_query`, `create_table`
- **Schema Tools**: `list_tables`, `describe_table`, `show_object_list`
- **Search Tools**: `vector_search`, `match_all`
- **Data Import**: `import_data_into_table_from_url`, `import_data_into_table_from_database`
- **Analytics**: `append_insight`, `get_knowledge_about_how_to_do_something`
- **Connection Management**: `switch_lakehouse_connection`, `create_lakehouse_connection`
- **Index Management**: `create_index` (VECTOR, INVERTED, BLOOMFILTER)

Full tool documentation available in [tools.json](tools.json).

## Documentation

- **GitHub Repository**: [yunqiqiliang/mcp-clickzetta-server](https://github.com/yunqiqiliang/mcp-clickzetta-server)
- **Configuration Guide**: [Lakehouse MCP Configuration Guide](https://github.com/yunqiqiliang/mcp-clickzetta-server/blob/main/docs/user_guide/Lakehouse_MCP_CONFIGURATION_GUIDE.md)
- **User Guide**: [Complete User Guide](https://github.com/yunqiqiliang/mcp-clickzetta-server/blob/main/docs/user_guide/Lakehouse_MCP_USER_GUIDE.md)
- **Release Notes**: [v2.0.0 Release Notes](https://github.com/yunqiqiliang/mcp-clickzetta-server/blob/main/docs/RELEASE_NOTES_v2.0.0.md)

## Support

- **Issues**: [GitHub Issues](https://github.com/yunqiqiliang/mcp-clickzetta-server/issues)
- **ClickZetta Documentation**: [yunqi.tech/documents](https://www.yunqi.tech/documents)
- **License**: Apache-2.0

## Version

Latest version: 3.3+

Key updates:
- v3.3: Intelligent session recovery + standard MCP endpoints
- v2.0: Unified configuration system + simplified Docker deployment
- Multi-connection management and environment support

---

**Ready to start?** Create your configuration file and add the server to your MCP client!

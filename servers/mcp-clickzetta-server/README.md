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

Create the configuration file manually:

```bash
# Create config directory
mkdir -p ~/.clickzetta

# Create configuration file
cat > ~/.clickzetta/connections.json << 'EOF'
{
  "connections": [
    {
      "is_default": true,
      "connection_name": "my_clickzetta",
      "service": "cn-shanghai-alicloud.api.clickzetta.com",
      "username": "YOUR_USERNAME",
      "password": "YOUR_PASSWORD",
      "instance": "YOUR_INSTANCE",
      "workspace": "quick_start",
      "schema": "public",
      "vcluster": "default_ap"
    }
  ],
  "system_config": {
    "allow_write": false,
    "prefetch": true,
    "log_level": "INFO"
  }
}
EOF

# Edit with your actual credentials
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

## Links

- **Official ClickZetta Lakehouse MCP Server Guide**: [Lakehouse MCP Server Documentation](https://www.yunqi.tech/documents/LakehouseMCPServer_intro)
- **Docker Image**: [czqiliang/mcp-clickzetta-server](https://hub.docker.com/r/czqiliang/mcp-clickzetta-server)
- **ClickZetta Platform**: [yunqi.tech/documents](https://www.yunqi.tech/documents)
- **MCP Protocol**: [Model Context Protocol](https://modelcontextprotocol.io)

## License

Apache-2.0

## Configuration Reference

### Key Configuration Parameters

| Parameter | Description | Example |
|-----------|-------------|---------|
| `service` | API endpoint based on your region | `cn-shanghai-alicloud.api.clickzetta.com` |
| `username` | ClickZetta account username | `user@example.com` |
| `password` | ClickZetta account password | `your_password` |
| `instance` | Your instance name from console URL | `19d58db8` |
| `workspace` | Workspace name | `quick_start` |
| `schema` | Schema name | `public` |
| `vcluster` | Virtual cluster name | `default_ap` |

### Common Service Endpoints

- **China (Shanghai, Aliyun)**: `cn-shanghai-alicloud.api.clickzetta.com`
- **China (Beijing, Tencent)**: `ap-beijing-tencentcloud.api.clickzetta.com`
- **Singapore (AWS)**: `ap-southeast-1-aws.api.singdata.com`

### Troubleshooting

**Connection Issues**:
- Verify credentials at [accounts.clickzetta.com](https://accounts.clickzetta.com)
- Check service endpoint matches your region
- Ensure instance name is correct (from browser URL)

**Permission Issues**:
- Set `allow_write: false` for read-only access (recommended)
- Set `allow_write: true` only if you need data modification capabilities

**Tool Issues**:
- Use `exclude_tools` in system_config to disable specific tools
- Check logs with `log_level: DEBUG` for detailed information

---

**Ready to start?** Create your configuration file and add the server to your MCP client!

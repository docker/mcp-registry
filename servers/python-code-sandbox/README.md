# <img src="icon.svg" width="32" height="32" align="center" /> Python Code Sandbox

A secure and isolated Python execution environment implemented as an MCP server using Docker.

## Features
- **Ephemeral Runs**: Execute code in one-off containers with automatic cleanup.
- **Persistent Sessions**: Maintain state across multiple tool calls using sessions.
- **Auto-dependency**: On-the-fly pip installations from PyPI.
- **Plotting Support**: Automatically retrieves generated images and files.
- **File Persistence**: Smart default file persistence to host filesystem.
- **Bilingual Documentation**: Comprehensive English and Chinese documentation.

## Requirements
- Docker Desktop must be running.

## Usage

### Claude Desktop
Add to your Claude Desktop config:
```json
{
  "mcpServers": {
    "python-sandbox": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "-v", "/var/run/docker.sock:/var/run/docker.sock",
        "ghcr.io/li-xiu-qi/python-code-sandbox-mcp"
      ]
    }
  }
}
```

### Environment Variables
- `SANDBOX_MEMORY_LIMIT`: Memory limit per container (default: `2g`)
- `SANDBOX_CPU_QUOTA`: CPU quota per container (default: `50000`)
- `ENABLE_PIP_CACHE`: Enable pip caching (default: `true`)
- `SANDBOX_FILES_DIR`: Custom directory for file persistence (optional)

## Documentation
- [GitHub Repository](https://github.com/li-xiu-qi/python-code-sandbox-mcp)
- [Execution Modes Guide](https://github.com/li-xiu-qi/python-code-sandbox-mcp/blob/main/docs/en/EXECUTION_MODES.md)
- [API Reference](https://github.com/li-xiu-qi/python-code-sandbox-mcp/blob/main/docs/en/API.md)

## Version
Current: v0.2.0

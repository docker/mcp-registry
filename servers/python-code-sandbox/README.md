# <img src="icon.svg" width="32" height="32" align="center" /> Python Code Sandbox

A secure and isolated Python execution environment implemented as an MCP server using Docker.

## Features
- **Ephemeral Runs**: Execute code in one-off containers.
- **Persistent Sessions**: Maintain state across multiple tool calls.
- **Auto-dependency**: On-the-fly pip installations.
- **Plotting Support**: Automatically retrieves generated images/files.

## Requirements
- Docker Desktop must be running.

## Usage
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

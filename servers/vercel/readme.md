# Vercel MCP Server

Manage Vercel projects and deployments through AI assistants using the Model Context Protocol.

## Features

The Vercel MCP server provides two categories of tools:

### Public Tools (No Authentication Required)
- **Documentation Search**: Search and navigate Vercel documentation
- **Information Retrieval**: Get information about Vercel features and services

### Authenticated Tools (Require Vercel Authentication)
- **Project Management**: Manage Vercel projects and settings
- **Deployment Management**: Deploy, view, and manage deployments
- **Log Analysis**: Analyze deployment logs and troubleshoot issues
- **Team Collaboration**: Work with team projects and resources

## Authentication

The server uses **OAuth** authentication following the latest MCP Authorization specifications. Each client connection requires explicit user consent to protect against confused deputy attacks.

## Configuration

### General Access
Connect to Vercel MCP for general access:
```
URL: https://mcp.vercel.com
```

### Project-Specific Access
For enhanced context with specific projects:
```
URL: https://mcp.vercel.com/<teamSlug>/<projectSlug>
```

You can add multiple Vercel MCP connections with different names for different projects.

## Transport

- **Type**: Streamable HTTP
- **Protocol Version**: MCP 2025-06-18
- **Authentication**: OAuth 2.1

## Supported Clients

- Claude Code
- Claude.ai
- ChatGPT
- Cursor
- VS Code with Copilot
- Devin
- Raycast
- Goose
- Windsurf
- Gemini Code Assist
- Gemini CLI
- Any MCP-compatible client with HTTP transport support

## Documentation

For more information, visit:
- [Vercel MCP Documentation](https://vercel.com/docs/mcp/vercel-mcp)
- [Vercel Documentation](https://vercel.com/docs)
- [MCP Specification](https://modelcontextprotocol.io)

# NetBird MCP Server

Comprehensive Model Context Protocol server for managing NetBird VPN infrastructure through AI assistants.

## Features

- **Complete API Coverage**: 64 tools covering all NetBird management operations
- **Account Management**: View and configure account settings
- **Peer Management**: List, view, update, and delete VPN peers
- **Group Management**: Organize peers into groups with advanced helper tools
- **Policy Management**: Create and manage network access policies with template support
- **Network Management**: Configure networks, resources, and routers
- **Route Management**: Set up routing rules and exit nodes
- **DNS Management**: Configure nameservers and DNS settings
- **Setup Keys**: Generate and manage peer enrollment keys
- **User Management**: Invite and manage team members
- **Posture Checks**: Enforce security compliance requirements
- **Port Allocations**: Manage ingress port forwarding

## Quick Start

### 1. Get Your NetBird API Token

Generate an API token from your NetBird dashboard:
- **Cloud**: https://app.netbird.io/settings/tokens
- **Self-hosted**: https://your-domain/settings/tokens

### 2. Configure in Docker Desktop

1. Open Docker Desktop → MCP Toolkit
2. Search for "NetBird MCP Server"
3. Click "Enable"
4. Enter your API token in the secret field
5. Set API host:
   - Use `api.netbird.io` for NetBird Cloud
   - Use your domain for self-hosted instances
6. Save configuration

### 3. Use with AI Assistants

Connect your AI assistant (Claude Desktop, Kiro, etc.) to Docker MCP Gateway and start managing your NetBird infrastructure with natural language commands.

## Documentation

- [Full Documentation](https://github.com/XNet-NGO/mcp-netbird#readme)
- [MCP Setup Guide](https://github.com/XNet-NGO/mcp-netbird/blob/main/docs/MCP_SETUP_GUIDE.md)
- [NetBird API Documentation](https://docs.netbird.io/api)

## Support

- [Report Issues](https://github.com/XNet-NGO/mcp-netbird/issues)
- [GitHub Discussions](https://github.com/XNet-NGO/mcp-netbird/discussions)

## License

Apache License 2.0 - See [LICENSE](https://github.com/XNet-NGO/mcp-netbird/blob/main/LICENSE)

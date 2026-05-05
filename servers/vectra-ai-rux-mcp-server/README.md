# Vectra AI MCP Server

Connect your Vectra AI security platform to AI assistants like Claude, Cursor, and Windsurf through the Model Context Protocol (MCP).

## Overview

The Vectra AI MCP Server standardizes how Large Language Models (LLMs) interact with Vectra AI's threat detection and response platform. It provides AI assistants with direct access to your Vectra AI environment, enabling them to:

- Analyze security detections and threats
- Manage investigation assignments
- Query entity information (hosts and accounts)
- Access detection summaries and details
- Visualize threat relationships
- Manage platform users and assignments
- Work with lockdown entities

## Features

### Detection Management
- List and filter security detections with detailed information
- Get comprehensive detection details and summaries
- Retrieve detection PCAP files for analysis
- Count detections based on various criteria
- Mark detections as fixed or unfixed

### Entity Management
- List and query host and account entities
- Get detailed entity information by ID, name, or IP address
- Filter entities by state, priority, and tags
- Access entity detection history

### Investigation & Response
- Create and manage investigation assignments
- Add investigation notes to entities
- List assignments by user or entity
- Delete investigation assignments
- Work with lockdown entities

### Platform Administration
- List platform users with role filtering
- Manage user assignments and permissions
- Access platform configuration and logs

### Advanced Analytics
- Generate detection summaries with AI analysis
- Visualize entity detection relationships with interactive graphs
- Export detection data for further analysis

## Configuration

### Required
- **VECTRA_CLIENT_SECRET**: Your Vectra AI service account token (get from [Vectra AI Portal](https://portal.vectra.ai))
- **VECTRA_BASE_URL**: Your Vectra AI platform URL (e.g., `https://123456789.ab1.portal.vectra.ai`)
- **VECTRA_CLIENT_ID**: Your Vectra AI client ID

### Optional
- **VECTRA_MCP_DEBUG**: Enable debug logging - default: `false`

## Security Best Practices

⚠️ **Important Security Considerations:**

1. **Service Account Tokens**: Use dedicated service account tokens with minimal required permissions
2. **Network Security**: Ensure secure network connections to your Vectra AI platform
4. **Review Operations**: Always review and approve security operations before execution

## Usage

The MCP server provides the following tool categories:

- **Detection Management**: Query, analyze, and manage security detections
- **Entity Management**: Work with hosts, accounts, and other entities
- **Investigation Tools**: Create assignments, add notes, and manage investigations
- **Platform Administration**: Manage users and platform configuration
- **Analytics & Visualization**: Generate summaries and visualizations

## Links

- [Vectra AI MCP Server Repository](https://github.com/vectra-ai-research/vectra-ai-mcp-server)
- [Vectra AI Platform Documentation](https://docs.vectra.ai/)
- [Model Context Protocol](https://modelcontextprotocol.io/)
- [Vectra AI Blog - MCP Server Introduction](https://www.vectra.ai/blog/introducing-the-vectra-ai-mcp-server)

## License

MIT License

## Support

For issues and questions:
- [GitHub Issues](https://github.com/vectra-ai-research/vectra-ai-mcp-server/issues)
- [Vectra AI Support](https://support.vectra.ai/)

---

> **Note**: MCP (Model Context Protocol) is an emerging technology. Exercise caution when using this server and follow security best practices, including proper credential management and network security measures.

# Supabase MCP Server

Connect AI assistants like Cursor and Claude to Supabase projects, enabling them to manage databases, execute queries, and deploy functions through the Model Context Protocol.

## Features

The Supabase MCP server provides tools for:

### Database Management
- Execute SQL queries
- Manage database schemas
- View table structures
- Data operations (CRUD)

### Function Deployment
- Deploy Edge Functions
- Manage serverless functions
- Configure function settings

### Project Management
- Manage Supabase projects
- Configure project settings
- Access project resources

## Authentication

The server uses Streamable HTTP transport with Dynamic Client Registration OAuth 2.1 authentication. You'll authenticate through your Supabase account during initial MCP client setup.

## Security Considerations

- **Read-only mode**: Use read-only access for safer operations
- **Project scoping**: Scope servers to specific projects to limit AI access
- **Access control**: Configure appropriate permissions for your use case

## Supported Clients

- Cursor
- VS Code / VS Code Insiders
- Claude Desktop
- Any MCP-compatible client with HTTP transport support

## Documentation

For more information, visit:
- [Supabase MCP GitHub](https://github.com/supabase-community/supabase-mcp)
- [Supabase Documentation](https://supabase.com/docs)
- [MCP Documentation](https://modelcontextprotocol.io)

# Supabase MCP Server

Connect your Supabase projects to AI assistants like Claude, Cursor, and Windsurf through the Model Context Protocol (MCP).

## Overview

The Supabase MCP Server standardizes how Large Language Models (LLMs) interact with Supabase services. It provides AI assistants with direct access to your Supabase project, enabling them to:

- Manage database tables and schemas
- Execute SQL queries and migrations
- Deploy and manage Edge Functions
- Access project configuration and logs
- Generate TypeScript types from your schema
- Work with development branches
- Search Supabase documentation

## Features

### Database Management
- List tables, extensions, and migrations
- Execute SQL queries (read-only mode available)
- Apply database migrations with automatic tracking
- Generate TypeScript types from your database schema

### Edge Functions
- List and inspect Edge Functions
- Deploy new functions or update existing ones
- Access function source code

### Development & Debugging
- Get project URLs and API keys
- View service logs (API, Postgres, Auth, Storage, Realtime)
- Check security and performance advisories
- Work with development branches (paid plans)

### Knowledge Base
- Search Supabase documentation for up-to-date information

## Configuration

### Required
- **SUPABASE_ACCESS_TOKEN**: Your Supabase access token (get from [Supabase Dashboard](https://supabase.com/dashboard/account/tokens))

### Optional
- **PROJECT_REF**: Scope to a specific project (recommended for security)
- **API_URL**: Custom Supabase API URL (defaults to https://api.supabase.com)
- **READ_ONLY**: Set to `true` to restrict to read-only operations (recommended)

## Security Best Practices

⚠️ **Important Security Considerations:**

1. **Use Read-Only Mode**: Enable read-only mode by default to prevent accidental data modifications
2. **Project Scoping**: Limit access to a specific project using PROJECT_REF
3. **Development Environment**: Use with development projects, not production
4. **Review Tool Calls**: Always review and approve tool calls before execution
5. **Feature Groups**: Enable only the tool groups you need

## Usage

The MCP server provides the following tool groups:

- **account**: Project and organization management
- **docs**: Documentation search
- **database**: Schema and data management
- **debugging**: Logs and advisories
- **development**: Project configuration
- **functions**: Edge Functions deployment
- **branching**: Development branch management (experimental)
- **storage**: Storage bucket configuration

## Links

- [Supabase MCP Documentation](https://supabase.com/docs/guides/getting-started/mcp)
- [Source Repository](https://github.com/Sixteen-Digits/supabase-mcp)
- [Supabase Documentation](https://supabase.com/docs)
- [Model Context Protocol](https://modelcontextprotocol.io/)

## License

Apache 2.0

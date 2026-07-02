# Dart Tooling MCP Server

Expose Dart and Flutter development tools to AI assistant clients through the Model Context Protocol.

## Features

The Dart Tooling MCP server provides **40+ tools** for AI agents working with Dart/Flutter projects:

### Project Analysis
- Analyze Dart/Flutter projects
- Inspect project structure
- Review dependencies
- Code analysis and diagnostics

### Testing
- Run tests
- View test results
- Debug test failures
- Test coverage analysis

### Widget Inspection
- Inspect Flutter widgets
- View widget tree
- Analyze widget properties
- Debug UI layouts

### Hot Reload/Restart
- Hot reload during development
- Hot restart application
- Fast development iteration
- Live code updates

### Package Management
- Manage pub packages
- Add/remove dependencies
- Update packages
- Resolve dependency conflicts

## Requirements

- **Dart**: 3.9.0-163.0.dev or later
- For Flutter projects: Current stable Flutter SDK

## Configuration

The server supports various flags:
- `--experimental-mcp-server`: For Dart < 3.9.0
- `--force-roots-fallback`: Cursor workaround for multi-root workspaces

## Compatible Clients

- Gemini CLI
- Cursor
- GitHub Copilot
- VS Code (via MCP API)
- Any stdio-based MCP client supporting Tools and Resources capabilities

## Documentation

For more information, visit:
- [Dart Tooling MCP Server](https://github.com/dart-lang/ai/tree/main/pkgs/dart_mcp_server)
- [Dart Documentation](https://dart.dev/guides)
- [Flutter Documentation](https://flutter.dev/docs)

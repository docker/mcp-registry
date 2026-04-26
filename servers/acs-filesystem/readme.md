# MCP Filesystem Server

Advanced filesystem operations for AI agents with strict security boundaries.

## Overview

The MCP Filesystem Server provides AI agents with advanced file operations beyond basic read/write, including batch operations, directory watching, file search/indexing, and permission management, all within strict security boundaries.

## Features

### Core Capabilities

- **Batch Operations**: Execute multiple file operations (copy, move, delete) atomically with automatic rollback on failure
- **Directory Watching**: Monitor filesystem changes in real-time with event filtering and recursive watching
- **File Search & Indexing**: Fast full-text search with metadata filtering using Lunr.js
- **Checksum Operations**: Compute and verify file integrity using MD5, SHA-1, SHA-256, or SHA-512
- **Symlink Management**: Create and manage symbolic links within workspace boundaries
- **Disk Usage Analysis**: Analyze directory sizes, identify large files, and get file type breakdowns
- **Directory Operations**: Recursive copy, sync (only newer/missing files), and atomic file replacement

### Security Features

The server implements defense-in-depth security with 10 layers of path validation:

1. **Absolute Path Resolution**: Prevents relative path tricks
2. **Workspace Boundary Check**: Ensures path is within workspace
3. **Path Traversal Detection**: Blocks `..` and `./` sequences
4. **System Path Blocklist**: Hardcoded system directories (cannot be overridden)
5. **Sensitive Pattern Blocklist**: Hardcoded sensitive files (cannot be overridden)
6. **Subdirectory Restrictions**: Optional allowlist within workspace
7. **User Blocklist**: Custom blocked paths
8. **User Pattern Blocklist**: Custom blocked patterns
9. **Read-Only Mode**: Prevents write/delete operations
10. **Symlink Validation**: Validates symlink targets are within workspace

### Hardcoded Security (Cannot Be Disabled)

**System Paths (Always Blocked):**

- `/etc`, `/sys`, `/proc`, `/dev`, `/boot`, `/root`
- `C:\Windows`, `C:\Program Files`
- `/System`, `/Library`, `/Applications` (macOS)
- `/bin`, `/sbin`, `/usr/bin`, `/usr/sbin`

**Sensitive Patterns (Always Blocked):**

- `.ssh/`, `.aws/`, `.kube/`
- `id_rsa`, `*.pem`, `*.key`, `*.p12`, `*.pfx`
- Files containing: `password`, `secret`, `token`
- `.env` files

## Installation

### NPM

```bash
npm install -g @ai-capabilities-suite/mcp-filesystem
```

### Docker

```bash
docker pull digitaldefiance/mcp-filesystem:latest
```

## Configuration

### Required Configuration

Create a configuration file (e.g., `mcp-filesystem-config.json`):

```json
{
  "workspaceRoot": "/path/to/your/workspace",
  "blockedPaths": [".git", ".env", "node_modules"],
  "blockedPatterns": ["*.key", "*.pem", "*.env"],
  "maxFileSize": 104857600,
  "maxBatchSize": 1073741824,
  "maxOperationsPerMinute": 100,
  "enableAuditLog": true,
  "readOnly": false
}
```

### MCP Client Configuration

Add to your MCP client configuration:

```json
{
  "mcpServers": {
    "filesystem": {
      "command": "mcp-filesystem",
      "args": ["--config", "/path/to/mcp-filesystem-config.json"]
    }
  }
}
```

## Available Tools

The server exposes 12 MCP tools:

1. **fs_batch_operations** - Execute multiple operations atomically
2. **fs_watch_directory** - Monitor directory for changes
3. **fs_get_watch_events** - Retrieve accumulated events
4. **fs_stop_watch** - Stop watch session
5. **fs_search_files** - Search by name, content, or metadata
6. **fs_build_index** - Build searchable file index
7. **fs_create_symlink** - Create symbolic link
8. **fs_compute_checksum** - Compute file checksum
9. **fs_verify_checksum** - Verify file integrity
10. **fs_analyze_disk_usage** - Analyze disk usage
11. **fs_copy_directory** - Recursively copy directory
12. **fs_sync_directory** - Sync directories (only newer/missing)

## Usage Examples

### Batch File Operations

```typescript
{
  "operations": [
    { "type": "copy", "source": "file1.txt", "destination": "backup/file1.txt" },
    { "type": "move", "source": "temp.txt", "destination": "archive/temp.txt" },
    { "type": "delete", "source": "old.txt" }
  ],
  "atomic": true
}
```

### Watch Directory

```typescript
{
  "path": "src",
  "recursive": true,
  "filters": ["*.ts", "*.js"]
}
```

### Search Files

```typescript
{
  "query": "TODO",
  "searchType": "content",
  "fileTypes": [".ts", ".js"],
  "useIndex": true
}
```

### Verify File Integrity

```typescript
{
  "path": "important-file.zip",
  "checksum": "abc123...",
  "algorithm": "sha256"
}
```

## Security Warnings

⚠️ **CRITICAL SECURITY CONSIDERATIONS**

1. **Workspace Jail**: All operations are confined to the configured workspace root. This cannot be changed after server starts.

2. **System Paths**: The server blocks access to system directories. This is hardcoded and cannot be disabled.

3. **Sensitive Files**: The server blocks access to SSH keys, AWS credentials, and other sensitive files. This is hardcoded and cannot be disabled.

4. **Rate Limiting**: Configure appropriate rate limits to prevent abuse.

5. **Audit Logging**: Enable audit logging for security monitoring and forensics.

6. **Read-Only Mode**: Consider using read-only mode for untrusted agents.

## Documentation

- **Full Documentation**: [GitHub Repository](https://github.com/Digital-Defiance/ai-capabilities-suite/tree/main/packages/mcp-filesystem)
- **Security Guide**: [SECURITY.md](https://github.com/Digital-Defiance/ai-capabilities-suite/blob/main/packages/mcp-filesystem/SECURITY.md)
- **Docker Guide**: [DOCKER.md](https://github.com/Digital-Defiance/ai-capabilities-suite/blob/main/packages/mcp-filesystem/DOCKER.md)
- **API Reference**: [README.md](https://github.com/Digital-Defiance/ai-capabilities-suite/blob/main/packages/mcp-filesystem/README.md)

## Support

- **Issues**: [GitHub Issues](https://github.com/Digital-Defiance/ai-capabilities-suite/issues)
- **Email**: info@digitaldefiance.org

## License

MIT License - see [LICENSE](https://github.com/Digital-Defiance/ai-capabilities-suite/blob/main/packages/mcp-filesystem/LICENSE)

## Related Projects

- [MCP Debugger](https://github.com/Digital-Defiance/ai-capabilities-suite/tree/main/packages/mcp-debugger-server) - Debug Node.js applications via MCP
- [MCP Process](https://github.com/Digital-Defiance/ai-capabilities-suite/tree/main/packages/mcp-process) - Process management via MCP
- [MCP Screenshot](https://github.com/Digital-Defiance/ai-capabilities-suite/tree/main/packages/mcp-screenshot) - Screenshot capture via MCP

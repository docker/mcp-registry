# Disk Clean MCP

A Model Context Protocol (MCP) server that provides read-only disk usage analysis tools. Help AI assistants analyze storage, identify space hogs, and suggest cleanup opportunities without modifying any files.

## Features

- **📊 Quick Summary**: Get instant overview of directory size, file count, and structure
- **📁 Type Analysis**: Find which file types consume the most space
- **🔍 Directory Ranking**: Identify the heaviest directories
- **📈 Large Files**: Locate biggest files with flexible filtering
- **⏰ Stale Files**: Find old, unused files taking up space
- **🔄 Duplicate Detection**: Identify duplicate files by content hash

## Safety First

This server is **read-only** by design. It only analyzes and reports—never deletes, moves, or modifies files. Perfect for safe AI-assisted disk cleanup workflows.

## Installation

### Option 1: NPM (Recommended)

```bash
npm install -g disk-clean-mcp
```

### Option 2: Use with Claude Desktop

Add to your Claude Desktop configuration:

```json
{
  "mcpServers": {
    "disk-clean": {
      "command": "npx",
      "args": ["-y", "disk-clean-mcp"]
    }
  }
}
```

### Option 3: Docker

```bash
docker run -v /path/to/analyze:/data disk-clean-mcp
```

## Usage Examples

### Analyze a directory
```
"What's taking up space in my Downloads folder?"
```

### Find large old files
```
"Show me files larger than 100MB that haven't been accessed in 6 months"
```

### Detect duplicates
```
"Find duplicate files in my Documents folder"
```

## Available Tools

1. **scan_summary** - Get directory overview
2. **by_type** - Analyze by file extension
3. **top_dirs** - Find heaviest directories
4. **top_files** - List largest files
5. **stale_candidates** - Find old unused files
6. **duplicate_candidates** - Detect duplicate content

## Configuration

- Supports ignore patterns (glob syntax)
- Configurable depth limits
- Size and age filters
- Result limits to prevent overwhelming output

## License

MIT

## Links

- [GitHub Repository](https://github.com/Superandyfre/disk-clean-mcp)
- [NPM Package](https://www.npmjs.com/package/disk-clean-mcp)
- [Report Issues](https://github.com/Superandyfre/disk-clean-mcp/issues)

## Contributing

Contributions welcome! Please open an issue or PR on GitHub.

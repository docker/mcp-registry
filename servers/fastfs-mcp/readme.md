# FastFS MCP

High-performance filesystem and Git operations for AI-assisted development workflows.

## Overview

FastFS-MCP is a comprehensive MCP server providing 69 tools for filesystem operations, Git repository management, and code navigation. Built with performance in mind, it enables AI assistants like Claude to seamlessly interact with local filesystems and Git repositories through Docker.

**Key capabilities:**
- **60+ filesystem tools**: Complete UNIX-style file operations (ls, find, grep, sed, awk, tar, etc.)
- **30+ Git tools**: Full Git workflow support plus advanced repository analysis
- **GitHub authentication**: Secure access via Personal Access Tokens or GitHub Apps
- **High performance**: Optimized for speed with tools like ripgrep and efficient file handling
- **Docker-native**: Runs in isolated container with /mnt/workspace mount point

Perfect for AI-assisted coding, repository analysis, project exploration, and automated file management tasks.

## Tools

### Filesystem Operations (40+ tools)

#### Basic File Operations
- **`fastfs_ls`**: List files and directories at a path
- **`fastfs_pwd`**: Print current working directory
- **`fastfs_cd`**: Change current directory
- **`fastfs_read`**: Read file contents (UTF-8 text files)
- **`fastfs_write`**: Write/overwrite file contents (creates parent dirs)
- **`fastfs_grep`**: Search file with pattern and line numbers
- **`fastfs_cat`**: Concatenate and display multiple files
- **`fastfs_head`**: Show first N lines of file (default 10)
- **`fastfs_tail`**: Show last N lines of file (default 10)

#### File Management
- **`fastfs_cp`**: Copy files or directories (recursive support)
- **`fastfs_mv`**: Move or rename files/directories
- **`fastfs_rm`**: Remove files or directories (recursive, force options)
- **`fastfs_mkdir`**: Create directories (with parent support)
- **`fastfs_touch`**: Create empty file or update timestamp
- **`fastfs_chmod`**: Change permissions (octal or symbolic modes)
- **`fastfs_chown`**: Change file owner and group

#### File Discovery & Analysis
- **`fastfs_find`**: Find files by glob pattern, type, max depth
- **`fastfs_tree`**: Display directory tree with depth limit
- **`fastfs_stat`**: Show detailed file metadata
- **`fastfs_which`**: Locate executables in PATH
- **`fastfs_du`**: Show disk usage of directory
- **`fastfs_df`**: Show disk space and usage
- **`fastfs_readlink`**: Resolve symbolic link target
- **`fastfs_realpath`**: Get absolute resolved path

#### Text Processing
- **`fastfs_sed`**: Stream editor for text transformation
- **`fastfs_gawk`**: AWK text processing and filtering
- **`fastfs_cut`**: Extract columns by delimiter
- **`fastfs_sort`**: Sort lines (alphabetic, numeric, reverse)
- **`fastfs_uniq`**: Filter/report repeated lines
- **`fastfs_wc`**: Count lines, words, bytes
- **`fastfs_nl`**: Number lines in file
- **`fastfs_split`**: Split file into smaller parts

#### Archive & Compression
- **`fastfs_tar`**: Create/extract/list tar archives (gzip, bzip2, xz)
- **`fastfs_gzip`**: Compress/decompress with gzip
- **`fastfs_zip`**: Create/extract zip archives

### Git Operations (29 tools)

#### Basic Git Workflow
- **`fastfs_clone`**: Clone repository (GitHub auth support)
- **`fastfs_init`**: Initialize new repository
- **`fastfs_add`**: Stage files for commit
- **`fastfs_commit`**: Commit staged changes
- **`fastfs_status`**: Show working tree status
- **`fastfs_push`**: Push to remote
- **`fastfs_pull`**: Pull from remote
- **`fastfs_fetch`**: Download refs and objects

#### Branching & Merging
- **`fastfs_branch`**: List/create/delete branches
- **`fastfs_checkout`**: Switch branches or restore files
- **`fastfs_merge`**: Merge branch into current
- **`fastfs_rebase`**: Reapply commits on new base

#### History & Inspection
- **`fastfs_log`**: Show commit history
- **`fastfs_show`**: Show commit details
- **`fastfs_diff`**: Show changes between versions
- **`fastfs_blame`**: Show line-by-line authorship
- **`fastfs_git_grep`**: Search tracked files efficiently

#### Repository Management
- **`fastfs_remote`**: Manage remote repositories
- **`fastfs_tag`**: Create/list/delete tags
- **`fastfs_config`**: Get/set Git configuration
- **`fastfs_stash`**: Temporarily save uncommitted changes
- **`fastfs_reset`**: Reset HEAD to state (DESTRUCTIVE)
- **`fastfs_clean`**: Remove untracked files (DESTRUCTIVE)

#### Advanced Analysis Tools
- **`fastfs_context`**: Get comprehensive repo status in one call (GO-TO tool)
  - Current branch, repo root, clean status
  - HEAD commit, remotes, recent commits
  - Branches and tags lists
- **`fastfs_repo_info`**: Detailed repository statistics
  - Commit count, contributors, file count, size
  - Per-author commit counts
- **`fastfs_validate`**: Check repository health
  - Identifies common issues and warnings
- **`fastfs_summarize_log`**: Analyze commit history
  - Author statistics, date distribution
  - Per-commit change metrics
- **`fastfs_suggest_commit`**: Auto-generate commit messages
  - Analyzes staged changes
  - Suggests conventional commit format
- **`fastfs_audit_history`**: Security audit of history
  - Finds large files, potential secrets
  - Identifies merge conflicts, binary files

#### Utility Tools
- **`fastfs_rev_parse`**: Parse Git revisions
- **`fastfs_ls_files`**: List files in index/working tree
- **`fastfs_describe`**: Human-readable object names
- **`fastfs_git_show_head`**: Show current HEAD details
- **`fastfs_version`**: Get Git version

## Configuration

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `GITHUB_PERSONAL_ACCESS_TOKEN` | No | None | GitHub PAT for private repo access |
| `GITHUB_APP_ID` | No | None | GitHub App ID for GitHub App auth |
| `GITHUB_APP_INSTALLATION_ID` | No | None | GitHub App installation ID |

### Secrets

| Secret | Required | Description |
|--------|----------|-------------|
| `GITHUB_PERSONAL_ACCESS_TOKEN` | No | GitHub Personal Access Token for authenticating with private repositories |
| `GITHUB_APP_PRIVATE_KEY` | No | GitHub App private key (PEM format) for GitHub App authentication |

### Volumes

| Volume | Path | Description |
|--------|------|-------------|
| `workspace` | `/mnt/workspace` | Local filesystem mount point for all file operations |

## Usage Examples

### Basic File Operations

```
# List current directory
fastfs_ls(path=".")

# Read a file
fastfs_read(path="README.md")

# Search for pattern
fastfs_grep(pattern="TODO", path="src/main.py")

# Find Python files
fastfs_find(pattern="*.py", file_type="f")
```

### Git Workflow

```
# Get full repository context (START HERE)
fastfs_context()

# Check status before changes
fastfs_status()

# Stage and commit
fastfs_add(paths=".")
fastfs_suggest_commit()  # Get AI-suggested commit message
fastfs_commit(message="feat: add new feature")

# Push changes
fastfs_push()
```

### Repository Analysis

```
# Get repository statistics
fastfs_repo_info()

# Audit for security issues
fastfs_audit_history()

# Analyze recent commits
fastfs_summarize_log(count=20)

# Check repository health
fastfs_validate()
```

### Project Exploration

```
# Visualize structure
fastfs_tree(path=".", depth=3)

# Find all config files
fastfs_find(pattern="*.{json,yaml,yml}")

# Search codebase
fastfs_git_grep(pattern="function.*export")

# Check file details
fastfs_stat(path="package.json")
```

### Working with Archives

```
# Create tar.gz archive
fastfs_tar(operation="create", archive_file="backup.tar.gz", files=["src/", "docs/"])

# Extract zip archive
fastfs_zip(operation="extract", archive_file="download.zip")

# Compress file
fastfs_gzip(path="large-file.log")
```

## Authentication

FastFS-MCP supports two GitHub authentication methods:

### Personal Access Token (Recommended for individuals)

Set the `GITHUB_PERSONAL_ACCESS_TOKEN` environment variable:

```bash
docker run -i --rm \
  -v /path/to/workspace:/mnt/workspace:rw \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=ghp_your_token \
  mcp/fastfs-mcp
```

### GitHub App (Recommended for organizations)

Provides fine-grained permissions and better security:

```bash
docker run -i --rm \
  -v /path/to/workspace:/mnt/workspace:rw \
  -e GITHUB_APP_ID=123456 \
  -e GITHUB_APP_PRIVATE_KEY="$(cat key.pem)" \
  -e GITHUB_APP_INSTALLATION_ID=789012 \
  mcp/fastfs-mcp
```

Or use a file-based private key:

```bash
docker run -i --rm \
  -v /path/to/workspace:/mnt/workspace:rw \
  -e GITHUB_APP_ID=123456 \
  -e GITHUB_APP_PRIVATE_KEY_PATH=/mnt/workspace/key.pem \
  -e GITHUB_APP_INSTALLATION_ID=789012 \
  mcp/fastfs-mcp
```

## Best Practices

1. **Start with context**: Use `fastfs_context()` to understand repository state before making changes
2. **Check before commit**: Always run `fastfs_status()` before `fastfs_add()` and `fastfs_commit()`
3. **Preview destructive operations**: Use `fastfs_clean(options="-n")` to preview before `fastfs_clean(options="-f")`
4. **Use tree for navigation**: `fastfs_tree()` provides better overview than multiple `fastfs_ls()` calls
5. **Leverage analysis tools**: Use `fastfs_suggest_commit()`, `fastfs_validate()`, and `fastfs_audit_history()` for insights

## Performance Notes

- Uses ripgrep for fast file searching
- Optimized for large repositories with efficient Git operations
- Tree operations limited by depth to prevent excessive output
- All operations run in isolated Docker container

## Security Considerations

- Only mount directories you need to expose
- Use GitHub App authentication over PAT for better security
- Store private keys with proper file permissions (chmod 600)
- Regularly review executed commands in logs
- Be cautious with destructive operations (rm, reset --hard, clean -f)

## Links

- [Source Repository](https://github.com/aj-geddes/fastfs-mcp)
- [Issue Tracker](https://github.com/aj-geddes/fastfs-mcp/issues)
- [Docker Hub](https://hub.docker.com/r/mcp/fastfs-mcp)

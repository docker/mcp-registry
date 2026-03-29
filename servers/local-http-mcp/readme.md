# Local HTTP Bridge

Production-ready MCP server that bridges HTTP requests from Claude to your private network resources.

## Overview

When Claude runs in containerized environments (like Docker Desktop), it cannot access resources on your host machine's private network - localhost services, internal APIs with custom DNS (like `*.hvs` domains), or development servers. The Local HTTP Bridge solves this by creating a secure, allowlist-based HTTP proxy that runs on the host machine.

This enables Claude to:
- Access localhost development servers (e.g., `http://localhost:3000`)
- Query internal APIs with custom DNS entries (e.g., `https://api.internal.hvs`)
- Interact with private network services (e.g., `http://192.168.1.100:8080`)
- Test and debug local web applications during development

Security is maintained through domain allowlisting - only pre-approved domains and IP patterns can be accessed.

## Tools

### `fetch`
Makes HTTP requests to local network resources that are only accessible from the host machine.

**Parameters:**
- `url` (string, required): Target URL to request. Must be in the allowlist and begin with http:// or https://. Examples: `http://localhost:3000/api`, `https://internal.hvs/health`
- `method` (string, optional): HTTP method to use. One of: GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS. Defaults to GET.
- `headers` (object, optional): Custom HTTP headers as key-value pairs. Example: `{"Authorization": "Bearer token", "Content-Type": "application/json"}`
- `body` (string, optional): Request body for POST/PUT/PATCH operations. Sent as-is to the target server. For JSON, stringify the object first.
- `verify_ssl` (boolean, optional): Whether to verify SSL certificates. Set to false for self-signed certificates in development. Defaults to true.
- `timeout` (number, optional): Request timeout in seconds. Defaults to 30, maximum 300.
- `follow_redirects` (boolean, optional): Whether to automatically follow HTTP redirects. Defaults to true.

**Returns:**
Dictionary containing:
- `success` (boolean): Whether the request succeeded
- `status_code` (integer): HTTP status code (if successful)
- `headers` (object): Response headers
- `body` (string/object): Response body (parsed as JSON if possible, otherwise text)
- `error` (string): Error message (if unsuccessful)
- `troubleshooting` (string): Debugging guidance (if unsuccessful)

## Configuration

### Environment Variables
| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `ALLOWED_DOMAINS` | No | `localhost` | Comma-separated list of domains that Claude can access. Supports wildcards (e.g., `localhost,*.hvs,192.168.1.*`) |
| `LOG_LEVEL` | No | `INFO` | Logging verbosity. One of: DEBUG, INFO, WARNING, ERROR |

### Security Model
The server uses an allowlist approach for security:
1. Configure `ALLOWED_DOMAINS` to specify which domains/IPs Claude can access
2. All requests are validated against this allowlist before execution
3. Wildcards (`*`) are supported for pattern matching (e.g., `*.internal.com`, `192.168.*`)
4. Requests to non-allowlisted domains are rejected with a clear error message

## Usage Examples

### Basic Usage
Ask Claude:
```
"Check if my local development server at http://localhost:3000 is running"
```

Claude will use the `fetch` tool to make a GET request and report the status.

### Interacting with Internal APIs
```
"Get the user list from our internal API at https://api.internal.hvs/users"
```

### Testing with Custom Headers
```
"Make a POST request to http://localhost:8080/api/data with Authorization header 'Bearer my-token' and JSON body {\"name\": \"test\"}"
```

### Development Workflow
```
"Test all endpoints of my REST API running at http://localhost:4000 - check /health, /api/users, and /api/posts"
```

## Use Cases

**Local Development**: Test and debug web applications while they're running on localhost
**Internal Infrastructure**: Access company internal APIs and services with custom DNS
**Private Networks**: Query services on private IP ranges (e.g., home lab, company intranet)
**DevOps**: Health check internal services, query metrics endpoints, test API integrations
**Security Testing**: Interact with services behind firewalls or on isolated networks

## Links

- [Source Repository](https://github.com/aj-geddes/local-http-mcp)
- [Issue Tracker](https://github.com/aj-geddes/local-http-mcp/issues)

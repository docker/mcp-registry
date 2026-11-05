# Linked API

Linked API connects AI assistants to your LinkedIn account through a cloud browser. Tools cover messaging, connections, posts, search, and profile data, including Sales Navigator operations.

A Linked API subscription and a connected LinkedIn account are required. Copy the Linked API token and the account identification token from [your dashboard](https://app.linkedapi.io), then configure both secrets in Docker MCP Toolkit.

The catalog runs the existing Docker image in stdio mode with `node dist/index.js`.

- [Installation and token setup](https://linkedapi.io/mcp/installation)
- [Source code and issue tracker](https://github.com/Linked-API/linkedapi-mcp)

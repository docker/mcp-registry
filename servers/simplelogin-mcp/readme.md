# simplelogin-mcp

Independent, self-hostable MCP server for SimpleLogin email-alias workflows. This is not an official
SimpleLogin or Proton AG product and is not affiliated with or endorsed by either company.

Configuration requires a SimpleLogin API key from Settings -> API Keys. The Docker MCP Registry
entry runs the container in stdio mode with `TRANSPORT=stdio` and passes the key through
`SL_API_KEY`.

Full documentation: https://simplelogin-mcp.com/

The catalog uses the project's versioned GHCR image. The same stable image is also available from
[Docker Hub](https://hub.docker.com/r/enthouan/simplelogin-mcp), mirrored from the verified GHCR
digest. These are maintainer-provided images; Docker Hub availability does not imply Docker-managed
building, signing, or MCP Catalog acceptance.

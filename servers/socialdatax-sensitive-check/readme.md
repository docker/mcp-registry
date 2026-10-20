# SocialDataX Sensitive Text Check

Check text for sensitive terms and moderation risks, and query account points balance.

## Setup

1. Get your API Key and credits from [SocialDataX](https://socialdatax.com/ai?from=docker).
2. Set `socialdatax-sensitive-check.api_key` in Docker MCP Toolkit to your own API Key (without the `Bearer` prefix).
3. Enable this server and connect your MCP client to the Docker MCP Gateway.

The gateway connects to `https://mcp.socialdatax.com/sensitive-check/mcp` over Streamable HTTP and sends the key in the Authorization header. Business tool calls require sufficient SocialDataX credits; tool discovery does not establish access to paid operations.

Sensitive-text results are screening aids and do not guarantee platform acceptance. Text sent to the checking tool is transmitted to the hosted service.

Public [connection documentation](https://github.com/DevinChen2014/socialdatax-sensitive-check-mcp) describes available tools and client configuration. The business service is privately hosted.

Website and API Key: https://socialdatax.com/ai?from=docker

# SocialDataX Weibo

Search Weibo posts and users, read comments and hot topics, and get video transcripts.

## Setup

1. Get your API Key and credits from [SocialDataX](https://socialdatax.com/ai?from=docker).
2. Set `socialdatax-weibo.api_key` in Docker MCP Toolkit to your own API Key (without the `Bearer` prefix).
3. Enable this server and connect your MCP client to the Docker MCP Gateway.

The gateway connects to `https://mcp.socialdatax.com/weibo/mcp` over Streamable HTTP and sends the key in the Authorization header. Business tool calls require sufficient SocialDataX credits; tool discovery does not establish access to paid operations.

Public [connection documentation](https://github.com/DevinChen2014/weibo-mcp) describes available tools and client configuration. The business service is privately hosted.

Website and API Key: https://socialdatax.com/ai?from=docker

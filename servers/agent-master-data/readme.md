# Agent Master Data MCP

[Documentation](https://api.agentmasterdata.com/docs) · [Website](https://agentmasterdata.com)

Connect directly to `https://api.agentmasterdata.com/mcp` using Streamable HTTP.
No OAuth or API key is required for initialization, tool discovery, or the
`get_master_data_schema` tool. Inspect current schemas and availability first.

Data-returning operations use x402 payments. This entry does not install a
payment signer, supply wallet credentials, or claim Docker MCP Toolkit can
complete paid calls. Clients without an x402 payment flow can use free discovery;
paid delivery requires a separately configured compatible client. Published
schemas are not a claim that every entity dataset is loaded.

[Current pricing](https://api.agentmasterdata.com/v1/pricing) ·
[Current readiness](https://api.agentmasterdata.com/readyz)

Support: support@agentmasterdata.com

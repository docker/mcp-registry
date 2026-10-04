# BlockVectra MCP Server

Remote MCP server for BlockVectra: multi-chain EVM JSON-RPC, an indexed Data API, documentation, Compute Unit pricing and live status. An API key is optional.

## Features

- Documentation, supported chains, live prices and network status tools work without a key.
- `rpc_call` runs read-only JSON-RPC methods; `send_raw_transaction` broadcasts signed transactions.
- `data_api_get`, `get_account` and `get_deposit_address` use an API key sent in the optional `x-api-key` header.

## More Information

- Website: https://blockvectra.com/en/agents/
- Source: https://github.com/blockvectra/mcp
- Endpoint: https://docs.blockvectra.com/mcp (Streamable HTTP)

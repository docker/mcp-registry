# Local Model Suitability MCP

AI-powered routing advisor for AI agents. Determines whether a task can run on a local model (Ollama, LM Studio, Llama) instead of a cloud API, saving inference costs on every eligible call.

**Homepage:** https://kordagencies.com
**npm:** [local-model-suitability-mcp](https://www.npmjs.com/package/local-model-suitability-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://local-model-suitability-mcp-production.up.railway.app/mcp`

Free tier: 20 calls/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Check whether a summarisation task can run locally before sending to a cloud model:

```bash
curl -X POST https://local-model-suitability-mcp-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "check_local_viability",
      "arguments": {
        "task": "Summarise a 500-word news article into 3 bullet points",
        "quality_threshold": "PRODUCTION",
        "data_sensitivity": "PUBLIC"
      }
    }
  }'
```

Returns: `verdict` (LOCAL/CLOUD/EITHER), `recommended_model`, `confidence`, `reasoning`, and `agent_action`.

## Tools

- **check_local_viability** — Routes tasks to local or cloud models based on complexity and sensitivity

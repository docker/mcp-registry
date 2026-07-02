# HS Code Classifier MCP

HS code classification and validation for AI agents. Classifies product descriptions against official government tariff schedules from USITC (US), Singapore Customs, CBSA (Canada), Australia Border Force, and WCO. Validates supplier-provided codes for misclassification.

**Homepage:** https://kordagencies.com
**npm:** [hs-code-classifier-mcp](https://www.npmjs.com/package/hs-code-classifier-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://hs-code-classifier-mcp-server-production.up.railway.app/mcp`

Free tier: 10 calls/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Classify a product before generating a customs declaration:

```bash
curl -X POST https://hs-code-classifier-mcp-server-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "hs_classify_product",
      "arguments": {
        "product_description": "solid oak dining chair with upholstered seat for dining room use",
        "country": "US"
      }
    }
  }'
```

Returns: `hs_code`, `confidence` (HIGH/MEDIUM/LOW/AMBIGUOUS), `official_heading`, `source`, `tariff_version`, and `agent_action`.

## Tools

- **hs_classify_product** — Classify product description to HS code (free tier: 10/month)
- **hs_validate_code** — Validate supplier-provided HS code for misclassification (paid tier only)

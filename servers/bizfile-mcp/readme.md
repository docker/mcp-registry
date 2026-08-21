# Bizfile Counterparty Validator MCP

Real-time company verification, KYC, and sanctions screening for AI agents. Validates counterparties against UK Companies House, Singapore ACRA, US SEC EDGAR, and 328 global sanctions lists before payments, contracts, or supplier onboarding.

**Homepage:** https://kordagencies.com
**npm:** [bizfile-mcp](https://www.npmjs.com/package/bizfile-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://bizfile-mcp-production.up.railway.app/mcp`

Free tier: 20 calls/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Validate a counterparty before onboarding:

```bash
curl -X POST https://bizfile-mcp-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "validate_counterparty",
      "arguments": {
        "company_name": "Acme Limited",
        "jurisdiction": "gb"
      }
    }
  }'
```

Returns: `agent_action` (PROCEED/ENHANCED_DUE_DILIGENCE/BLOCK), `risk_score`, `kyc_confidence`, company status, officers, and registered address.

## Tools

- **validate_counterparty** — Full registry lookup + AI risk analysis + officer details
- **screen_counterparty** — Sanctions screening against 328 global lists
- **validate_counterparty_lite** — Budget-optimized registry-only check (no AI call)

# VAT Validator MCP

VAT number validation and invoice fraud detection for AI agents. Validates EU VIES (27 states), UK HMRC, and Australian ABN in real time. AI-powered fraud risk analysis and invoice comparison included.

**Homepage:** https://kordagencies.com
**npm:** [vat-validator-mcp](https://www.npmjs.com/package/vat-validator-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://vat-validator-mcp-production.up.railway.app/mcp`

Free tier: 20 calls/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Validate a VAT number before paying an invoice:

```bash
curl -X POST https://vat-validator-mcp-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "validate_vat",
      "arguments": {
        "vat_number": "GB220430231"
      }
    }
  }'
```

Returns: `valid` (true/false), registered `company_name`, `address`, and VAT status.

## Tools

- **validate_vat** — Validates EU, UK, or Australian VAT numbers
- **validate_uk_vat** — UK HMRC validation with consultation number
- **get_vat_rates** — Current rates for 27 EU states, UK, and Australia
- **batch_validate** — Up to 10 VAT numbers in one call
- **analyse_vat_risk** — AI fraud risk analysis on a validated number
- **compare_invoice_details** — Invoice details vs registry comparison

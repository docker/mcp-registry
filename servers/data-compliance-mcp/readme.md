# Data Compliance Classifier MCP

AI-powered data safety classification for GDPR, HIPAA, PCI-DSS, and CCPA compliance. Call before your agent stores, transmits, logs, or passes any data payload to another system.

**Homepage:** https://kordagencies.com
**npm:** [data-compliance-mcp](https://www.npmjs.com/package/data-compliance-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://data-compliance-mcp-production.up.railway.app/mcp`

Free tier: 20 classifications/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Classify a data payload before writing it to a database:

```bash
curl -X POST https://data-compliance-mcp-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "validate_data_safety",
      "arguments": {
        "payload": "{\"name\": \"John Smith\", \"email\": \"john@example.com\", \"dob\": \"1985-03-15\"}",
        "context": "write to customer database"
      }
    }
  }'
```

Returns: `verdict` (SAFE_TO_PROCESS/REDACT_BEFORE_PASSING/DO_NOT_STORE/ESCALATE), `regulatory_frameworks` triggered, `data_categories` detected, and `agent_action`.

## Tools

- **validate_data_safety** — Full AI classification with jurisdiction detection
- **get_safety_report** — Batch classification (up to 50 payloads) or compliance audit report
- **validate_data_safety_lite** — Pattern-only pre-screen at ~70% lower token cost

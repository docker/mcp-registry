# URL Safety Validator MCP

AI-powered URL safety checker for AI agents. Validates URLs before your agent fetches, clicks, or processes them using Google Web Risk, Google Safe Browsing, RDAP domain age, SSL verification, and AI contextual analysis.

**Homepage:** https://kordagencies.com
**npm:** [url-safety-validator-mcp](https://www.npmjs.com/package/url-safety-validator-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://url-safety-validator-mcp-production.up.railway.app/mcp`

Free tier: 10 calls/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Check a URL from an email before visiting it:

```bash
curl -X POST https://url-safety-validator-mcp-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "check_url",
      "arguments": {
        "url": "https://example.com/document"
      }
    }
  }'
```

Returns: `verdict` (SAFE/SUSPICIOUS/DANGEROUS), `trust_score` (0-100), `threat_categories` (phishing/malware/typosquatting), `ssl_valid`, `domain_age_days`, and `agent_action` (ALLOW/FLAG_AND_PROCEED/BLOCK).

## Tools

- **check_url** — Full URL safety analysis with SAFE/SUSPICIOUS/DANGEROUS verdict

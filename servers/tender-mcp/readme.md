# Tender MCP - Government Contract Search

AI-powered government tender search and bid scoring for AI agents. Searches UK Contracts Finder, EU TED, and US SAM.gov simultaneously. Returns scored opportunities with BID/INVESTIGATE/SKIP recommendations.

**Homepage:** https://kordagencies.com
**npm:** [tender-mcp](https://www.npmjs.com/package/tender-mcp)

## Usage

Connect to this MCP server via Docker MCP toolkit or any MCP client supporting streamable-http transport.

**Endpoint:** `https://tender-mcp-production.up.railway.app/mcp`

Free tier: 10 searches/month, no API key required. For paid access, obtain a key at kordagencies.com.

## Example

Search for cybersecurity tenders across UK, EU, and US markets:

```bash
curl -X POST https://tender-mcp-production.up.railway.app/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json, text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "search_tenders",
      "arguments": {
        "keyword": "cybersecurity",
        "company_profile": "UK-based cybersecurity firm specialising in penetration testing and SOC services",
        "sources": ["uk", "eu", "us"],
        "limit": 10
      }
    }
  }'
```

Returns: tenders with `fit_score` (0-100), `agent_action` (BID/INVESTIGATE/SKIP), deadline, estimated value, and key requirements.

## Tools

- **search_tenders** — Multi-market search with AI fit scoring
- **get_tender_intelligence** — Daily digest and award history intelligence

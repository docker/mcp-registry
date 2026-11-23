# CGC Airport Transfers

Private fixed-price airport transfers on the Gold Coast, Brisbane, Byron Bay and Northern NSW, Australia.

- **MCP endpoint (streamable HTTP):** https://cgctransfers.com.au/wp-json/cgc-gateway/v1/mcp
- **Auth:** none for quotes, availability and business info (rate-limited); `request_booking` needs an agent key, issued on request
- **Tools:** `get_service_info`, `get_quote` (fixed prices plus a pre-filled booking link), `check_availability`, `request_booking`
- **Official MCP Registry:** `au.com.cgctransfers/booking`

Docs: https://cgctransfers.com.au/ai-agent-gateway/

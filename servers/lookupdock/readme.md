# LookupDock

Setup and tool documentation: https://lookupdock.com/docs/agents

Use your own LookupDock bearer API key. The canonical remote endpoint is
`https://lookupdock.com/mcp` (Streamable HTTP). Enter the key through the Docker
MCP secret store; never commit it or include it in a public configuration.

The subscription is USD 99/month for 20,000 shared credits. A card is required
for the seven-day, 500-credit trial, which renews unless canceled. A combined
email, phone and IP check uses up to three credits. See current terms at
https://lookupdock.com/pricing.

Live checks cover local email syntax/mail-domain/DNS signals, basic phone
numbering plans and IP network/hosting/Tor signals. They do not check mailbox
existence or current carrier and do not claim VPN detection. US/Canada phone
line type stays unknown. Signals do not prove fraud.

This catalog metadata is contributed under the registry's MIT license.
The hosted LookupDock service is proprietary and is used under its published
terms: https://lookupdock.com/terms. No server implementation source is included.

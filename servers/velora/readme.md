Docs: https://github.com/crossi-dev/velora-mcp

Velora is an AI-native commerce toolkit for Argentine businesses, exposed over MCP. It covers fiscal invoicing (AFIP/ARCA), payments (MercadoPago), product catalog, sales, cash register, and messaging (WhatsApp).

## Endpoints

- **Public (no auth)** — `https://tools.somosvelora.com/api/mcp/public`
  Exposes 9 no-auth fiscal utility tools (CUIT/CUIL/CBU validation, IVA split, AFIP date/QR helpers, invoice-type reference, ARS formatting). This is the endpoint registered here.
- **Full suite (auth)** — `https://tools.somosvelora.com/api/mcp`
  The complete 61-tool suite (invoicing, payments, catalog, sales, caja, messaging) behind authentication.

## Notes

- Transport: streamable HTTP. Clients must send `Accept: application/json, text/event-stream`.
- Status: AFIP invoicing and MercadoPago payments are live in production; credit/debit notes and logistics dispatch are currently sandbox.

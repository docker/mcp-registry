# InvoiceIn MCP Server

Docs: https://invoicein-api.peculiar.systems/docs

The **InvoiceIn MCP Server** gives an agent the receiving side of European e-invoicing. Point it at a file a supplier sent — XRechnung (UBL or CII), EN 16931 UBL 2.1, Peppol BIS Billing 3, CII D16B, ZUGFeRD 1.0/2.x and Factur-X hybrid PDFs, Italy's FatturaPA 1.2 or Poland's KSeF FA(2)/FA(3) — and it comes back as structured data instead of an attachment nobody can read.

## Tools

- `read_invoice` — canonical EN 16931 JSON, the same shape whatever syntax came in, fields named by their BT/BG number
- `validate_invoice` — the official rule sets (EN 16931, XRechnung/KoSIT, Peppol, FatturaPA, KSeF) plus arithmetic cross-checks, each finding with a plain-language fix in English, German, Polish, Italian or French
- `invoice_to_html` — a human-readable rendering, the same A4 layout for every input
- `invoice_to_csv` — flat CSV, one row per line item or per document
- `invoice_to_datev` — a DATEV Buchungsstapel (EXTF 700) for the incoming invoice

## Authentication

Optional. Without a key the demo quota applies: 20 invoices per IP per day. A key is sent as the `X-Api-Key` header and raises that limit.

## Privacy

Nothing is stored. Each request is parsed, validated and rendered in memory and discarded when the response is sent — no document is written to disk, and there is no queue, cache or log of document contents.

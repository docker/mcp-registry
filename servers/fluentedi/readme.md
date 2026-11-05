# FluentEDI

Deterministic tools for AI agents, over a remote streamable-HTTP MCP endpoint. No API key,
no signup, no install.

A language model has no clock, no calculator, no resolver and no parser. It approximates all
four, confidently. FluentEDI answers those questions exactly instead.

## What agents can do with it

**Retail EDI**, where coverage is deepest and no other free service offers anything:

- `edi_parse` — read X12 850 purchase orders, 856 ASNs, 810 invoices, 855 acknowledgments
  and 997 functional acknowledgments into structured JSON
- `edi_build` — generate a valid 856 from a packing manifest, assigning HL identifiers during
  the walk so every level names the ID its parent actually received, with a fixed-width
  106-character ISA and SE01/CTT01/GE01/IEA01 computed from what was emitted
- `edi_validate` — the checks a trading partner runs before rejecting a file: control-number
  mismatches, segment miscounts, orphaned HL parents, missing carton levels
- `edi_acknowledge` — turn `AK5*R*5` with `AK3*HL*4**7` into "segment HL at position 4 is out
  of sequence", and report which documents you sent were never acknowledged at all
- `gs1_checkdigit`, `id_validate` — GS1 mod-10 for SSCC-18 and GTIN, plus IBAN, ISIN, CUSIP,
  LEI, ABA, IMEI, ORCID, VIN and GSTIN

**General determinism**, for the rest of an agent's work:

- Time: current time in any IANA zone, timezone conversion, calendar-aware arithmetic,
  whether an instant falls inside a delivery window and by how much it missed
- Cron: validate an expression, describe it in English, list its next runs in a timezone
- JSON: repair malformed output with the line and column where it broke, JSONPath, structural
  diff, schema inference, RFC 8785 canonicalization with a CIDv1
- Verification: assert an endpoint returned the expected status, headers and JSON values;
  check whether a link still resolves; verify Ed25519, ECDSA, RSA and HMAC signatures
- Safety: scan text or a URL for credentials and personal data, and get back a redacted copy

Every tool is read-only, idempotent and side-effect free, so it is safe to auto-run.

## Connecting

```
https://fluentedi.com/mcp
```

`tools/list` returns 17 tools by default rather than all 42, because clients cap how many
tools they hold across all servers at once. Everything else is reachable through
`tool_search`, `tool_describe` and `tool_call`, or by appending `?tools=all` to the URL.

## Also available over plain HTTP

Every tool is a single stateless GET or POST at `https://fluentedi.com/v1/{tool}`, described
by OpenAPI 3.1 at `/openapi.json` and by an agent-facing reference at `/llms.txt`. Failed
calls return the tool's full parameter schema plus working examples, so a wrong call teaches
the caller how to make the next one succeed.

## Privacy

Stateless. Request and response bodies are processed in memory and discarded; nothing sent is
stored, logged as content, or used for training. `dns_lookup`, `http_check` and `http_assert`
reach the network on the caller's behalf and say so in their descriptions; every other tool is
computed at the edge and touches nothing external.

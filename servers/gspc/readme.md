# Council of AI GSPC — measurement over MCP

Independent AI governance / model safety / EU AI Act / agent verification / MCP governance server.

Quote live `totals.public_count` from `GET https://councilof.ai/api/gspc` — never freeze the slot/measured pair in this file. Every result is an Ed25519-signed, offline-verifiable measurement card. UNMEASURED is reported, never hidden. Measurement, not certification.

HTTP (7 tools, streamable-http, no auth): `POST https://councilof.ai/mcp`
- `board_totals` — live slot count and measured count as two labelled numbers
- `get_axis` — one axis row (MEASURED or UNMEASURED, never a zero)
- `verify_card` — three-state VALID / INVALID / UNCHECKABLE under `did:web:csoai.org#card-attestation-1`
- `list_cards` — published index vs card store (two labelled numbers)
- `get_root` — public-root merkle head
- `get_card` — one signed leaf by id
- `verify_inclusion` — prove a leaf is under the root

stdio npm `csoai-gspc-mcp@0.1.0` is four tools; prefer the remote URL until 0.1.1.

- **Verify:** https://councilof.ai/gspc-verify
- **Connect:** https://councilof.ai/connect-gspc
- **Source:** https://github.com/CSOAI-ORG/councilof-ai
- **Official MCP Registry:** `io.github.CSOAI-ORG/gspc`

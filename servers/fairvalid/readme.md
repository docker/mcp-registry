https://fairvalid.com/help/mcp

# FairValid and FairMatch

FairValid integration 1.1.0 exposes twelve tools through the same hosted OAuth
endpoint: six email-verification tools and six FairMatch lead-discovery tools.
The remote catalog uses dynamic tool discovery; `tools.json` remains `[]`.

FairMatch adds `fairvalid_get_lead_filters`, `fairvalid_search_leads`,
`fairvalid_quote_leads`, `fairvalid_unlock_leads`, `fairvalid_get_lead_batch` and
`fairvalid_get_leads`. Filter discovery, search, quotes, batch reads and saved
leads require `leads:read`; unlocking requires `leads:acquire`. Existing grants
do not gain scopes when a catalog entry is updated. Reconnect through native
OAuth only when FairMatch is wanted, selecting a key with the intended workspace,
permissions and lifetime cap.

Use returned selection IDs, quote the exact 1–50 selected leads, then obtain
approval for that selection and a numerical credit maximum before unlocking.
Pass integer `maximumCredits` from 0–200, `confirm: true` and one stable
`idempotency_key`. A quote reserves no credits. Each newly acquired verified
work email costs four credits; already-owned and unavailable contacts cost zero.
Retain the returned batch ID, poll with the returned retry hints, and retrieve
actual completed results. For uncertain retries, preserve the original API-key
identity, ordered selection IDs, credit maximum and idempotency key.

Installation, connection, searching and quoting do not authorize acquisition.
Tool filters do not replace scopes, caps or confirmation. FairMatch does not
send outreach or grant permission to contact a person. Live verification retains
its separate credit policy.

OAuth selects Live keys. Use separate direct Test-key connections and the
documented synthetic fixtures for sandbox checks; never acquire real leads as an
installation test.

Exact contracts and review fixtures:
https://github.com/moonlitnayeem/fairvalid-ai/blob/f5f04f4354a5c2fe075728f0d664b90d8848640d/docs/fairmatch.md

Support: nayeem@fairvalid.com

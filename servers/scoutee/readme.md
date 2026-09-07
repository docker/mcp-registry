Docs: https://scoutee.org/en/api-docs

Scoutee aggregates the public procurement notices published every day by official journals and buyer
platforms across Europe (TED and the national portals of 30 countries) and North America (SAM.gov,
CanadaBuys, SEAO), deduplicated into one row per notice.

Tools (both read-only):

- `search_tenders` — full-text search over title, buyer and description, with filters for keywords,
  countries, source portal, estimated value range, sort order, closed notices and freshness. Returns a
  page of notices, the total, a per-country facet and the caller's remaining hourly quota.
- `get_tender` — one notice in full by its Scoutee id: buyer, description, CPV codes, estimated value,
  publication and deadline dates, documents link and the source URL where a bid is filed.

Authentication: a workspace API key sent as `X-API-Key: sct_...` (`Authorization: Bearer` is also
accepted). Create a key at https://scoutee.org on the Standard plan, workspace page, section "API".
Quota: 10,000 searches per hour per key. Pricing: https://scoutee.org/pricing

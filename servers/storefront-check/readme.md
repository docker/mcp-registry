# Storefront Check

Audit a public ecommerce storefront from the outside. No credentials and no
cooperation from the store, so it works on a competitor or a prospect as
readily as your own site.

Documentation: https://tools.profitpather.com/store

- `detect_tracking_tags` - every ad and analytics tag visible in the HTML, with ids
- `check_agent_readiness` - whether AI crawlers may read the store, and whether the catalogue is machine readable
- `check_structured_data` - JSON-LD types, parse failures, and commerce markup
- `audit_product_feed` - Shopify catalogue health for Merchant Center and Meta feeds

The server distinguishes **present, absent and not observable** and never
conflates them. A Shopify Custom Pixel runs in a sandboxed iframe and cannot be
seen from outside at all, so an empty tag list is not evidence a store is
untagged; every tool carries a `not_observable` list saying what it could not
see. Nothing is executed, so only server-rendered HTML is read.

The check worth knowing about: each Google measurement id is verified against
Google Tag Manager, which answers `404` for an id it does not recognise. A dead
id means `gtag` never initialises, so a store looks correctly tagged while
recording nothing, and a campaign can spend for months at zero recorded
conversions. A consent or cookie app carrying a stale id from an old property is
the usual cause, and this failure shows up nowhere else.

Docs: https://app.forthclear.io/developers/mcp

# ForthClear Liquidation Marketplace

Remote MCP server for [ForthClear](https://forthclear.io), a B2B marketplace for surplus, overstock and customer-return inventory. Buyers' agents can search live lots, read listing details, request bulk quotes from sellers and check out. Sellers' agents can list lots and answer quotes.

- Endpoint: `https://app.forthclear.io/mcp` (streamable HTTP)
- Authentication: none for the public tools; OAuth 2.1 with dynamic client registration for buyer and seller tools.
- Payments: Stripe. ForthClear runs the marketplace; the listing's seller sells the goods.
- Publisher: Forthsuite (Hylke Reitsma), info@forthclear.io
- Also published in the official MCP registry as `io.forthclear/liquidation`.

## Tools that need no sign-in

- `search_inventory`: keyword search of published, in-stock lots.
- `search_inventory_semantic`: describe what you want in everyday language.
- `get_product_details`: one listing with price, quantity, condition and photo.
- `search_suppliers`: published supplier directory profiles.
- `get_pricing_recommendation`: suggested liquidation unit price from MSRP, condition and lot size.
- `request_quote`: ask a seller to price a bulk order; pass `buyer_email`, where the seller replies.
- `get_quote_status`: check a quote with its status token.
- `preview_checkout`: seller of record, price, shipping and total for an instant-checkout listing. Does not charge.

## Tools that need sign-in (OAuth)

Buyer: `quick_buy`, `confirm_purchase`, `accept_quote_and_pay`, `list_my_quotes`, `list_my_orders`, `get_order_status`, `mark_order_received`, `confirm_order`, `cancel_order`, `open_dispute`, `get_dispute_status`, `subscribe_to_query`.

Seller: `create_listing`, `get_quote_requests`, `respond_to_quote`, `bulk_update_listings`, `bulk_mark_shipped`, `bulk_export_orders`, `list_seller_orders`, `seller_inbox_summarise`.

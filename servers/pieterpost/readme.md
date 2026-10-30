# PieterPost

PieterPost prints and posts physical letters, postcards, and greeting cards from AI assistants.

Connect to https://pieterpost.com/mcp/public/ with Streamable HTTP. No PieterPost account, API key, or OAuth sign-in is required.

Use the tools to quote postage, upload files, prepare unpaid drafts, create editable letter previews or hosted checkout links, and check payment, fulfillment, and available delivery updates. The user reviews and pays on PieterPost before printing and posting. For demonstrations, use mode=test. Test orders never send physical mail.

Design greeting cards with `list_card_options`, `create_card_design`, `get_card_design`, and `update_card_design`. Use `upload_asset` for a supplied photo, inspect character and print warnings, and give the user the private editor link to refine and order. Drafts return PNG previews and do not order or print anything. The server currently exposes thirteen guest tools through dynamic discovery; `tools.json` stays empty for this remote server.

Documentation: https://pieterpost.com/api/docs/mcp/

Machine-readable documentation: https://pieterpost.com/api/docs/mcp.md

Support: pieter@pieterpost.com

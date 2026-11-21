# Jobs MCP Server

Documentation: https://apify.com/conserving_celerytop/jobs-mcp-server

A remote MCP server (Streamable HTTP) with two read-only tools:

- `get_company_jobs`: open jobs of 1 to 25 companies you name, read live from their career pages.
- `search_tech_jobs`: search of tech jobs open today.

Endpoint: https://conserving-celerytop--jobs-mcp-server.apify.actor/mcp

Authentication: your own Apify API token, sent as `Authorization: Bearer <token>`. Get it at https://console.apify.com/settings/integrations. Runs happen in your own Apify account.

Price: $0.005 per successful tool call, plus the per-result charges of the Apify Actors it calls. All billed to your Apify account.

Source code: the server code is not public.

Disclosure: I built this server and publish it on the Apify Store.

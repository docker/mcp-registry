Docs: https://askwatch.ai/free-tools/google-search-console-mcp

Connect a Google account once at the page above and AskWatch issues a personal key. There is no Google Cloud project to create and no credentials file to keep.

The key goes in the `Authorization: Bearer` header, or as a path segment (`https://mcp.askwatch.ai/gsc/<key>`) for clients that prefer that.

The grant asks for `webmasters.readonly`, so none of the eleven tools can submit a sitemap, request indexing, remove a URL or change a property.

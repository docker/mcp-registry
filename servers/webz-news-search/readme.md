# Webz.io News Search

Search global news in natural language and retrieve articles with source links, publication dates, and available metadata.

- [Documentation](https://docs.webz.io/docs/webz/news-search-api-mcp)
- Hosted endpoint: `https://news-search-mcp.webz.io/mcp` (Streamable HTTP)
- Tool: `news_search_by_webz`; tool definitions are discovered after authorization.

## Access and authorization

An active Webz.io account with News Search API access and available credits. Coverage, history, usage limits, and charges depend on the connected account and product entitlement.

Enable this entry in Docker MCP Toolkit and authorize it through Docker's OAuth flow. On the Webz consent page, enter the API token from your Webz dashboard and approve the connection. Webz publishes OAuth discovery and dynamic client registration endpoints with the `mcp` scope and PKCE support. The catalog's secret and environment names are symbolic OAuth storage slots; no credential is included in these files.

## Usage

Use natural-language `query` and the discovered date, source, language, and other filters. `k` limits returned results; it does not establish the total matching corpus. Returned content and metadata depend on indexed coverage. Use available source links and publication dates when citing records.

## Privacy, support, and terms

Queries are sent to Webz.io, and returned material can contain third-party content. See the [privacy policy](https://webz.io/privacy/) and [service terms](https://webz.io/tos/). For account access, service questions, or security reports, contact [support@webz.io](mailto:support@webz.io).

This catalog entry's metadata is contributed under the registry's MIT license. The hosted Webz.io service and third-party content remain subject to their applicable terms and licenses.

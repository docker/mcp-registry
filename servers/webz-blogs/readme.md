# Webz.io Blogs

Search indexed blog posts with Boolean queries and filters, returning post records with source links and available metadata.

- [Documentation](https://docs.webz.io/docs/webz/blogs-api-mcp)
- Hosted endpoint: `https://blogs-mcp.webz.io/mcp` (Streamable HTTP)
- Tool: `search_webz_blogs`; tool definitions are discovered after authorization.

## Access and authorization

An active Webz.io account with Blogs API access and available credits. Coverage, history, usage limits, and charges depend on the connected account and product entitlement.

Enable this entry in Docker MCP Toolkit and authorize it through Docker's OAuth flow. On the Webz consent page, enter the API token from your Webz dashboard and approve the connection. Webz publishes OAuth discovery and dynamic client registration endpoints with the `mcp` scope and PKCE support. The catalog's secret and environment names are symbolic OAuth storage slots; no credential is included in these files.

## Usage

Use Webz Boolean `q` and the discovered filters to retrieve matching records. Set a time window and result size appropriate to the task. A returned page is not necessarily the complete matching corpus. Returned content and metadata depend on indexed coverage. Use available source links and publication dates when citing records.

## Privacy, support, and terms

Queries are sent to Webz.io, and returned material can contain third-party content. See the [privacy policy](https://webz.io/privacy/) and [service terms](https://webz.io/tos/). For account access, service questions, or security reports, contact [support@webz.io](mailto:support@webz.io).

This catalog entry's metadata is contributed under the registry's MIT license. The hosted Webz.io service and third-party content remain subject to their applicable terms and licenses.

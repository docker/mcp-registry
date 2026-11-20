Docs: https://mcp.trymimetic.com/docs
Home: https://trymimetic.com

[Mimetic](https://trymimetic.com) is a growth and conversion tool for web and ecommerce sites. Over MCP an
assistant can read a site's audit findings and session recordings, query its
Google Analytics 4, Search Console, Google Ads and raw GA4 BigQuery export,
connect an existing analytics or experimentation platform with the owner's own
credentials, and provision analytics or ads for a site that has none.

Authentication is OAuth 2.1, discovered automatically from the endpoint: the
client reads the protected resource metadata, registers itself dynamically, and
sends the user through a consent screen. There is no API key or personal access
token to configure, which is why this entry declares no secrets.

Scopes are mcp:read for reads and mcp:write for anything that changes
configuration or creates an account.

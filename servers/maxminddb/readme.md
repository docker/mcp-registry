# MaxMindDB

[Documentation](https://github.com/oschwald/maxminddb-mcp#readme)

Query local MaxMind DB files for IP geolocation, autonomous system information,
and network analysis. This is an unofficial integration, not a MaxMind product.

Set your MaxMind account ID and license key in MCP Toolkit. The default editions
are GeoLite2-City and GeoLite2-ASN. You can select other MMDB editions available
to your account. No databases are bundled with the image.

Databases download in the background and remain in the `maxminddb-mcp-data`
Docker volume between sessions. Ask the assistant to call `list_databases` to
check download status before the first lookup.

Example: "Use MaxMindDB to find the country and autonomous system for 8.8.8.8."

IP lookups run locally. Downloads contact MaxMind or a URL supplied to the
`install_database` tool. Database licenses and download access are managed by
the database provider.

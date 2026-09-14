Docs: https://www.microburbs.com.au/api-access

Microburbs is an Australian property data provider. The MCP server exposes the
same endpoints as the REST API — suburb metrics, property valuations, comparable
sales, schools, demographics, risk and zoning — as tools.

Endpoint reference: https://www.microburbs.com.au/developers/api-docs
Per-endpoint pricing: https://www.microburbs.com.au/developers/api-docs/pricing

Authentication is an API key sent as `Authorization: Bearer mib_live_...`. Create
one at https://www.microburbs.com.au/api-access. A free sandbox key, the literal
string `test`, works without an account and covers one sample suburb and one
sample property per state.

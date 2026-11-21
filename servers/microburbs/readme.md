Docs: https://www.microburbs.com.au/developers/api-docs

Microburbs publishes the largest and most comprehensive Australian real estate
and geospatial data. The MCP server exposes the same 173 endpoints as the REST
API — across 14,764 suburbs and every Australian address — as tools.

Four things no other Australian property API publishes: a price forecast for
every street rather than the suburb's, modelled crime for small Census areas
plus crime by type, small-area geography on real ABS mesh blocks rather than hex
grids, and what is actually within 200 m of a property. Everything goes micro —
demographics, income, crime, risk, liveability, growth and tenure are all
available at mesh-block grain, not just a couple of headline fields. Sale and
rent history, comparable sales, schools, zoning, development and valuations are
there too. 2026 Census work covers updated boundaries and modelled current
estimates, ahead of official ABS results in June 2027.

Authentication is an API key sent as `Authorization: Bearer mib_live_...`. Create
one at https://www.microburbs.com.au/api-access. A free sandbox key, the literal
string `test`, works without an account and covers one sample suburb and one
sample property per state.

# InfraNode

Keyless remote MCP server unifying German public-infrastructure open data:
weather, air quality, traffic, public transit (incl. live departures), parking
and roadworks for 84+ German cities. Data is sourced and normalized from official
German providers (DWD, Umweltbundesamt, Mobilithek, GovData and others).

- **Endpoint:** `https://mcp.infranode.dev/mcp` (remote, streamable HTTP, no API key)
- **Tools:** 38, all read-only
- **License:** Apache-2.0
- **Docs:** https://infranode.dev/mcp
- **Source:** https://github.com/street1983nk/infranode

## Connect

```bash
claude mcp add --transport http infranode https://mcp.infranode.dev/mcp
```

No installation, no key, no account. Tools return a canonical JSON envelope with
`data` and `meta` (including per-source status and attribution).

## Example tools

- `list_cities` — all covered cities with slug, state, population, coverage
- `weather` — current DWD observations for a city
- `air_quality` — Umweltbundesamt air quality (PM10/NO2, ...)
- `transit_departures` — live public-transit departures with real-time delay
- `compare` — compare one resource (weather/air) across multiple cities
- `sources` — overview of all data sources with license and attribution

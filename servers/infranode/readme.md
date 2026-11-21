# InfraNode

Keyless remote MCP server unifying German public-infrastructure open data:
weather, air quality, traffic, public transit (incl. live departures),
electricity prices, land values, station departures, parking and roadworks for
84+ German cities. Data is sourced and normalized from official German providers
(DWD, Umweltbundesamt, SMARD, BORIS, Deutsche Bahn, Mobilithek, GovData and
others).

- **Endpoint:** `https://mcp.infranode.dev/mcp` (remote, streamable HTTP, no API key)
- **Tools:** 42, all read-only
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

- `list_cities`: all covered cities with slug, state, population, coverage
- `weather`: current DWD observations for a city
- `air_quality`: Umweltbundesamt air quality (PM10/NO2, ...)
- `power_price`: SMARD day-ahead electricity prices
- `land_values`: official BORIS land values per city
- `station_departures`: live Deutsche Bahn departures at a city's main station
- `transit_departures`: live public-transit departures with real-time delay
- `compare`: compare one resource (weather/air) across multiple cities
- `sources`: overview of all data sources with license and attribution

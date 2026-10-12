# Criora Climate Risk

Criora is a hosted, free MCP server (streamable HTTP, no authentication) that answers climate and disaster risk questions for any place on Earth.

## Tools

- **assess_place** – one answer for a place: this week's forecast risk by day, hazards nearby, headline climate exposures and the risk profile of its country. The best first call for a general question about a location.
- **find_place** – resolve a place name or address to coordinates (OpenStreetMap Nominatim).
- **get_forecast** – this week's forecast and the daily risk class with the hazard behind it (heat, cold, wind, rain, snow, freezing rain).
- **get_hazards_near** – disasters happening now or recently near a place: wildfires (NASA FIRMS), earthquakes, floods, tropical cyclones, volcanic eruptions and droughts (GDACS), US NWS warnings and air quality (OpenAQ).
- **get_climate_profile** – long-term climate exposure: the latest ERA5 year, warming and rainfall change projected to 2050 and 2080 (CMIP6), sea level rise (IPCC AR6) and more.
- **get_country_profile** – a country's overall climate and disaster risk score with climate, nature, social, governance and adaptation dimensions, and about 160 indicators (INFORM, ND-GAIN, World Bank).
- **compare_countries** – several countries side by side on overall risk and each dimension, with ranks.

## Example questions

- "How exposed is Ghana to climate disasters?"
- "What is the weather risk in Lisbon this week?"
- "Are there any wildfires or earthquakes near Tokyo right now?"
- "Compare climate risk for Spain and Portugal."

## More Information

Documentation: https://criora.com/docs/references/api-and-agents

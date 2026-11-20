Docs: https://mubboo.flights/mcp

Mubboo Flights is an all-in-one travel comparison and tickets platform. This
remote MCP helps travellers prepare flight comparisons with less manual setup.

- `find_airports`: look up airport and city codes.
- `prepare_flight_search`: validate trip details and prepare a flight-search link.
- `get_service_info`: inspect available tools, product links, connection details
  and service information.

Connect to `https://mubboo.flights/api/mcp` using Streamable HTTP. No API key,
OAuth sign-in or local server package is required. Tool discovery is provided
by the hosted endpoint.

The MCP does not return live fares, schedules, inventory, fare rules or baggage
allowances. It does not book travel or process payments. Travellers open the
prepared link to compare options and choose a booking provider.

Do not send passenger names, contact details, passport data or payment details.
Use the traveller's explicit departure market, currency, dates and passenger
counts; do not infer the departure market from the destination or language.

Support and security reports: support@mubboo.flights.

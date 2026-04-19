# Kultur.dev — Cultural Intelligence MCP Server

**Kultur.dev** provides cultural intelligence as an MCP tool layer for AI agents. Analyze content for cultural risks across 195 countries, get localization guidance, check cultural sensitivity, compare Hofstede cultural dimensions, retrieve business communication styles, and access holiday calendars.

## Tools

| Tool | Description |
|---|---|
| `analyze_cultural_context` | Analyze text for cultural risks across 195 countries |
| `get_localization_guidance` | Get localization guidance for a target market |
| `check_cultural_sensitivity` | Check for cultural taboos and offensive content |
| `compare_cultures` | Compare Hofstede dimensions between two countries |
| `get_communication_style` | Get business communication style for a country |
| `get_holiday_calendar` | Get cultural/religious holidays by country/month |

## Authentication

Requires a Bearer token (API key). Get your key at [kultur.dev/dashboard](https://kultur.dev/dashboard) after signing up.

## Transport

- **Type:** SSE (Server-Sent Events)
- **Endpoint:** `https://kultur.dev/api/mcp-sse/sse`

## Links

- **Website:** [kultur.dev](https://kultur.dev)
- **Documentation:** [kultur.dev/docs](https://kultur.dev/docs)

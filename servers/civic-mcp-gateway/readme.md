# Civic MCP Gateway

Connecting AI Agents to tools and data via the Civic MCP Gateway gives builders access to guardrails, scoped permissions, audit trails, and revocable access when calling MCP tools. Civic separates the permission layer from the AI agent so they can't get around restrictions.

## Features

- **85+ MCP servers** — Gmail, Google Calendar, PostgreSQL, Slack, and more through a single gateway
- **Guardrails** — block destructive operations, redact PII from responses, enforce rate limits
- **Audit trail** — every tool call logged with agent identity and timestamp
- **Scoped permissions** — grant agents access to specific tools only
- **Revocable access** — revoke a token and the agent loses access immediately

## Setup

1. Sign up at [app.civic.com](https://app.civic.com) and complete onboarding
2. On first connection, complete the OAuth flow in your browser

## Documentation

- Docs: https://docs.civic.com/civic/quickstart
- Source: https://github.com/civicteam/civic-mcp-gateway

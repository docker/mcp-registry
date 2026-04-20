# MCP Billing Gateway

Billing infrastructure for MCP server operators. Add monetization to any MCP server without building payment infrastructure from scratch.

## What it does

Register your MCP server with the billing gateway to get:
- **Billing proxy**: A managed endpoint that handles payment verification before forwarding to your server
- **Pricing plans**: Per-call, flat-rate, or Stripe subscription billing models
- **Revenue tracking**: Usage analytics, revenue dashboards, and Stripe Connect payouts
- **x402 support**: Native x402 micropayment verification for agent-to-agent billing

## Authentication

Obtain an operator API key by calling `register_operator` with no auth header (free signup), then use the returned key as `Authorization: Bearer <key>` for all subsequent calls.

## Source

- SDK: https://github.com/sapph1re/mcp-billing-gateway-sdk
- Docs: https://mcp-billing-gateway-production.up.railway.app/docs

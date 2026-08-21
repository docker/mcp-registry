# Feedback Synthesis MCP

Customer feedback aggregation and synthesis MCP server. Analyzes feedback from GitHub Issues, Hacker News discussions, and App Store reviews to surface actionable insights.

## Tools

| Tool | Description | Cost |
|------|-------------|------|
| `synthesize_feedback` | Full feedback synthesis with themes and sentiment | $0.05/call |
| `get_pain_points` | Ranked list of user pain points | $0.02/call |
| `search_feedback` | Search feedback by keyword/topic | $0.01/call |
| `get_sentiment_trends` | Sentiment trends over time | $0.03/call |

## Authentication

This server uses the [x402 payment protocol](https://x402.org) for pay-per-use access. No API key or subscription is required — each tool call is priced individually and settled via Base network USDC micropayments.

## Source

https://github.com/sapph1re/feedback-synthesis-mcp

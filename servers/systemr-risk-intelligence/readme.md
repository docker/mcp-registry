Docs: https://github.com/System-R-AI/systemr-python

## System R Risk Intelligence

AI-native risk intelligence for trading agents. Provides:

- **G-formula position sizing** — Optimal position size using geometric growth optimization ($0.003/call)
- **Iron Fist risk validation** — Pre-trade risk checks against configurable rules ($0.004/call)
- **System R Score** — Performance evaluation from R-multiples ($0.10-$1.00/call)
- **Pricing** — View current operation pricing (free)

### Getting Started

1. Register an agent API key: `POST https://agents.systemr.ai/v1/agents/register`
2. Connect via MCP: `https://agents.systemr.ai/mcp/sse` with `X-API-Key` header

### Also available as

- PyPI package: `pip install systemr` (stdio transport)
- Official MCP Registry: `io.github.System-R-AI/systemr-risk-intelligence`

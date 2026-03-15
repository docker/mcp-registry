Docs: https://github.com/System-R-AI/systemr-python

## System R Risk Intelligence

Trading & Investment Operating System for AI agents. 48 tools via MCP:

| Category | Count | Examples | Cost |
|----------|-------|----------|------|
| Compound | 2 | pre_trade_gate, assess_trading_system | $0.01-$2.00 |
| Core | 4 | position_sizing, risk_check, evaluate_performance | $0.003-$1.00 |
| Analysis | 18 | drawdown, monte_carlo, kelly, equity_curve, system_r_score | $0.004-$0.008 |
| Intelligence | 11 | detect_regime, detect_patterns, analyze_greeks | $0.004-$0.008 |
| Planning | 4 | options_sizing, futures_sizing, options_plan | $0.004-$0.008 |
| Data | 3 | calculate_pnl, expected_value, compliance | $0.003-$0.004 |
| System | 5 | equity_curve, score_signal, trade_outcome, scanner | $0.002-$0.005 |
| Journal | 1 | record_trade_outcome | $0.003 |

### Getting Started

1. Register (free — $30 USDC credited): `POST https://agents.systemr.ai/v1/agents/register`
2. Connect via MCP: `https://agents.systemr.ai/mcp/sse` with `X-API-Key` header
3. Or use Python SDK: `pip install systemr`

### Links

- [Python SDK](https://pypi.org/project/systemr/) — `pip install systemr`
- [Demo Agent](https://github.com/System-R-AI/demo-trading-agent) — Reference implementation
- [API Docs](https://agents.systemr.ai/v1/tools/list) — Full tool catalog

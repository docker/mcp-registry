# OptionsAhoy

OptionsAhoy is an equity-compensation tax optimizer exposed as a remote MCP server. It computes optimal schedules for ISO/AMT exercise, NSO exercise, RSU sell-vs-hold, QSBS eligibility, single-stock concentration risk, and protective puts/collars against the full federal tax code plus all 50 states and DC. Each tool returns a globally-optimal, deterministic result over the candidate space rather than an approximation.

## Tools

- `amt_iso_optimize` — Multi-year ISO exercise schedule maximizing after-tax Net Final Value, modeling AMT credit recovery and grant timing.
- `nso_calculate` — After-tax NSO exercise payout (federal, state, FICA) comparing sell-at-exercise vs hold-for-LTCG.
- `rsu_sell_vs_hold` — After-tax RSU vest payout, including the 22% supplemental withholding gap, comparing sell-at-vest vs hold.
- `concentration_analyze` — Single-stock concentration risk analysis on an existing position.
- `protective_put_price` — Black-Scholes pricing of a protective put or zero-cost collar.
- `qsbs_check` — Section 1202 Qualified Small Business Stock qualification check.
- `equity_funding_plan` — Multi-year, multi-stack optimizer to fund a cash goal from equity at minimum after-tax cost.

## Usage

Remote streamable-http endpoint, no authentication required (open endpoint, no OAuth, no API key):

```
https://optionsahoy.com/mcp
```

Docs: https://optionsahoy.com/for-agents

## License

MIT

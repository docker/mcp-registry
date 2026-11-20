# FxSocket

Read-only access to MetaTrader 4 and MetaTrader 5 trading accounts hosted by
[FxSocket](https://fxsocket.com): accounts with balance and equity, open
positions and pending orders, closed-trade history with per-account statistics,
broker symbol search and contract specifications, live quotes and OHLC price
history.

The server requires an FxSocket API key, sent as the `X-API-Key` header. Keys
are created in the FxSocket dashboard; without one the server answers `401`.

Docs: https://fxsocket.com/docs/mcp

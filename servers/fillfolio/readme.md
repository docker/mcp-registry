Docs: https://fillfolio.com/mcp

# Fillfolio

Fillfolio is a read-only portfolio and spending tracker. This remote server lets an AI assistant read the signed-in user's stored Fillfolio data: portfolios, holdings (stocks, ETFs, funds, crypto), cash and liabilities, investment activity, bank spending and transactions, and monthly budgets.

The tools cannot trade, move money, edit records or refresh connected accounts.

## Connecting

- Endpoint: `https://fillfolio.com/api/mcp` (Streamable HTTP)
- Sign-in: OAuth with PKCE in the browser, scope `mcp:read`
- Account: an eligible Fillfolio account
- Spending, transactions and budgets are off by default for each connected app. Turn them on for "MCP Gateway - fillfolio" in Fillfolio Settings.
- Revoke access at any time in Fillfolio Settings, under Connected apps.

## Tools

| Tool | What it reads |
| --- | --- |
| `get_portfolio_summary` | Net worth, holdings value, cash, debt and coverage |
| `list_portfolios` | The user's portfolios |
| `list_holdings` | Holdings with values and cost basis where available |
| `list_cash_and_liabilities` | Cash and debt balances, without account numbers |
| `list_activity` | Investment activity for a date range |
| `get_asset` | One asset and the user's holdings of it |
| `get_spending_summary` | Spending and income totals for a date range |
| `list_bank_transactions` | A page of bank transactions for a date range |
| `list_budgets` | Monthly budget limits and actuals |
| `open_fillfolio` | Opens the Fillfolio dashboard on hosts that support MCP Apps |

## Links

- [Website](https://fillfolio.com/)
- [Privacy policy](https://fillfolio.com/privacy)
- [Terms](https://fillfolio.com/terms)
- [Support](https://fillfolio.com/contact)
- [Public MCP contract and examples](https://github.com/fillfolio/portfolio-mcp)

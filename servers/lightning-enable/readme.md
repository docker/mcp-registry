# Lightning Enable MCP

The payment layer for AI agents. Pay Lightning invoices, manage spending budgets, and access L402-protected APIs.

## Features

- **Pay Lightning Invoices** - Send payments to any BOLT11 invoice
- **L402 Auto-Payment** - Automatically pay API access fees
- **Budget Controls** - Set per-request and per-session spending limits
- **Multiple Wallets** - Support for NWC, OpenNode, and Strike

## Quick Start

1. Configure your wallet (NWC, OpenNode, or Strike API key)
2. Set budget limits for safety
3. Start paying invoices or accessing L402 APIs

## Pricing

| Tier | Price | Tools |
|------|-------|-------|
| FREE | $0 | pay_invoice, check_wallet_balance, get_payment_history, get_budget_status, get_all_balances, create_invoice, check_invoice_status, exchange_currency*, get_btc_price* |
| PAID | 6,000 sats one-time | access_l402_resource, pay_l402_challenge |

*Strike wallet only

## Documentation

- [Quick Start Guide](https://docs.lightningenable.com/products/l402-microtransactions/mcp-quickstart)
- [Wallet Setup](https://docs.lightningenable.com/products/l402-microtransactions/mcp-wallet-setup)
- [AI Spending Security](https://docs.lightningenable.com/products/l402-microtransactions/ai-spending-security)

## Links

- **Website**: https://lightningenable.com
- **GitHub**: https://github.com/refined-element/lightning-enable
- **PyPI**: https://pypi.org/project/lightning-enable-mcp
- **NuGet**: https://nuget.org/packages/LightningEnable.Mcp
- **Docker**: https://hub.docker.com/r/refinedelement/lightning-enable-mcp

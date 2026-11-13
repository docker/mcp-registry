Docs: https://github.com/base-buzz/paystage

Read-only payment preflight for agents. Given a proposed x402 payment and a policy, it returns `pay`,
`review` or `do_not_pay` **with the rule that fired** — before a wallet signs.

It holds no keys and cannot sign, which is the point: the agent stays the only signer. Five tools,
including `evaluate_mcp_payment` and `verify_receipt`.

The same endpoint is listed in the official MCP Registry as `io.github.base-buzz/mcp` and speaks
Streamable HTTP at `https://api.paystage.xyz/mcp`. There is also a stdio package for clients that
prefer one: `npx -y -p @paystage/mcp paystage-mcp`.

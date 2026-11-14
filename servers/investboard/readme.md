Docs: https://investboard.com/product/mcp

Install options and companion skills: https://investboard.com/product/ai-toolkit#install

Investboard is a hosted Streamable HTTP service with 24 MCP tools. Two tools,
`analyze_holdings` and `stress_test_holdings`, work on supplied holdings without
sign-in and read or save nothing in an account. Account tools use OAuth and
require the appropriate permissions, enabled sharing and the applicable Free
or Premium plan.

The OAuth secret/env names in server.yaml follow the Docker remote-server
catalog convention; they contain no credentials and do not imply a separate
Investboard API-key or personal-access-token issuance flow. Account access is
obtained through the investor's OAuth consent in a compatible client.

The public https://github.com/Investboard/ai-toolkit repository contains the
Proprietary integration bundle, skills and docs, not the private hosted runtime.
No container image or open-source runtime is claimed by this remote entry.

Checks and stated decisions may be saved. Imports and prepared reviews require
the investor's confirmation in Investboard. Rebalance comparisons execute no
orders and save no plan or check. Historical stress results are hypothetical
replays, not forecasts. No money transfers or trade execution.

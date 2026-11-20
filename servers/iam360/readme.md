# iAM360

Connect your assistant to your wearable measurements, nutrition, workouts, challenges, and journal in iAM360.

## Connect

1. Create an account or sign in at https://app.iam360.ai. Connect WHOOP or Oura in iAM360 if you want wearable data.
2. Add IAM360 in Docker MCP Toolkit and start its OAuth authorization.
3. Sign in to iAM360 in the browser, review the requested access, and choose Allow.
4. Return to your assistant and try a read request, such as "Show my logged meals and nutrition totals for today."

MCP endpoint: https://app.iam360.ai/api/mcp

The server uses Streamable HTTP and OAuth with PKCE. The authorization issuer is https://clerk.iam360.ai. The `fitness` scope permits reading and writing supported fitness records, including deletion; `offline_access` enables refresh tokens when requested by the client. There is no separate read-only scope. Keep your assistant's tool confirmations enabled.

An iAM360 account is required. Some hosted AI features require an iAM360 plan; see the documentation for feature availability. You do not need an API key or a local server.

## Links

- Docs: https://app.iam360.ai/docs/mcp
- Website: https://www.iam360.ai
- Connector repository and support: https://github.com/iam360-ai/iam360-mcp-connector
- Manage connected agents: https://app.iam360.ai/dashboard/settings/agents
- Privacy: https://app.iam360.ai/privacy
- Terms: https://app.iam360.ai/terms

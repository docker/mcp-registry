Docs: https://skysay.ai/docs/mcp

Skysay is a hosted telephony service using Streamable HTTP. No local server image is required.

Connect to `https://api.skysay.ai/mcp` and sign in to Skysay using OAuth authorization code with PKCE. Review the workspace and requested permissions before approving access. Public sign-up is available at https://skysay.ai.

The default OAuth grant is read-only. Calling, SMS, transcripts, number requests, and wallet top-ups require additional consent. For a connector that exposes only five read tools, use `https://api.skysay.ai/mcp/read-only` instead.

Connecting is free. Phone numbers, phone calls, and SMS require purchased Skysay funds. The welcome credit covers browser testing only. See https://skysay.ai/pricing.

Number availability, capabilities, and verification requirements vary by country. 100+ countries covered.

The public connection configuration is MIT-licensed at https://github.com/skysay-ai/skysay-mcp; the hosted service is governed by https://skysay.ai/terms. The public repository does not contain the hosted application's source.

Support: hello@skysay.ai

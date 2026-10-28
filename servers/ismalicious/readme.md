# isMalicious

Source and installation documentation: https://github.com/hexablob/ismalicious-mcp-server
API documentation: https://ismalicious.com/api-docs
Security reports: contact@ismalicious.com (do not include credentials in public issues).

This MIT stdio server runs locally in a container and calls the hosted isMalicious API. Configure both secret fields with a key pair from https://ismalicious.com/app/account. Free account quotas are available; request and scan meters are separate. Inputs supplied to intelligence and scan tools are sent to the API. Scan decisions are advisory until the calling application enforces them.

The pinned source revision is version 0.5.0, independently of newer npm releases. The tools inventory is captured from that revision's `tools/list` response in configured mode. The container's default key values are nonfunctional placeholders for discovery, not test credentials. Real tool calls require a user-owned key pair.

For Docker review without production credentials, run the source's `npm ci`, `npm run typecheck`, `npm test`, `npm run build`, and `node scripts/smoke.mjs`. The smoke test covers initialize, discovery, calls, quota, cache and cancellation against a local stub API; it does not establish production API availability.

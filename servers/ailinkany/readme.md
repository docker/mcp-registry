# AiLinkAny

Public documentation: https://github.com/AiLinkAnyCom/ailinkany

Website and Windows client: https://ailinkany.com/

AiLinkAny is a hosted Streamable HTTP MCP service. Users run the proprietary AiLinkAny client on their own Windows computer, bind it to their own account, and authorize the MCP host through individual account OAuth. The client and cloud implementation are not included in the public documentation repository. No Docker image is required for this remote entry.

Registration, OAuth and device discovery are separate from execution entitlement. Actual computer execution requires an eligible AiLinkAny subscription. File changes, Shell commands and application control can have side effects; review tool calls and use only explicitly authorized devices and paths. With multiple computers, resolve the user's explicit target; do not silently select the most recently online computer.

Use list_devices or connection_diagnostics for read-only discovery. Reuse the original command ID to query long-running task status/results rather than replaying a side-effect command.

OAuth integration notes for catalog review:

- Resource: https://mcp.ailinkany.com/mcp
- Issuer: https://ailinkany.com
- Protected resource metadata: https://mcp.ailinkany.com/.well-known/oauth-protected-resource/mcp
- Authorization server metadata: https://ailinkany.com/.well-known/oauth-authorization-server
- Authorization Code with PKCE S256, dynamic client registration, scope mcp:tools.
- The personal_access_token key in server.yaml follows the catalog's credential-storage naming convention for OAuth results. It does not provide or request a shared static PAT, administrator token or device token.
- Docker's ailinkany OAuth provider onboarding is pending. Do not represent this entry as a completed Docker OAuth integration until it is configured and tested.

The public endpoint and metadata were checked; authenticated service checks on another host are not Docker Toolkit end-to-end acceptance. No reviewer credentials are included in public files. Any necessary review access must be provisioned separately with explicit owner authorization.

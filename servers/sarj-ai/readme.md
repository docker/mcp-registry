Docs: https://github.com/sarj-ai/mcp-integrations#readme

# Sarj AI

Connect to Sarj AI, an Arabic-first enterprise voice and chat platform, using the hosted Streamable HTTP endpoint at `https://platform-api.sarj.ai/api/v1/mcp`.

## Account and authorization

An existing Sarj organization account is required. Complete OAuth using an account role that can create API credentials, and review the permissions on the consent screen. Access remains limited by the approved scopes and the account's organization permissions. Contact your organization administrator or [Sarj support](https://sarj.ai/contact) if authorization is denied.

The Docker remote catalog schema does not declare an OAuth scope restriction. This entry must not be treated as a read-only connection or as a four-tool allowlist. Use the explicitly read-scoped configurations in the [public client repository](https://github.com/sarj-ai/mcp-integrations) when you need that restriction.

## Read workflows

| Workflow | Tool | Required scope |
| --- | --- | --- |
| Find recent calls and filter by status, scenario, or creation date | `listCalls` | `calls:read` |
| Inspect a call's status, duration, and recorded outcome | `getCall` | `calls:read` |
| Read an available conversation transcript | `getCallTranscript` | `calls:read` |
| Review available outcomes, caller experience, and analysis | `getCallAnalysis` | `calls:read` |

Transcript and analysis availability depends on call completion, processing, and the features enabled for the organization. Missing or pending values remain unavailable.

The production endpoint currently advertises eight tools, including separately scoped call-management operations. Starting, cancelling, or rescheduling calls requires `calls:write` and can change live state. Approve only permissions and tool calls you intend to grant. Tools are discovered dynamically; `tools.json` is intentionally empty.

## Links and licensing

- [Client configuration and setup](https://github.com/sarj-ai/mcp-integrations)
- [MCP setup documentation](https://platform-docs.sarj.ai/mcp-server)
- [Support and security reports](https://sarj.ai/contact)
- [Privacy policy](https://sarj.ai/privacy)
- [Terms of service](https://sarj.ai/terms)

The public client configuration and documentation repository is MIT-licensed. It does not contain or license the hosted Sarj service implementation. The hosted service is governed by Sarj's terms and account access requirements. Sarj names and logos are brand assets; the configuration license grants no trademark rights.

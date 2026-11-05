# Firedrill

[Setup and documentation](https://docs.firedrill.run/guides/control-mcp)
· [Agent quickstart](https://firedrill.run/agents)

Firedrill is a simulation platform for AI agents. A coding agent can prepare
scenarios, configure synthetic dependencies, and inspect saved simulations and
evidence on your behalf. Exploration, demonstrations, comparison, and testing
are supported uses; assessments are optional.

Connect to `https://mcp.firedrill.run/mcp` over Streamable HTTP. A Firedrill
account and browser OAuth authorization are required. Approve the intended
organization, then select its project and environment using the available tools.
Tool discovery is dynamic; no static credential belongs in this catalog.
The OAuth secret and environment names in the catalog are Docker's standard
OAuth storage identifiers, not instructions to paste a token into chat.

The MCP connection does not execute a repository on your computer. Your agent
stays in your own process, repository, CI runner, or existing endpoint. For local
execution, run `npx firedrill-run init` from your application's repository and
follow the CLI setup. CLI and MCP authorization are separate; choose the same
intended organization.

Firedrill is free to use. Your application's model and runner costs remain
separate. Firedrill is a closed-source service; this submission contains only
public connector metadata and documentation, not its service implementation.

[Privacy](https://firedrill.run/privacy) · [Terms](https://firedrill.run/terms)
· [Support](mailto:support@firedrill.run)

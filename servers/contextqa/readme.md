# ContextQA

[ContextQA MCP documentation](https://mcp.contextqa.com/docs)

Connect to the hosted Streamable HTTP endpoint at `https://mcp.contextqa.com/mcp` using browser OAuth with your ContextQA account. Tool discovery occurs after authentication; available operations depend on the selected workspace and account permissions.

[Get started with ContextQA](https://ship.contextqa.com/). The REST API token flow is separate from MCP OAuth. Do not paste a REST token in place of the MCP OAuth credential.

Use authorized staging or synthetic targets for validation. Test creation changes the workspace, and test execution can affect the target application and consume execution capacity. Begin with `get_current_workspace`, select the intended test or plan, and obtain approval before execution. Retain the returned execution handle and inspect it after a timeout before considering another execution.

ContextAI voice/agent evaluation uses a separate service and is not included in this catalog entry.

[Support and security reports](https://contextqa.com/contact-us/) · [Privacy](https://contextqa.com/privacy-policy/) · [Terms](https://contextqa.com/terms-and-conditions/)

# ToolRouter

ToolRouter is the OpenRouter for tools. Where OpenRouter gives you every model behind one API, ToolRouter gives every AI agent one MCP connection to a hosted catalog of 250+ specialist tools and 1,300+ skills, maintained and billed centrally.

Your agent connects once, discovers what it needs with `discover`, and runs it with `use_tool`. Free tools work before you claim an account; paid tools are pay per call from $0.005, with no provider accounts to set up for the catalog's built-in tools.

## Features

- **One connection, many tools** – web search, web extraction and scraping, image and video generation, SEO and AEO analysis, finance, property, compliance, document analysis and more.
- **Compact routing surface** – the whole catalog costs an agent roughly 9,000 tokens of context instead of loading every skill as its own tool definition.
- **Live schemas and prices** – `discover` returns current tool schemas and per-call prices at request time, so the catalog can grow without a client update.
- **No setup to start** – connect through the OAuth consent screen and choose "Connect" with no account; the server provisions a free account and returns a claim link. Free tools work before you claim it.

## Configuration

No credentials are required to connect. Free tools run immediately. Paid tools are pay per call and require credits on a claimed account; an API key can be sent as a bearer token to bind calls to an existing account.

## More Information

- Website: https://toolrouter.com
- Documentation: https://toolrouter.com/docs
- Source (MIT): https://github.com/Humanleap/toolrouter-mcp

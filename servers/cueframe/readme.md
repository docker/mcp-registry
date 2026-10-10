<a href="https://cueframe.ai"><img src="https://cueframe.ai/icon.svg" alt="CueFrame" width="80" height="80"></a>

# CueFrame

CueFrame is a hosted video composition, editing, and rendering service for agents.
[Try your first render](https://cueframe.ai/start) · [Connection guide](https://docs.cueframe.ai/docs/setup) · [Templates](https://cueframe.ai/templates)

[![A frame from CueFrame's Yosemite Peregrines template, with a title behind the subject and word-timed captions](https://raw.githubusercontent.com/cueframe-ai/cueframe-mcp/main/assets/readme/yosemite-render.jpg)](https://cueframe.ai/start)

*A frame from the Yosemite Peregrines template. [Watch the video and try the same edit](https://cueframe.ai/start). The project remains editable in CueFrame Studio.*

Connect to `https://api.cueframe.ai/v1/mcp` using Streamable HTTP and authorize
with your CueFrame account. Interactive clients use browser OAuth; headless
clients can send a CueFrame API key in the Authorization Bearer header.

Setup documentation: https://docs.cueframe.ai/docs/setup

Agent installation instructions: https://github.com/cueframe-ai/cueframe-mcp/blob/main/llms-install.md

Public agent surface (Apache-2.0): https://github.com/cueframe-ai/cueframe-mcp
The hosted server implementation is private. The public repository contains
agent skills, client tooling, and registry metadata; it does not host the server.

Cloud video operations use account credits. Call `get_account` before metered
operations and review pricing at https://cueframe.ai/pricing.

Tools are discovered dynamically after authentication. Dedicated reviewer
credentials have not been provided; authenticated Docker OAuth and tool-call
verification remain pending.

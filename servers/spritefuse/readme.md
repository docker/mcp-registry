# SpriteFuse

[Documentation and connection setup](https://spritefuse.com/mcp)

SpriteFuse provides a hosted Streamable HTTP MCP endpoint at
`https://spritefuse.com/api/mcp`. Sign in to your SpriteFuse account with OAuth
when prompted by a compatible MCP client. No local server installation is needed.

The server exposes 11 tools for capabilities and account information, character
and animation generation, sprite revision, job inspection and cancellation,
owned projects and assets, private reference uploads, frame timing, and exports.
Exports include PNG frames, sprite sheets, GIF previews, and metadata.

Generation uses the authenticated account's SpriteFuse Credits. Each generation
request requires a maximum credit limit; asynchronous jobs are subject to account
and daily limits. New accounts receive trial credits for personal, noncommercial
evaluation. See the documentation for current trial terms and generation limits.

Start with `get_capabilities` and `get_account` to review available operations and
account limits. Generation creates account jobs and can consume credits, so ask
for the user's generation intent and maximum credit budget before creating a job.

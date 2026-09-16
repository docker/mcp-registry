# chat-recall

One searchable memory across every AI coding tool you use.

chat-recall indexes the session transcripts your AI coding tools already write to
disk — Claude Code, Codex, Gemini CLI, OpenCode, Antigravity and Cursor — into a
single searchable history, and exposes it back through MCP so the agent can search
its own past work: resume where a previous session stopped, recall a decision and
why it was made, or find which sessions touched a file.

Secrets are redacted on your own machine before anything is uploaded.

- Documentation: https://chatrecall.dev/mcp/
- How it works: https://chatrecall.dev/how-it-works/
- Self-hosting (free for one person): https://chatrecall.dev/self-hosting/
- Privacy policy: https://chatrecall.dev/privacy/

## Authentication

OAuth 2.1 with dynamic client registration and PKCE. Connecting from a client
takes one browser approval; there is no token to paste.

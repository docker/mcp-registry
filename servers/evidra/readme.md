# Evidra

Flight recorder for AI agents that touch infrastructure.

## Overview

Evidra records what an agent intended to do, what it actually did, and what it deliberately refused to do. It produces an append-only evidence chain around infrastructure mutations and returns reliability context through MCP.

This Docker MCP Registry entry runs the `evidra-mcp` stdio server for local-first usage in Docker Desktop and the MCP Toolkit.

## MCP tools

- `prescribe` records intent before an infrastructure mutation
- `report` records the terminal outcome or deliberate refusal
- `get_event` looks up a previously recorded evidence event

## Runtime behavior

- starts in local-first mode
- sets `EVIDRA_SIGNING_MODE=optional` so the image works without an external signing key
- can forward evidence to a self-hosted Evidra API if the client provides `EVIDRA_URL` and `EVIDRA_API_KEY`

## Links

- [GitHub Repository](https://github.com/vitas/evidra)
- [MCP Setup Guide](https://github.com/vitas/evidra/blob/main/docs/guides/mcp-setup.md)
- [MCP Registry Publication Guide](https://github.com/vitas/evidra/blob/main/docs/guides/mcp-registry-publication.md)
- [Landing Page](https://evidra.samebits.com)

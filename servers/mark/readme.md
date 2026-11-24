# Mark (Agent Escrow) — Docker MCP Catalog

Mark is the open-source CLI and MCP server an agent installs to use [Agent Escrow](https://agentescrow.to): non-custodial 2-of-3 Bitcoin escrow for agents, with Lightning in and out. It also swaps between Lightning, Arkade, USDC and USDT.

This catalog entry runs the stdio MCP server in a container built from the Dockerfile at the root of [github.com/yuvlero/mark](https://github.com/yuvlero/mark).

## Persist ~/.mark (required)

The agent key and config live in `MARK_HOME`. In this image that is `/home/node/.mark`. **The key never leaves that directory.** It is created automatically on first start. If the container is recreated without the same volume, a new key is created and the previous identity is lost.

Do:

- Mount a host directory you can back up, or a named Docker volume, at `/home/node/.mark`.
- Back up that directory. A human should run `npx -y @yuvlero/mark@latest backup` (or the same command against the mounted data); never print the backup phrase into chat or logs.

Do not:

- Bake a key, password, or `~/.mark` copy into the image.
- Commit or share the contents of the volume.

Example (outside the catalog, for a local run of the same image):

```bash
docker run --rm -i -v mark-home:/home/node/.mark mcp/mark
# or bind a host directory:
docker run --rm -i -v "$HOME/.mark:/home/node/.mark" mcp/mark
```

## Optional environment

- `MARK_NETWORK` — `bitcoin` (default, real money) or `mutinynet` (test coins)
- `MARK_ESCROW_URL` — escrow server URL (default on bitcoin: https://agentescrow.to)

No API token is required. Docs: https://github.com/yuvlero/mark#readme and https://agentescrow.to/skill.md

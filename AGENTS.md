# AGENTS.md — mcp-registry

Docker's curated catalog of MCP servers, plus a Go CLI toolset (`cmd/*`) to
create, build, validate, and import MCP server definitions. Server definitions
live under `servers/<name>/server.yaml` (+ optional `tools.json`). Not a
long-running service.

## Cursor Cloud specific instructions

### Toolchain

- **Go 1.24+ is required** (root `go.mod` declares `go 1.24.0`; the
  `agents/security-reviewer` submodule needs 1.25.x). The base Cloud image may
  ship an older Go, so Go 1.24 is installed to `/usr/local/go` during environment
  setup and added to `PATH` via the agent's `~/.bashrc`
  (`export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin`). `go.mod`'s `toolchain`
  directive also auto-fetches the pinned toolchain on first use.
- Module deps: `go mod download` (idempotent; cached in the module cache).

### Build / test / run

- **Build:** `go build ./...`
- **Test:** `go test ./...` (CONTRIBUTING uses `task unittest`; `task` is not
  installed in the VM, so call `go` directly). Tests pass; most packages have no
  test files.
- **Generate a catalog tile (offline):** `go run ./cmd/catalog <server-name>`
  (e.g. `time`) → writes to gitignored `catalogs/<name>/catalog.yaml`.
- Commands that need extra infra (NOT available by default): `cmd/validate`
  (hits the GitHub API for license + `npx prettier`), `cmd/build` / `task import`
  (need Docker + the `docker mcp` plugin), and `task` (taskfile.dev).

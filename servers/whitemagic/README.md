# WhiteMagic local MCP server

The source pin is `0aea631425795eca6e116a875dd5e67abee5ab61` (`v11.2.1`, Linux `amd64`/`arm64` via the root `Dockerfile`).
The entry describes the local stdio server (`wm serve`), not the separate hosted MCP service.

`tools.json` matches the `v11.2.1` default `cyberbrain` profile `tools/list` output (19 tools, including the `wm` dispatcher and `memory_*`, `session_*`, `galaxy_*`, `mesh_sync`, and `sutra_*` surfaces). Passing `--profile full` exposes the 64-tool expanded surface separately. Contextual store paths are never baked into `tools.json`.

Validation relies on `go run ./cmd/validate --name whitemagic` locally plus the exact-source Docker-enabled GitHub Actions validation run for `go run ./cmd/build --tools whitemagic`.

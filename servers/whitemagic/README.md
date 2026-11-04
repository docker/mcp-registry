# WhiteMagic stable local server

The source pin is f022e99f20dac969ab2370ec65579456f998e23e (stable9.3.4).
The entry describes local stdio, not the separate hosted MCP service.

`tools.json` was refreshed from the stable9.3.4 Linux x86-64 server's actual
`tools/list`: 30 entries including the wm dispatcher. Its wider routed catalogue
is a different count. Property constraints are retained while MCP descriptions
are mapped to registry `desc` fields. The wm description's runtime store path
is removed; that contextual suffix is not a portable catalogue declaration.

The pinned npm launcher passed an owned scratch-store create/read/process-
restart/read control. Those results do not establish a Docker registry build,
Windows/macOS install, independent operator or hosted service outcome. Registry
validation and image build remain required; the local follow-through host has
no running Docker daemon. Preserve failed/pending CI as such.

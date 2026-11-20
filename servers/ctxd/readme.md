# CTXD

ctxd is open-source context infrastructure between agents and data. It synchronizes
source content into a context filesystem so agents can inspect paths and read
the current source.

## Setup

Set `source` to an absolute folder path on your Docker host. It must be readable
by container UID 10001. The folder is mounted read-only at `/sources`; no other
host directories are mounted. No account, credentials or runtime network access
are required.

The container packages ctxd 0.1.16 from PyPI. It creates an isolated context store,
synchronizes the mounted folder and starts the tracker before serving MCP over
stdio. The store is rebuilt at container startup. Source paths retain the
`local-files/source/` prefix when a gateway starts another container.

## Tools and workflow

- Inspect paths with `ls`, `tree`, `glob` and `stat`.
- Read source content with `read` and `read_lines`.
- Edit a source file, allow synchronization to finish, then read its path again.
  A running tracker checks for changes every ten seconds plus processing time.

The tracker accepts its documented text formats, not arbitrary file formats.
The published package also advertises `grep`, but that call fails because
`CtxfsFiles.grep` is missing in 0.1.16. Python and the CLI provide exact text search.
The MCP client may send returned source content to its configured model provider.

[Project and setup](https://github.com/ctxddev/ctxd-agent-plugin)
· [Product source](https://github.com/ctxd-dev/ctxd)
· [ctxd website](https://ctxd.dev/_lzx/)

Support and private security reports: hi@ctxd.dev.

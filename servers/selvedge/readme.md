# Selvedge

[Documentation and setup](https://selvedge.sh/start/quickstart/)

Selvedge stores explicitly recorded decisions and rejected approaches in local
SQLite. Agents must call its tools to record and retrieve decisions; adding an
MCP connection alone does not guarantee capture.

## Persistent project storage

Create a `.selvedge` directory in the project, then set `data_dir` to its absolute
host path in the MCP Toolkit configuration. Docker mounts that directory at
`/data`; the pinned image uses `/data/selvedge.db`.
Create the directory yourself and check the path before connecting. Docker's
bind-mount behavior can otherwise create a missing path for you.

Use a different directory for unrelated projects. Point a native Selvedge
installation at the same `.selvedge` directory to share that project's history.
The directory contains the decision data you or your agent explicitly save.

This entry uses the MCP server over stdio with runtime networking disabled. It
does not install the native coding-client lifecycle hooks. Follow the setup
guide for your client and verify a saved decision can be retrieved in a later
session.

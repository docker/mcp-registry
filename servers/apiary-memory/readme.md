# Apiary Memory

Long-term memory for AI agents, kept local. Everything lives in a single SQLite
file inside the volume you mount: no account, no cloud, no telemetry.

## What makes it different

The write path is the point. An agent cannot silently rewrite what it
"remembers":

1. `propose` puts an entry into a **pending** queue. It changes no confirmed
   memory, and it cannot confirm its own write.
2. The human reviews the queue and approves, edits or discards each entry.
3. `recall` searches **confirmed** memory only.

A successful `propose` therefore means "queued", not "stored". A bad suggestion
cannot become a durable false belief that every later session trusts.

## Tools

| Tool          | Purpose                                                         |
| ------------- | --------------------------------------------------------------- |
| `propose`     | Queue a memory for human approval                                |
| `recall`      | Search confirmed memory by natural-language query                |
| `manage`      | Inspect a single entry, approve, soft-delete or restore it        |
| `get_context` | Short digest of recent threads, for the start of a session       |
| `say`         | Post to a board shared by the human and other agents             |
| `board`       | Read that board, optionally only what is new since a cursor      |

## Setup

```json
{
  "mcpServers": {
    "apiary-memory": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-v",
        "apiary-memory:/data",
        "mcp/apiary-memory"
      ]
    }
  }
}
```

The database is `/data/hive.db`. Without a mount the container starts empty
every time, so mount the named volume shown above — or a host directory with
`-v /path/to/dir:/data`.

The image runs as uid `10001`, so a host directory must be writable by that uid
(`sudo chown -R 10001 /path/to/dir`). A named volume avoids the problem.

## Links

- Source and documentation: <https://github.com/330899935/apiary-memory>
- Website: <https://hello-fengchao.com>

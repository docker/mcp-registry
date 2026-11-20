# Progressor

Progressor keeps your AI coding agent's plan (goals, steps, decisions, notes and things for you to try by hand) in a `.progressor` folder inside your project. A new conversation, or one whose context was compressed, picks the plan up where the last one left off. Free for personal and other noncommercial use ([license](https://github.com/racstan/progressor-mcp/blob/main/LICENSE.md)).

Full documentation: https://github.com/racstan/progressor-mcp#readme

## Setup

1. Set **project** to the full path of your project's folder. Progressor creates `.progressor/` there and keeps the plan in it, so the folder is mounted writable.
2. The project list and your settings are kept in the `progressor-data` Docker volume, shared with the dashboard below.

The MCP gateway only mounts a host folder writable when you allow that exact folder. If the gateway says the folder is "outside allowed roots", start it with:

```
MCP_GATEWAY_DOCKER_BIND_ALLOW_WRITABLE_PATHS=/path/to/your/project
```

## The dashboard

The Toolkit does not publish ports, so the live dashboard runs in its own small container, which reads the same plan. Mount the same project folder at the same path the Toolkit uses, and the same `progressor-data` volume:

```
docker run -d --name progressor --restart unless-stopped \
  -v progressor-data:/data \
  -v /path/to/your/projects:/path/to/your/projects \
  -p 127.0.0.1:4800:4800 \
  mcp/progressor daemon run
```

On Windows the Toolkit mounts `C:\Projects\shop` at `/C/Projects/shop`, so mount it the same way: `-v "C:\Projects:/C/Projects"`.

Then open http://localhost:4800. Every project the Toolkit has used shows there, and the page follows the plan as your AI changes it.

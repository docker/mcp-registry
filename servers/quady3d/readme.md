# Quady3D

[Official installation guide](https://quady3d.com/integrations/mcp) · [Agent Skill](https://quady3d.com/integrations/agent-skill)

Requires a Quady3D account and a limited Generate and download API key. Cloud generation uses account credits and model delivery requires download access. The MIT client license does not cover the hosted service or generated models.

Mount an existing writable project directory to `/project` and a separate private recovery directory to `/home/node/.local/state/quady3d`. Both must be writable by the image’s non-root UID 1000; the recovery directory must be outside the project on the host as well. Do not expose the recovery directory as an allowed project root. Preserve it across restarts so interrupted submissions can be recovered using the original action ID. Configure the key through the Toolkit secret store; never put it in a public file.

Eight tools cover text/single-image generation, quoting, account checks, original-task recovery and saving files. New paid tasks need the user’s prior credit authorization. Quotes do not create tasks. Retry an interrupted delivery independently of generation.

## Authorize the selected folders

Docker MCP Gateway versions that enforce writable bind permissions also need the two selected host folders in the gateway launch environment. On macOS or Linux, set:

```sh
MCP_GATEWAY_DOCKER_BIND_ALLOW_WRITABLE_PATHS="/absolute/path/to/project:/absolute/path/to/private-recovery"
```

Use the exact existing project and private recovery folders configured in this entry. On Windows use the operating system path-list separator (`;`) between the two paths. This permission lets the gateway mount the selected folders; the Quady3D server still limits asset access to `/project` and keeps its recovery records in the separate private mount.

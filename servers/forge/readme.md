# Reliant Forge

Documentation: https://github.com/reliant-labs/forge/blob/main/docs/mcp.md#with-docker

Forge is an Apache-2.0 CLI and MCP server, separate from the Reliant agent workspace (Business Source License 1.1, not OSI open source).

The source-built image starts `forge mcp` over stdio as uid 10001 in `/workspace`. Mount a supported project there with suitable write permissions; do not mount credentials or the Docker socket. Deployment is an explicit, networked operation, not part of startup or tool discovery.

`deploy_claimable` supports existing Vite sites, Next.js static exports, and supported Dockerfile-backed HTTP apps through Forge's documented adoption/deployment path. Do not assume an arbitrary Next.js SSR app or privileged workload is supported anonymously. A successful deployment returns a live URL and a private claim link. Unclaimed deployments expire after 72 hours; per-network/per-machine rate limits and container egress restrictions apply. A URL or HTTP 200 alone is not evidence of application functionality or persistence.

The image includes git and CA certificates, but not Node/npm, Go, buf, or a Docker daemon. Local builds or tools requiring the project's own toolchain need additional prerequisites; use the installed CLI or npm route on the host for those cases. Mounting a project alone does not install its build toolchain. The catalog entry does not claim hosted end-to-end deployment validation.

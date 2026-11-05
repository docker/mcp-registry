# MotionSpec

Hosted MCP server for verified web motion: `motion_catalog` (40 verified primitives + authoring rules) and `motion_validate` (fail-closed spec validation that reports WCAG 2.2.2 pause-path candidates). Keyless, read-only, rate-limited (60 req / 10 s per IP).

- Documentation: https://motionspec.dev/docs
- Install guide (Cline/Claude/stdio): https://github.com/MasterPlayspots/motionspec/blob/main/llms-install.md
- Endpoint: https://api.motionspec.dev/mcp (streamable HTTP)

```markdown
# Add BioLab MCP Research Server

This pull request adds BioLab MCP Research Server to the MCP Registry.

- Repository: https://github.com/aguennoune/biolab-mcp-research-server
- Image: ghcr.io/aguennoune/biolab-mcp:latest (recommended: pin to a release tag like v0.1.0)
- Description: Plateforme de recherche biomoléculaire multi‑agents supervisée par IA pour la recherche CRISPR/Cas.
- Healthcheck endpoint: /healthz
- Exposed port: 8080
- License: MIT

Checklist:
- [x] Image published on GHCR
- [ ] Image tag pinned to a release (recommended)
- [x] Repo includes README with usage and healthcheck
- [x] Manifest file added at services/biolab-mcp-research-server.yaml
- [ ] Optional: SBOM/security scan attached

Notes for maintainers:
- If you prefer a different manifest path or schema keys, adapt the PR and the manifest accordingly.
- The contributor can provide a Trivy scan or SBOM on request.
```
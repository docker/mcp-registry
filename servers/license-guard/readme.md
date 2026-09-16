# LicenseGuard MCP Server

Docs: https://github.com/rccaoki-wq/license-guard

Decides whether an open source dependency's license actually obligates you, given how your project ships.

Generic license scanners answer a different question — *what license is this?* — and then warn on everything. LicenseGuard evaluates the license against your distribution model, so the same license produces different verdicts:

| How you ship | AGPL-3.0 dependency |
|---|---|
| SaaS (network-accessible) | **blocked** — §13 network clause |
| Internal use only | allowed |
| Distributed binary / on-prem | **blocked** — inherited GPL distribution terms |
| **devDependency** (never in the artifact) | **allowed** |

That last row is the point: a build-time linter under AGPL never ships, so it triggers nothing. Tools that warn on it anyway train people to ignore every warning they produce.

The same distinctions run through the rest of the license landscape and are not interchangeable:

- **GPL** obligations attach to *distribution*. Running GPL code as a network service is not distribution.
- **AGPL** adds §13, which attaches to *network interaction* — a separate trigger from GPL's distribution terms.
- **MPL / EPL / CDDL** are file-scoped and linkage-independent. MPL-2.0 §3.3 explicitly permits distributing a Larger Work under your own terms.
- **LGPL** is the one that actually depends on linkage: static linking carries a relinking obligation, dynamic linking does not.

## Tools

| Tool | When to call it |
|---|---|
| `check_dependency_license` | Before adding a single dependency |
| `check_manifest_licenses` | To audit a whole manifest or lockfile |
| `explain_license` | To see what a license requires across every distribution model |

## Coverage

Ecosystems: npm, PyPI, Go modules, crates.io

Manifests: `package-lock.json`, `pnpm-lock.yaml`, `yarn.lock`, `requirements.txt`, `pyproject.toml`, `poetry.lock`, `uv.lock`, `go.mod`, `go.sum`, `Cargo.toml`, `Cargo.lock`

Transitive dependencies are covered from lockfiles. `package-lock.json` v2/v3 embeds a license for every entry, so a full transitive audit runs with zero registry lookups.

Dependencies that cannot be resolved are reported as `not-checked` or `review` and never as `allowed` — an incomplete scan is never presented as clean.

## Access

No authentication. Stateless Streamable HTTP at `https://license-guard.rcc-aoki.workers.dev/mcp`.

Listed in the official MCP registry as `io.github.rccaoki-wq/license-guard`.

## Privacy

Manifest contents, package names and IP addresses are not stored. A local stdio build is also available in the repository for teams that prefer the manifest never to leave their machine — in that mode only package names and versions reach public registries.

## Disclaimer

LicenseGuard provides information derived from published license texts and declared dependency metadata. It is not legal advice, and using it does not create an attorney-client relationship. Verdicts rest on the license information a package declares; they do not claim to identify every obligation or violation.

## License

Apache-2.0

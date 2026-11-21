# Incident Triage MCP

`incident-triage-mcp` is a Model Context Protocol (MCP) server for evidence-driven incident response workflows.

It exposes auditable triage and action tools to AI hosts/agents while keeping direct infrastructure access out of the model runtime.

## What It Provides

- Incident triage summaries from normalized evidence bundles (`EvidenceBundle v1`)
- Evidence retrieval / wait-for-bundle orchestration
- Jira / ServiceNow ticket drafting and creation
- Slack / Teams notifications
- Optional Airflow DAG trigger integration
- Safe actions (RBAC, confirm tokens, audit logging, idempotency)

## Safe Defaults (Docker MCP Toolkit / Catalog)

This registry entry is intended to be safe to run locally without external dependencies:

- `MCP_TRANSPORT=stdio`
- `WORKFLOW_BACKEND=none`
- `EVIDENCE_BACKEND=fs`
- `JIRA_PROVIDER=mock`
- `BUNDLE_ONLY_MODE=true`

You can enable Airflow/Jira/Slack/Teams later by supplying env vars and secrets.

## Common Modes

### Standalone / Local

- `WORKFLOW_BACKEND=none`
- `EVIDENCE_BACKEND=fs`

### Airflow-Orchestrated (local or prod)

- `WORKFLOW_BACKEND=airflow`
- `EVIDENCE_BACKEND=fs` (local PVC/filesystem) or `s3` (prod object storage)

## Project

- Repository: https://github.com/felixkwasisarpong/incident-triage-mcp
- PyPI: https://pypi.org/project/incident-triage-mcp/

For full setup (Docker, Kubernetes, Airflow, Helm, and provider configuration), see the repository `README.md`.

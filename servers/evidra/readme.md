# Evidra

Designed for AI agents operating your infrastructure.
Fail-closed policy guardrails for AI agents running kubectl, terraform, helm, argocd, and oc.

## Overview

Evidra is a kill-switch for AI agents managing infrastructure. Experimenting with AI in staging? Add a kill-switch first. Blocks dangerous ops. Allows safe ones. Every decision logged.

## What it catches

- Protected namespace deletions (e.g. kube-system, default)
- Mass resource removal
- Public S3 bucket creation
- Wildcard IAM policies
- Dangerous ArgoCD sync operations
- Other high-impact infrastructure mistakes

## Key features

- **Zero-config**: embedded OPA policy bundle, works out of the box
- **Fail-closed**: if policy evaluation fails, the operation is blocked
- **Deterministic**: pure OPA/Rego rules, no LLM in the decision loop
- **Evidence trail**: every allow/deny decision logged with cryptographic signatures
- **Lightweight**: focused on catastrophic scenarios, not a full compliance engine

## Links

- [GitHub Repository](https://github.com/vitas/evidra)
- [Security Model](https://github.com/vitas/evidra/blob/main/docs/SECURITY_MODEL.md)
- [Landing Page](https://evidra.samebits.com)

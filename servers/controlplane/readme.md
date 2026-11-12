# Control Plane MCP Server

The **Control Plane MCP Server** lets any AI tool with [Model Context Protocol](https://modelcontextprotocol.io/) support deploy and operate workloads, secrets, networking, storage, and managed Kubernetes across AWS, GCP, Azure, and private clouds — through natural-language conversations.

[Control Plane](https://controlplane.com) runs containerized workloads across AWS, GCP, Azure, OCI, and your own hardware under a single API. This server exposes that platform to your assistant as a focused set of typed tools that mirror the platform's own validation, so the assistant builds correct requests and surfaces precise errors instead of guessing.

This is a **remote** server — there is nothing to install or run locally. Point your assistant at the endpoint, complete the one-time OAuth consent for your organization, and start issuing requests.

## Endpoint

```text
https://mcp.cpln.io/mcp
```

Transport: `streamable-http`.

## Authentication

OAuth with per-org consent. Your AI tool walks you through the flow on first use — no API keys to copy or store. Every action runs as you, scoped by your organization's policies, so an assistant can never exceed your own permissions.

## Toolset profiles

The full catalog is larger than most AI tools can hold in context at once, so each connection selects a **toolset profile** with the `toolsets` query parameter. The profiles are nested — pick one:

| Profile | Tools | Endpoint | Includes |
| :------ | :---: | :------- | :------- |
| `core` *(default)* | 58 | `https://mcp.cpln.io/mcp` | Everyday deploy-and-operate workflows: workloads, GVCs, secrets, custom domains, identities, policies, the template catalog, volume sets, observability, and Terraform export — plus generic read/delete for **every** resource kind and a raw-API escape hatch. |
| `mk8s` | 86 | `https://mcp.cpln.io/mcp?toolsets=mk8s` | Everything in `core`, plus the BYOK managed-Kubernetes family for provisioning and updating clusters across cloud providers. |
| `full` | 151 | `https://mcp.cpln.io/mcp?toolsets=full` | Every tool, including cloud-account registration, org IAM, network agents, IP sets, audit contexts, external logging, quotas, traces, advanced workload tuning, and volume snapshots. |

Keep the default `core` for clients that load every tool up front (Cursor, Codex, Gemini CLI, Claude Desktop, VS Code). Use `full` for clients with deferred tool loading, like Claude Code.

## What you can do

- **Deploy & operate workloads** — Create serverless, standard, cron, stateful, and VM workloads with their containers, scaling, and exposure; update images and resources; run cron jobs; exec into containers.
- **Manage GVCs** — Create and configure Global Virtual Clouds and choose the cloud locations your apps run in.
- **Handle secrets** — Create and update opaque, dictionary, TLS, Docker-registry, and cloud-provider secrets, and wire them into workloads safely.
- **Configure custom domains** — Put domains in front of workloads with routing and TLS.
- **Control identity & access** — Manage identities, policies, groups, and service accounts for least-privilege RBAC.
- **Install from the template catalog** — Provision Postgres, Redis, Kafka, MongoDB, MySQL, and other managed components.
- **Provision storage** — Create volume sets, mount them to stateful workloads, and manage snapshots and expansion.
- **Observe** — Query logs, metrics, traces, and workload events to troubleshoot and right-size.
- **Export to IaC** — Convert live resources to Terraform.
- **Connect clouds & networks** — Register AWS/GCP/Azure/NATS cloud accounts and reach private VPCs and on-prem networks through network-bridge agents.
- **Run managed Kubernetes** — Provision and update BYOK mk8s clusters across 11 providers.
- **Work generically** — List, read, and delete *any* resource kind, with a raw-API tool for anything not yet typed.

## Example prompts

```text
# Deploy
Create a GVC called "my-app-prod" in Frankfurt and Virginia, then deploy a
public workload "frontend" from nginx:latest with min 2 replicas on port 80.

# Operate
Show me the status and replica count for the "backend" workload, then get the
last 15 minutes of error logs.

# Secure
Create an opaque secret "db-url" and grant the "backend" workload access to it.

# Promote
Show the differences between the staging and production "api" workloads, then
update production to match staging.

# Manage data services
Install Postgres from the template catalog into "my-app-prod".

# Infrastructure as code
Export the "my-app-prod" GVC and its workloads to Terraform.
```

## Safety

Destructive tools (deletes, removes, overwrites) are labeled as such, and assistants present the impact and ask for confirmation before running them. Everything you do is captured in Control Plane's audit trail.

## Compatible tools

Any client that supports remote MCP servers can connect, including Claude Code, Claude Desktop, Claude Web, Codex, Gemini CLI, Cursor, VS Code, Antigravity, Amp, and OpenCode.

> **Using Claude Code, Codex, or Gemini CLI?** Start with the [Control Plane AI Plugin](https://docs.controlplane.com/ai/plugin) instead — it auto-configures this MCP server and adds skills, agents, commands, and production guardrails on top.

## Documentation

- MCP Server: https://docs.controlplane.com/ai/mcp
- Tools reference: https://docs.controlplane.com/ai/tools
- Usage examples: https://docs.controlplane.com/ai/examples
- AI Plugin: https://docs.controlplane.com/ai/plugin
- Control Plane: https://controlplane.com

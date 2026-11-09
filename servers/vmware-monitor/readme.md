# VMware Monitor

Read-only VMware vCenter/ESXi MCP integration, maintained by the VMware Skills community. MIT licensed. No infrastructure mutation tools are exposed.

## Local setup

1. Configure targets using the source repository [README](https://github.com/vmware-skills/VMware-Monitor#readme). A suitable vCenter read-only account and network access from Docker to vCenter/ESXi are required. Passwords belong in the local `.env` file, never in `config.yaml` or this directory entry.
2. Select an absolute `config_dir` containing the prepared `config.yaml` and optional `.env`. Restrict credential files to the account owner. The mount is writable because existing server startup can normalize local credential-file permissions and encoding; it may update `.env` on the host.
3. Select a separate absolute, writable `audit_dir` for shared vmware-policy audit data. Do not point either parameter at an entire home directory. The source image currently runs as root; on Linux, newly created files may be root-owned.
4. Enable the server in Docker MCP Toolkit. No hosted VMware endpoint or shared credentials are supplied by this project. Docker Desktop must have network reachability and appropriate trust for the configured endpoint.

The static tools.json records 33 tools discovered from published package 2.6.0 using MCP initialize, ping and tools/list, without infrastructure calls. The source container was separately built and protocol-tested in [CI](https://github.com/vmware-skills/VMware-Monitor/actions/runs/42547928630), with networking disabled. These metadata checks do not claim an authenticated vCenter test through Docker Toolkit.

Security contact: zhouwei008@gmail.com. Do not post credentials in issues or pull requests. See the source repository for policy, audit, configuration and troubleshooting documentation.

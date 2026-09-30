# SousChef

Convert Chef, SaltStack, Puppet, PowerShell and Bash automation to Ansible, and
plan Ansible upgrades.

## Workspace configuration

Set **workspace** to an existing, absolute host directory containing the
automation to migrate. Docker mounts that directory read-write at
`/workspace` and sets `SOUSCHEF_WORKSPACE_ROOT=/workspace`.

Use **container paths** in tool calls, for example:

- `parse_recipe(path="/workspace/cookbooks/web/recipes/default.rb")`
- `read_file(path="/workspace/cookbooks/web/metadata.rb")`

A Windows host directory such as `C:/work/souschef` is still exposed as
`/workspace` inside the Linux container. Do not pass the Windows host path to
a tool.

Use a dedicated migration directory: tools can read and write its contents.
Copy source automation into it and keep backups. Do not mount your entire
home directory. On native Linux, ensure UID/GID 1001 can read inputs and write
generated output; grant access to this directory only. The image runs as a
non-root user. Docker Desktop must be able to share the host directory.

No credentials are required for local parsing or deterministic resource
conversion. Operations connecting to Chef, AWX/AAP or other external services
have their own connection requirements; the basic catalogue configuration
does not configure those services.

## Reviewer fixture

`fixtures/default.rb` contains a minimal package-install recipe.
Copy it to `<workspace>/cookbooks/web/recipes/default.rb`, then ask your MCP
client to parse `/workspace/cookbooks/web/recipes/default.rb` and convert
the `package[nginx]` resource with action `install`.

Expected conversion: an `ansible.builtin.package` task with `name: nginx`
and `state: present`. Calling `read_file` with `/etc/passwd` must be rejected
because the path is outside the configured workspace.

## Import the proposed catalogue locally

From a checkout of the Docker MCP Registry branch containing this entry:

```bash
task validate -- --name souschef
task build -- --tools souschef
task catalog -- souschef
docker mcp catalog import ./catalogs/souschef/catalog.yaml
docker mcp config write '{"souschef":{"workspace":"/absolute/path/to/migration-workspace"}}'
docker mcp server enable souschef
docker mcp gateway run --servers souschef
```

Merge the `souschef.workspace` setting into your existing configuration if
you already use other servers; `config write` replaces the configuration.
The proposed entry is not available in the public Docker catalogue until
Docker approves and publishes it.

Documentation: https://kpeacocke.github.io/souschef/
Source: https://github.com/kpeacocke/souschef

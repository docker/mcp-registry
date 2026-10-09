# Agend-sh

[Agend](https://agend.sh) gives AI agents persistent Linux environments through
MCP. Run commands, drive REPLs and terminal apps, edit files, collect background
task output, and expose development servers through public HTTPS URLs.

The container runs the open source CLI's local stdio MCP bridge. Commands and
remote files live in your hosted Agend environment. An Agend account and
available environment quota are required.

## Setup

1. Install the [Agend CLI](https://github.com/agend-sh/cli#quick-start) and run
   `agend login`.
2. Follow the [container authentication guide](https://github.com/agend-sh/cli/blob/main/docs/mcp-registry.md#connect-using-the-container)
   to load your active account token without printing it.
3. Configure the `agend-sh.api_token` secret in Docker MCP Toolkit. Docker
   passes it to the bridge as `AGEND_API_TOKEN`.
4. Enable Agend-sh in your MCP Toolkit profile and connect your AI client.
5. Ask the agent to call `list_environments`, then use the returned environment
   ID or name in subsequent tool calls. Use `env_create` if your account has
   quota and needs an environment.

Refresh the secret after reauthentication when the token expires. The bridge
does not inherit the native CLI's selected environment.

## Interactive terminals

Use `shell_exec` with `interactive=true` to launch a REPL or terminal app, then
use `shell_send_raw` for subsequent input. Keep the MCP connection open while
interacting with the process.

The `input_wait` result records whether a guest terminal input-wait event was
observed during that response. It is process feedback; it does not prove that
an application or service is ready. See the tool descriptions for the full
event semantics and interaction workflow.

## Files

Use `file_write` to create remote text files directly. `file_upload` and
`file_download` use local paths inside the container's `/workspace`; they do
not automatically access your computer's directories. For transfers to or
from your computer, see the
[workspace mount instructions](https://github.com/agend-sh/cli/blob/main/docs/mcp-registry.md#local-file-transfers).

## Package and support

This entry uses the public Agend v1.2.14 container, pinned by digest, for
Linux amd64 and arm64. Its 23 tool descriptions are generated from the
same release's embedded schemas. Tool metadata can be reviewed without
credentials; hosted tool calls require a valid account token.

- [Source and documentation](https://github.com/agend-sh/cli)
- [MIT license](https://github.com/agend-sh/cli/blob/main/LICENSE)
- [Report security issues privately](https://github.com/agend-sh/cli/blob/main/SECURITY.md)

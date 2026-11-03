# sourcebolt (NotebookLM, unofficial)

Unofficial MCP server for Google NotebookLM. Not affiliated with Google; it uses NotebookLM's web calls with your own session and can break when NotebookLM changes.

Docs: https://github.com/candidloom-hub/sourcebolt

## Sign in

Set the `sourcebolt.notebooklm_cookies` secret to your NotebookLM session cookies as one Cookie header line. How to export them: https://github.com/candidloom-hub/sourcebolt/blob/main/docs/auth.md#sign-in-with-an-environment-variable-containers

- sourcebolt keeps the cookies in memory only and never writes them to disk.
- Treat the secret like a password. Anyone who can run `docker inspect` on the container can read it.
- Google sessions expire. When tool calls return `AUTH_REQUIRED`, export fresh cookies and update the secret.

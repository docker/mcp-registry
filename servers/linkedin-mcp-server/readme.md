# LinkedIn

Give MCP clients access to LinkedIn profiles, companies, jobs and messages through your own logged-in browser session.

## Sign in

Create the persistent session volume and sign in once:

```bash
docker volume create linkedin-mcp-server-session
docker run -it --rm \
  -v linkedin-mcp-server-session:/home/pwuser/.linkedin-mcp \
  -p 127.0.0.1:6080:6080 \
  mcp/linkedin-mcp-server \
  --login --login-viewer
```

Open the full URL printed by the command and complete the sign-in. Let the command exit on its own so the session is stored completely. Docker MCP reuses the same named volume when it starts the server.

Repeat the login command when the session expires. To remove the stored session, disable the server and run `docker volume rm linkedin-mcp-server-session`.

Docker MCP does not currently expose the server’s optional proxy settings. If your network requires a proxy, use the [direct Docker setup](https://github.com/stickerdaniel/linkedin-mcp-server#docker-setup).

For configuration and troubleshooting, see the [project documentation](https://github.com/stickerdaniel/linkedin-mcp-server#docker-setup).

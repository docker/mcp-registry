Docs: https://contentin.io/linkedin-mcp-server/setup/

Requires a ContentIn account. Sign in with OAuth when prompted. Scheduling and publishing show the exact post text and need your confirmation.

The default authorization grants read and write access, which covers writing and scheduling posts. Publishing a post immediately (`publish_post`) needs the publish permission:

```
docker mcp oauth authorize contentin --scopes "read write publish"
```

# Easy Render | Photo to 3D

Turn any Product Photo into a 3D model.

Start with a free angle-sheet preview, then build the model and download GLB and USDZ files.

## Connect

- **Endpoint:** `https://easyrender.ai/api/mcp`
- **Transport:** Streamable HTTP
- **Authentication:** Sign in with OAuth in a compatible MCP client, or send an easyrender API key as `Authorization: Bearer er_your_key`. An easyrender account is required; manage keys in [Connectors](https://easyrender.ai/connectors).

## Main tools

| Tool | Purpose |
| --- | --- |
| `create_render` | Start an angle-sheet preview from a photo. |
| `get_render` | Check progress and retrieve result links. |
| `get_credits` | Check current account usage. |
| `render_3d` | Build a 3D model from a ready preview. |

Previews are free, subject to account limits. Ask your agent to show the preview before you approve a 3D build. Poll `get_render` until the model is ready; private downloads require authentication.

Docs: https://easyrender.ai/mcp

REST API: https://easyrender.ai/image-to-3d-api

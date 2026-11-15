# easyrender

Turn one photo of furniture, decor or a building exterior into a 3D model: start with a free angle-sheet preview, then build the model and download GLB and USDZ files.

## Connect

- **Endpoint:** `https://easyrender.ai/api/mcp`
- **Transport:** Streamable HTTP
- **Authentication:** Sign in with OAuth in a compatible MCP client, or send an easyrender API key as `Authorization: Bearer er_your_key`. An easyrender account is required; manage keys in [Connectors](https://easyrender.ai/connectors).

## Main tools

| Tool | Purpose |
| --- | --- |
| `create_render` | Start an angle-sheet preview from a photo. |
| `get_render` | Check progress and retrieve result links. |
| `get_credits` | Check your credit balance and current build cost. |
| `render_3d` | Spend credits to build a 3D model from a ready preview. |

Previews are free, subject to account limits. A finished model costs **64 credits ($0.64)**. Ask your agent to show the preview and check the cost before you approve a paid build. Poll `get_render` until the model is ready; private downloads require authentication.

Docs: https://easyrender.ai/mcp

REST API: https://easyrender.ai/image-to-3d-api

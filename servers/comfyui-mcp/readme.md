# ComfyUI MCP

## Documentation

For complete documentation on the ComfyUI MCP server, visit:
- **Official Docs:** https://comfyui-mcp.artokun.io/docs
- **GitHub Repository:** https://github.com/artokun/comfyui-mcp
- **Agent Panel:** https://github.com/artokun/comfyui-mcp-panel
- **Local LLM Guide:** https://comfyui-mcp.artokun.io/docs/local-llms
- **Discord Community:** https://discord.gg/cW9arBhzCu

## Quick Start

This MCP server connects your LLM to ComfyUI, enabling:

- **Text-to-image generation** (Flux, SDXL, SD1.5, Turbo, Lightning)
- **Video generation** (WAN, LTX-Video 2.3)
- **Audio generation** (Stable Audio 3, ACE Step 1.5)
- **Workflow automation** (create, modify, validate, run)
- **Model management** (search, download, list)
- **Custom node installation** (discovery, installation, management)
- **Autonomous sidebar agent** (ComfyUI Agent Panel) for direct canvas control

## Standalone Headless + Sidebar Agent

This package provides TWO ways to use comfyui-mcp:

### 1. **Headless MCP Server** (this submission)
Run the MCP server standalone in Docker or Node.js:
- Use with Claude Desktop, Claude Code, ChatGPT, Cursor, etc.
- Controlled purely via tool calls
- No UI component required

### 2. **ComfyUI Agent Panel** (optional sidebar)
Install the autonomous agent directly in ComfyUI:
- GitHub: https://github.com/artokun/comfyui-mcp-panel
- ComfyUI-Manager: Search `comfyui-agent-panel`
- Sidebar agent that drives your live canvas
- Real-time workflow editing and execution
- Works with any backend (Claude, ChatGPT, Gemini, Ollama, etc.)

Use either independently or together for maximum flexibility.

## Supported LLMs

- Claude Desktop / Claude Code
- ChatGPT via Codex
- Google Gemini
- Ollama (local models)
- OpenAI-compatible endpoints (DeepSeek, Grok, etc.)

## Environment Variables

- `COMFYUI_URL` - URL to your ComfyUI instance (auto-detected if not set)
- `COMFYUI_PATH` - Path to ComfyUI data directory (auto-detected if not set)
- `CIVITAI_API_TOKEN` - For downloading models from CivitAI
- `HUGGINGFACE_TOKEN` - For accessing HuggingFace models
- `GITHUB_TOKEN` - For skill generation and avoiding API rate limits
- `LOG_LEVEL` - Logging verbosity: debug, info, warn, error

## Features

- **182 MCP tools** for complete ComfyUI control
- **35 AI skills** for model families (Flux, WAN, LTX, Qwen, etc.)
- **Workflow composition** from templates
- **Real-time progress monitoring**
- **Model & custom node management**
- **Generation tracking** with SQLite database
- **Both stdio and HTTP transports** supported

## Use Cases

1. **AI agents** that autonomously generate images/video
2. **Image workflows** with natural language
3. **Model management** and discovery
4. **Custom node** installation and testing
5. **Batch generation** and parameter sweeps

## More Information

Join the Discord for help, model tips, and release announcements:
https://discord.gg/cW9arBhzCu

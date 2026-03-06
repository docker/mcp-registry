# Recallium - AI Memory & Intelligence MCP Server

Recallium provides persistent memory, insights, and context for AI agents. It enables AI assistants to remember decisions, track progress, search across projects, and get AI-powered insights.

## Features

- **Persistent Memory**: Store and retrieve memories across sessions with automatic categorization
- **Semantic Search**: Find relevant context using AI embeddings (1024-dimensional vectors)
- **Project Management**: Organize work by projects with briefs, PRDs, and implementation plans
- **Insights Engine**: Get AI-powered analysis of patterns, trends, and learnings
- **Task Tracking**: Manage tasks with bidirectional memory linking
- **Structured Thinking**: Step-by-step reasoning sequences for complex problems
- **Document Processing**: Upload and search PDFs and text documents

## Prerequisites

- **Docker** - Required for containerization

## Quick Start

### 1. Pull the Image

```bash
docker pull recalliumai/recallium:latest
```

### 2. Create Environment File

Create a `recallium.env` file:

```bash
# Database (required)
POSTGRES_PASSWORD=recallium_password  # Change in production!

# Ports
HOST_UI_PORT=9001
HOST_MCP_PORT=8001
```

### 3. Run the Container

```bash
docker run -d \
  --name recallium \
  --env-file recallium.env \
  -p 8001:8000 \
  -p 9001:9000 \
  -v recallium-data:/data \
  recalliumai/recallium:latest
```

### 4. Complete Setup Wizard

Open http://localhost:9001 in your browser. The Setup Wizard will guide you through:

1. **Select LLM Provider** - Choose from 5 supported providers:
   - **Ollama** - Free local inference (requires [Ollama](https://ollama.ai) installed)
   - **OpenAI** - GPT-4, GPT-4o, GPT-3.5 Turbo
   - **Anthropic** - Claude 3.5 Sonnet, Claude 3 Opus/Sonnet/Haiku
   - **Google Gemini** - Gemini Pro, Gemini 1.5 Pro/Flash
   - **OpenRouter** - Access to 100+ models via single API

2. **Configure API Keys** - Enter your API key (securely stored in encrypted vault)

3. **Test Connection** - Verify your provider is working

4. **Set Up Failover** (Optional) - Configure backup providers for reliability

**Note:** Embeddings run locally at no cost using the bundled GTE-Large model.

**Access points:**
- Web UI: http://localhost:9001
- MCP API: http://localhost:8001/mcp
- Health check: http://localhost:8001/health

## MCP Client Configuration

### Claude Desktop (JSONB Extension)

1. Download the JSONB extension file from the [Recallium release](https://github.com/recallium-ai/recallium/releases) or create one with the Recallium MCP configuration
2. Open Claude Desktop → Settings → Extensions → Browse Extensions → Add Custom
3. Browse to and select the `.jsonb` file to install
4. Go to Connectors → Recallium and enable each tool you want to use

### Claude Code

```bash
claude mcp add --transport http recallium http://localhost:8001/mcp
```

### HTTP-Based IDEs (Cursor, VS Code, Windsurf)

Add to your MCP configuration:

```json
{
  "mcpServers": {
    "recallium": {
      "url": "http://localhost:8001/mcp",
      "transport": "http"
    }
  }
}
```

**Config file locations (Cursor):**
- macOS: `~/.cursor/mcp.json`
- Windows: `%APPDATA%\Cursor\mcp.json`
- Linux: `~/.config/cursor/mcp.json`

### STDIO-Based IDEs (Zed, JetBrains, Cline, BoltAI)

First install the MCP client:
```bash
npm install -g recallium
```

Then configure:
```json
{
  "mcpServers": {
    "recallium": {
      "command": "npx",
      "args": ["-y", "recallium-mcp"],
      "env": {
        "RECALLIUM_SERVER_URL": "http://localhost:8001/mcp"
      }
    }
  }
}
```

## Documentation

- Website: https://www.recallium.ai
- Help: https://www.recallium.ai/help
- GitHub: https://github.com/recallium-ai/recallium
- Docker Hub: https://hub.docker.com/r/recalliumai/recallium

## License

MIT

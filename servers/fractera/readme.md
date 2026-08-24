 # Fractera

  **Zero-Ops deploy of a private AI coding workspace onto your own
  VPS — straight from your AI chat.**

  Fractera turns a bare Ubuntu VPS into a complete, private AI
  coding workspace without any DevOps. You give your AI agent only
  the server credentials, and the connector configures everything
  automatically (Nginx, HTTPS, auth, database, services) in about
  10 minutes. The deploy is IP-first: when it finishes, the
  workspace is live on plain HTTP at `http://<your-ip>:3002`.
  Attaching a custom domain with HTTPS is an optional later step
  done inside the workspace.

  ## What you get on your own server
  - **5 AI coding engines** — Claude Code, Codex, Gemini CLI, Qwen
  Code, Kimi Code (switchable on the fly)
  - **Hermes** — an autonomous orchestrator agent that can work
  while you sleep
  - **Private graph memory (LightRAG)** over your code and docs
  - Fully self-hosted on **your** hardware — nothing leaks to
  third-party clouds

  ## Tools
  - `register_and_deploy` — create the user, register the server,
  run the full deploy
  - `retry_deploy` — recover a failed deployment
  - `check_status` — check deployment progress / state
  - `get_vps_recommendation` — suggest a suitable VPS
  - `get_subdomain` — return the workspace address
  - `get_project_info` — full architecture, specs and FAQ about the
  project

  ## How to use
  1. Add the connector by URL: `https://www.fractera.ai/api/mcp`
  2. In chat, ask to deploy Fractera and provide your VPS IP +
  credentials.
  3. The agent runs the full deploy and reports back when it's live
  at `http://<your-ip>:3002`.

  ## Good to know
  - **Free** — Fractera never charges for the deployment; it runs
  on your own hardware.
  - **No password is stored** — server credentials are used only
  during the deploy and are not persisted.
  - **Open & documented** — workspace is open-source (MIT):
  https://github.com/Fractera/ai-workspace

  Full knowledge base: https://www.fractera.ai/mcp-info

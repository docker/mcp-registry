# CouchLoop EQ — MCP Server

Your AI remembers conversations. Add persistent memory and safety checks to Claude and ChatGPT.

<p align="center">
  <img src="https://raw.githubusercontent.com/wisenbergg/couchloop-mcp/master/assets/logo/couchloop_EQ-IconLogo.png" alt="CouchLoop EQ" width="120" />
</p>

## The Problem

AI assistants forget everything between sessions. Users repeat context, lose progress on multi-step workflows, and get inconsistent responses. Worse, LLMs can hallucinate, contradict themselves, or drift into harmful territory without guardrails.

**CouchLoop EQ fixes this.** It's an MCP server that gives your AI persistent memory, progress tracking, and automatic safety validation.

## Quick Start (30 seconds)

### Try Without Signup

Use the public demo server to test immediately:

**For Claude Desktop** (requires Claude Desktop v0.7.0+), add to `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "couchloop-eq": {
      "url": "https://couchloop-mcp-production.up.railway.app/mcp",
      "transport": "sse"
    }
  }
}
```

Restart Claude and try: **"Start a daily reflection session"**

> Sessions automatically persist locally on your machine (v1.0.4+). No account required.

### For ChatGPT

1. Open ChatGPT → Settings → Developer Mode
2. Add MCP connector:
   - **URL**: `https://couchloop-mcp-production.up.railway.app/mcp`
   - **Auth**: None required
3. Try: **"List available journeys"**

> **Note**: ChatGPT MCP support is in beta—expect occasional disconnects.

## Production Setup

For unlimited sessions and custom journeys:

```bash
npm install -g couchloop-eq-mcp
```

Add to Claude Desktop config:

```json
{
  "mcpServers": {
    "couchloop-eq": {
      "command": "couchloop-eq-mcp",
      "env": {
        "COUCHLOOP_API_KEY": "your-api-key"
      }
    }
  }
}
```

**Get your API key:** [couchloop.com/signup](https://couchloop.com/signup)

## How It Works

CouchLoop runs as an MCP server between your AI client and its responses. It intercepts messages, applies safety checks, and manages session state.

```
┌─────────────────┐
│ Claude Desktop  │
│   or ChatGPT    │
└────────┬────────┘
         │ MCP Protocol
         ▼
┌─────────────────┐
│ CouchLoop EQ    │
│  ┌───────────┐  │
│  │  Safety   │  │ ← Fact-checking, tone monitoring
│  │  Layer    │  │
│  └───────────┘  │
│  ┌───────────┐  │
│  │  Session  │  │ ← State persistence, checkpoints
│  │  Manager  │  │
│  └───────────┘  │
│  ┌───────────┐  │
│  │  Journey  │  │ ← Guided workflows, templates
│  │  Engine   │  │
│  └───────────┘  │
└─────────────────┘
```

### Code Example

```javascript
import { MCPClient } from '@modelcontextprotocol/sdk';

const mcp = new MCPClient({ server: 'couchloop-eq' });

// Start a guided session
const session = await mcp.call('create_session', {
  journey: 'daily-reflection',
  userId: 'user_123'
});
// → Returns session ID and first prompt

// User completes session, days pass...

// Resume with full context restored
await mcp.call('resume_session', { userId: 'user_123' });
// → "Last time you mentioned feeling energized in the mornings.
//    How has that been this week?"
```

## Features

### Persistent Memory (New in v1.0.4!)
- **Zero-config persistence**: Sessions automatically saved locally to `~/.couchloop-mcp/identity.json`
- **Session state**: Conversations survive browser closes, app restarts on the same machine
- **Checkpoints**: Save progress at key moments in multi-step workflows
- **Insights**: Capture and retrieve important user reflections over time
- **No signup required**: Works immediately, no account needed for local persistence

### Safety Checks
- **Fact-checking**: Catches fabricated information and unsupported claims
- **Consistency tracking**: Flags contradictions across the conversation
- **Conversation boundaries**: Detects manipulation patterns and harmful suggestions
- **Tone stability**: Maintains consistent personality without emotional drift

### Guided Journeys
Pre-built workflows for common use cases:

| Journey | Duration | Description |
|---------|----------|-------------|
| Daily Reflection | 5 min | Brief check-in to process your day |
| Gratitude Practice | 3 min | Notice and name things you appreciate |
| Weekly Review | 10 min | Look back and set intentions |

## Available Tools

| Tool | Description | Returns |
|------|-------------|---------|
| `create_session` | Start new session, optionally with a guided journey | Session ID, initial prompt |
| `resume_session` | Continue previous session with context restored | Last checkpoint, next prompt |
| `send_message` | Send a message with safety validation | Validated response |
| `save_checkpoint` | Save progress at a key moment | Checkpoint ID |
| `get_checkpoints` | Retrieve all checkpoints for a session | Array of checkpoints |
| `list_journeys` | List available guided journeys | Journey definitions |
| `get_journey_status` | Get current progress in a journey | Step number, completion % |
| `save_insight` | Capture a meaningful user insight | Insight ID |
| `get_insights` | Retrieve saved insights | Array of insights with timestamps |
| `get_user_context` | Get user's insights, patterns, and history | Context summary for personalization |

## Example Usage

```
"Start a daily reflection session"
"Resume my last session"
"Save this insight: I notice I'm more energized in the mornings"
"What insights have I saved?"
"How far am I in this journey?"
```

## Authentication

| Mode | Access | Limits | Best For |
|------|--------|--------|----------|
| **Demo** | Public server, no signup | Unlimited sessions with local persistence | Testing, evaluation |
| **Production** | API key from couchloop.com | Unlimited sessions, custom journeys | Production apps |

## Use Cases

- **Customer support agents**: Maintain conversation history across tickets and channels
- **Onboarding flows**: Guide users through multi-step setup with progress tracking
- **AI assistants with memory**: Remember user preferences, past decisions, and context
- **Compliance-sensitive apps**: Add safety validation for financial, legal, or healthcare AI
- **Multi-session workflows**: Any AI interaction that spans days, devices, or interruptions

> **Building wellness or therapeutic apps?** See our [wellness integration guide](https://couchloop.com/docs/wellness) for purpose-built journeys and crisis detection features.


## Support

- **Issues**: [github.com/wisenbergg/couchloop-mcp/issues](https://github.com/wisenbergg/couchloop-mcp/issues)
- **Email**: support@couchloop.com
- **Docs**: [couchloop.com/docs](https://couchloop.com/docs)

## License

MIT

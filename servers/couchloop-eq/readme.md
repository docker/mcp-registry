# CouchLoop EQ MCP Server

## Overview

CouchLoop EQ is a behavioral governance layer for Large Language Models (LLMs) that monitors for hallucination, inconsistency, tone drift, and unsafe reasoning patterns. It provides stateful conversation management with session persistence, progress checkpoints, and guided journeys - enabling AI agents like ChatGPT and Claude to maintain context across interruptions and window changes.

## Key Features

### 🛡️ Behavioral Governance
- **Hallucination Detection**: Monitors AI responses for factual consistency
- **Tone Drift Prevention**: Ensures consistent emotional tone throughout conversations
- **Crisis Detection**: Identifies critical mental health signals and triggers appropriate responses
- **Unsafe Reasoning Detection**: Catches potentially harmful logical patterns

### 💾 Session Management
- **Stateful Conversations**: Maintains context across multiple interactions
- **Cross-Window Persistence**: Sessions survive browser refreshes and tab changes
- **Resume Capability**: Pick up conversations exactly where you left off
- **Progress Checkpoints**: Save key moments and insights during conversations

### 🎯 Guided Journeys
- **Pre-built Experiences**: Structured therapeutic journeys (daily reflection, gratitude, CBT exercises)
- **Step-by-Step Progression**: Advance through curated conversation flows
- **Optional Paths**: Skip or customize journey steps based on user needs
- **Journey Memory**: Track progress across multiple journey sessions

### 🧠 Memory & Insights
- **User Context**: Builds understanding of user patterns and preferences
- **Insight Capture**: Save and retrieve meaningful realizations from conversations
- **Session History**: Access past conversations and checkpoints
- **Personalization**: Adapts responses based on accumulated user context

## Quick Start

### Using with Docker Desktop MCP Toolkit

1. **Install from Docker Registry** (coming soon):
```bash
docker pull docker/mcp-couchloop:latest
```

2. **Configure Environment Variables**:
Create a `.env` file with required configuration:
```env
DATABASE_URL=postgresql://user:password@host:5432/dbname
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-anon-key
JWT_SECRET=your-secret-key-min-32-chars
```

3. **Run the Container**:
```bash
docker run -d \
  --name couchloop-mcp \
  --env-file .env \
  -p 3000:3000 \
  docker/mcp-couchloop:latest
```

4. **Connect via MCP**:
The server exposes MCP protocol on stdio and HTTP endpoints on port 3000.

## Available Tools

### Session Management
- `create_session` - Start a new conversation session or guided journey
- `resume_session` - Continue a previously paused session
- `save_checkpoint` - Capture progress or important moments

### Communication
- `send_message` - Send messages through the therapeutic AI stack with safety checks
- `get_checkpoints` - Retrieve saved checkpoints from a session

### Journey Navigation
- `list_journeys` - View available guided experiences
- `get_journey_status` - Check progress in current journey

### Insights & Memory
- `save_insight` - Capture meaningful realizations
- `get_insights` - Retrieve user insights and patterns
- `get_user_context` - Get personalized context for tailored responses

## Use Cases

### For AI Platforms (ChatGPT, Claude)
- Enable persistent memory across chat sessions
- Add therapeutic conversation capabilities
- Implement safety guardrails for sensitive topics
- Provide structured conversation flows

### For Wellness Applications
- Mental health support chatbots
- Guided meditation and reflection tools
- Mood tracking and journaling assistants
- Crisis intervention systems

### For Developers
- Add behavioral governance to any LLM integration
- Implement session persistence for conversational AI
- Build safety layers for production AI deployments
- Create guided conversation experiences

## Configuration

### Required Environment Variables
- `DATABASE_URL`: PostgreSQL connection string
- `SUPABASE_URL`: Your Supabase project URL
- `SUPABASE_ANON_KEY`: Supabase anonymous key
- `JWT_SECRET`: Secret for JWT tokens (min 32 characters)

### Optional Environment Variables
- `SHRINK_CHAT_API_KEY`: API key for therapeutic backend integration
- `NODE_ENV`: Environment mode (development/staging/production)
- `LOG_LEVEL`: Logging verbosity (debug/info/warn/error)

## Architecture

CouchLoop implements a three-layer architecture:

1. **MCP Protocol Layer**: Handles tool/resource requests from AI agents
2. **Governance Engine**: Monitors and evaluates AI behavior
3. **Persistence Layer**: PostgreSQL database for sessions, checkpoints, and insights

The server can operate in two modes:
- **Standalone**: Direct MCP integration with your AI agent
- **Pass-through**: Routes messages through shrink-chat backend for additional therapeutic features

## Database Setup

1. **Initialize Schema**:
```bash
npm run db:push
```

2. **Seed Journey Definitions**:
```bash
npm run db:seed
```

3. **View Database** (development):
```bash
npm run db:studio
```

## Security Considerations

- All sessions are user-scoped with JWT authentication
- OAuth flow for secure user identification
- Encrypted storage for sensitive checkpoint data
- Rate limiting on API endpoints
- Input validation on all tools

## Performance

- Supports 100+ concurrent sessions
- Sub-100ms checkpoint save/retrieve
- Automatic session cleanup after 30 days
- Optimized for ChatGPT's 5-second timeout requirements

## Development

### Building from Source
```bash
git clone https://github.com/wisenbergg/couchloop-mcp.git
cd couchloop-mcp
npm install
npm run build
npm start
```

### Running Tests
```bash
npm test
```

### Development Mode
```bash
npm run dev
```

## Support

- **Documentation**: https://github.com/wisenbergg/couchloop-mcp
- **Issues**: https://github.com/wisenbergg/couchloop-mcp/issues
- **Email**: greg@couchloop.com

## License

MIT - See [LICENSE](https://github.com/wisenbergg/couchloop-mcp/blob/main/LICENSE) for details

## Contributing

Contributions are welcome! Please see our [Contributing Guide](https://github.com/wisenbergg/couchloop-mcp/blob/main/CONTRIBUTING.md) for details.

## Roadmap

- [ ] Multi-language support
- [ ] Advanced crisis detection models
- [ ] Voice conversation support
- [ ] Mobile SDK
- [ ] Real-time collaboration features
- [ ] Advanced analytics dashboard
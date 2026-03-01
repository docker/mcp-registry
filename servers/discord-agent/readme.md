# Discord Agent MCP Server

Comprehensive Discord automation and management through the Model Context Protocol. This server provides 46+ tools for AI assistants to interact with Discord servers, including messaging, channel management, roles, moderation, and more.

## Features

- **10 Messaging Tools**: Send, edit, delete messages, add reactions, manage pins
- **10 Channel Tools**: Create, modify, delete channels (text, voice, forum, stage)
- **3 Thread Tools**: Create, find, and archive forum threads
- **6 Server Management Tools**: Server info, webhooks, invites, audit logs
- **3 Member Tools**: Member info, list members, set nicknames
- **6 Role Tools**: Assign, create, modify, delete roles
- **5 Moderation Tools**: Kick, ban, timeout members
- **Resources**: Guild listings and information
- **Prompts**: Interactive moderation and announcement assistants

## Prerequisites

Before using this server, you need:

1. **Discord Bot Token**
   - Create a bot at [Discord Developer Portal](https://discord.com/developers/applications)
   - Enable these Privileged Gateway Intents:
     - Server Members Intent ✅
     - Message Content Intent ✅
   - Invite the bot to your server with appropriate permissions

2. **Required Configuration**
   - `DISCORD_TOKEN`: Your Discord bot token (required, secret)

3. **Optional Configuration**
   - `TRANSPORT_MODE`: `http` or `stdio` (default: `http`)
   - `HTTP_PORT`: Port for HTTP mode (default: `3000`)
   - `LOG_LEVEL`: Logging level (default: `info`)
   - `LOG_FORMAT`: `json` or `pretty` (default: `json`)

## Quick Start

### 1. Create Discord Bot

1. Go to [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a new application and add a bot
3. Copy the bot token
4. Enable "Server Members Intent" and "Message Content Intent"
5. Generate invite URL with bot permissions:
   - Manage Channels, Manage Roles, Manage Messages
   - Send Messages, Read Message History
   - Manage Threads, Moderate Members

### 2. Configure the Server

When prompted for configuration:
- **DISCORD_TOKEN**: Paste your bot token
- Leave other settings as default or customize as needed

### 3. Use with AI Assistant

Once running, the server exposes 46+ tools that AI assistants can use:

**Example Interactions:**
- "Send a welcome message to the #general channel"
- "List all channels in the server"
- "Create a new forum channel called 'Q&A'"
- "Give the @member role to user123"
- "Show me the last 10 messages in #announcements"

## Available Tools

### Messaging
- `send_message` - Send text messages
- `send_rich_message` - Send formatted embeds
- `send_message_with_file` - Send with attachments
- `read_messages` - Get message history
- `edit_message`, `delete_message`, `bulk_delete_messages`
- `add_reaction`, `pin_message`, `unpin_message`

### Channel Management
- `list_channels`, `get_channel_details`
- `create_text_channel`, `create_voice_channel`, `create_forum_channel`
- `create_category`, `create_stage_channel`
- `modify_channel`, `delete_channel`, `set_channel_permissions`

### Threads
- `find_threads` - Search forum threads
- `create_thread` - Create new threads
- `archive_thread` - Archive threads

### Server
- `get_server_info` - Server details
- `modify_server` - Update server settings
- `get_audit_logs` - View audit logs
- `list_webhooks`, `create_webhook`
- `get_invites`, `create_invite`

### Members
- `get_member_info` - Member details
- `list_members` - List all members
- `set_nickname` - Change nicknames

### Roles
- `assign_role`, `remove_role`
- `create_role`, `delete_role`, `modify_role`
- `list_roles`, `get_role_info`

### Moderation
- `kick_member` - Remove member (can rejoin)
- `ban_member`, `unban_member` - Ban management
- `timeout_member`, `remove_timeout` - Temporary mutes
- `get_bans` - List banned users

## Use Cases

- **Community Management**: Automate welcome messages, role assignments, channel creation
- **Moderation**: AI-assisted moderation with timeout, kick, and ban capabilities
- **Content Distribution**: Post announcements, updates, and embeds across channels
- **Server Organization**: Create and manage channels, categories, and forum threads
- **Member Support**: Answer questions, provide information, manage roles
- **Event Management**: Create stage channels, manage voice channels, send invites

## Security Notes

- **Bot Token Security**: Your Discord token is stored securely as a secret
- **Permissions**: Bot can only perform actions it has permissions for
- **Rate Limits**: Respects Discord's rate limits automatically
- **Audit Trail**: All actions are logged in Discord's audit log

## Architecture

Built with:
- **TypeScript** - Type-safe implementation
- **Discord.js v14** - Official Discord API library
- **MCP Protocol** - Standard AI assistant integration
- **Docker** - Containerized deployment
- **Structured Logging** - JSON-formatted logs for monitoring

## Support & Documentation

- **Full Documentation**: [GitHub README](https://github.com/aj-geddes/discord-agent-mcp#readme)
- **Source Code**: [GitHub Repository](https://github.com/aj-geddes/discord-agent-mcp)
- **Issues**: [GitHub Issues](https://github.com/aj-geddes/discord-agent-mcp/issues)
- **Discord.js Docs**: [Discord.js Guide](https://discordjs.guide/)
- **MCP Protocol**: [Model Context Protocol](https://modelcontextprotocol.io/)

## License

MIT License - See [LICENSE](https://github.com/aj-geddes/discord-agent-mcp/blob/main/LICENSE)

---

**Note**: This is a third-party MCP server and is not officially affiliated with Discord, Inc.

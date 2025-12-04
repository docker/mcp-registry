# Discord Agent

AI-powered Discord server management through natural language with Claude.

## Overview

Discord Agent is a production-ready MCP server that bridges Claude AI with Discord, enabling comprehensive server management through natural language. Instead of navigating Discord's interface or writing code, simply tell Claude what you want done in your Discord server.

Built with TypeScript and Discord.js, this server provides 71 tools covering every aspect of Discord server management - from basic messaging to advanced moderation, role management, scheduled events, and auto-moderation rules.

**Key Features:**
- 71 comprehensive Discord management tools
- Production-ready with robust error handling
- Type-safe with full TypeScript and Zod validation
- Structured logging with configurable levels
- Supports stdio transport for Docker deployment

## Tools

### Messaging (10 tools)

- `send_message` - Send text messages to channels
- `send_rich_message` - Send messages with rich embeds and formatting
- `send_message_with_file` - Send messages with file attachments
- `read_messages` - Retrieve message history
- `edit_message` - Edit bot messages
- `delete_message` - Delete messages
- `bulk_delete_messages` - Delete multiple messages
- `add_reaction` - Add emoji reactions
- `pin_message` - Pin important messages
- `unpin_message` - Unpin messages

### Channels (10 tools)

- `list_channels` - Get all channels organized by type
- `get_channel_details` - Detailed channel information
- `create_text_channel` - Create text channels
- `create_voice_channel` - Create voice channels
- `create_category` - Create channel categories
- `create_forum_channel` - Create forum channels
- `create_stage_channel` - Create stage channels
- `modify_channel` - Update channel settings
- `delete_channel` - Delete channels
- `set_channel_permissions` - Manage channel permissions

### Threads (3 tools)

- `find_threads` - Search and list forum threads
- `create_thread` - Create new threads
- `archive_thread` - Archive and lock threads

### Server Management (7 tools)

- `get_server_info` - Server details and statistics
- `modify_server` - Update server settings
- `get_audit_logs` - View audit log history
- `list_webhooks` - List server webhooks
- `create_webhook` - Create new webhooks
- `get_invites` - List active invites
- `create_invite` - Generate invite links

### Members (3 tools)

- `get_member_info` - Member details and roles
- `list_members` - List all server members
- `set_nickname` - Change member nicknames

### Roles (7 tools)

- `create_role` - Create new roles with permissions
- `delete_role` - Remove roles
- `modify_role` - Update role settings
- `list_roles` - List all server roles
- `get_role_info` - Detailed role information
- `assign_role` - Add role to member
- `remove_role` - Remove role from member

### Moderation (6 tools)

- `kick_member` - Remove member (can rejoin)
- `ban_member` - Ban member (cannot rejoin)
- `unban_member` - Remove ban
- `timeout_member` - Temporarily mute member
- `remove_timeout` - Remove timeout
- `get_bans` - List banned users

### Emojis (4 tools)

- `list_guild_emojis` - List custom emojis
- `create_emoji` - Upload custom emojis
- `modify_emoji` - Update emoji settings
- `delete_emoji` - Remove emojis

### Stickers (4 tools)

- `list_guild_stickers` - List custom stickers
- `create_sticker` - Upload custom stickers
- `modify_sticker` - Update sticker settings
- `delete_sticker` - Remove stickers

### Scheduled Events (6 tools)

- `list_scheduled_events` - List upcoming events
- `get_event_details` - Event information
- `create_scheduled_event` - Create voice, stage, or external events
- `modify_scheduled_event` - Update event details
- `delete_scheduled_event` - Cancel events
- `get_event_users` - List interested users

### Auto-Moderation (5 tools)

- `list_automod_rules` - List moderation rules
- `get_automod_rule` - Rule details
- `create_automod_rule` - Create keyword filters and spam detection
- `modify_automod_rule` - Update rule settings
- `delete_automod_rule` - Remove rules

### Application Commands (6 tools)

- `list_application_commands` - List slash commands
- `get_application_command` - Command details
- `create_application_command` - Create slash commands
- `modify_application_command` - Update commands
- `delete_application_command` - Remove commands
- `bulk_overwrite_commands` - Sync all commands

## Configuration

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DISCORD_TOKEN` | Yes | - | Discord bot token from [Discord Developer Portal](https://discord.com/developers/applications) |
| `TRANSPORT_MODE` | No | `stdio` | Transport mode (stdio for Docker) |
| `LOG_LEVEL` | No | `info` | Logging level: debug, info, warn, error |
| `LOG_FORMAT` | No | `json` | Log format: json or pretty |

### Secrets

| Secret | Required | Description |
|--------|----------|-------------|
| `DISCORD_TOKEN` | Yes | Discord bot token with appropriate server permissions |

### Required Bot Permissions

The bot needs appropriate Discord permissions based on which tools you use:

- **Basic**: Read Messages, Send Messages
- **Moderation**: Kick Members, Ban Members, Moderate Members
- **Management**: Manage Channels, Manage Roles, Manage Server
- **Advanced**: Administrator (for full functionality)

Create your bot at: https://discord.com/developers/applications

## Usage Examples

### Server Setup
```
"Set up a gaming community with channels for Minecraft, Valorant,
general chat, and appropriate voice channels"
```

### Messaging
```
"Send a welcome message to the #general channel with an embed showing
our server rules"
```

### Moderation
```
"Timeout user 123456789012345678 for 1 hour for spam"
```

```
"Set up auto-moderation to block common spam phrases and timeout
repeat offenders for 10 minutes"
```

### Events
```
"Create a voice event called 'Weekly Game Night' scheduled for
Saturday at 8 PM EST in the Gaming voice channel"
```

### Role Management
```
"Create a 'Moderator' role with kick, ban, and message management
permissions, then assign it to users Alice and Bob"
```

### Channel Organization
```
"Create a category called 'Voice Channels' and move all voice channels
into it, then create a new voice channel called 'AFK Channel'"
```

## Security Best Practices

1. **Never commit your Discord token** - Use environment variables or Docker secrets
2. **Use least privilege** - Only grant the bot permissions it actually needs
3. **Rotate tokens periodically** - Regenerate bot tokens every few months
4. **Monitor audit logs** - Review bot actions regularly
5. **Test in a development server** - Use a separate test server before production

## Troubleshooting

### Bot is not responding
- Verify `DISCORD_TOKEN` is set correctly
- Check bot has required permissions in Discord
- Ensure bot is invited to the server

### Permission errors
- Bot needs appropriate Discord permissions for each tool
- Check role hierarchy (bot's role must be higher than roles it manages)
- Verify channel-specific permission overrides

### Message sending fails
- Check message length (max 2000 characters)
- Verify channel exists and bot has access
- Ensure bot has Send Messages permission

### Tools not discoverable
- Server supports stdio transport for tool discovery
- No configuration required to list tools
- If issues persist, check server logs with `LOG_LEVEL=debug`

## Links

- [Source Repository](https://github.com/aj-geddes/discord-agent-mcp)
- [Documentation](https://aj-geddes.github.io/discord-agent-mcp/)
- [Discord Developer Portal](https://discord.com/developers/applications)
- [Issue Tracker](https://github.com/aj-geddes/discord-agent-mcp/issues)
- [MCP Protocol](https://modelcontextprotocol.io/)

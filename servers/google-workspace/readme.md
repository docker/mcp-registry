# Google Workspace MCP Server

[![Docker](https://img.shields.io/badge/Docker-MCP%20Registry-blue)](https://hub.docker.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Comprehensive Model Context Protocol (MCP) server providing unified access to all major Google Workspace APIs through a single interface.

## Features

### Supported Google Services

- **Gmail**: Read, send, search emails, manage labels
- **Google Drive**: List, create, delete, share files and folders
- **Google Sheets**: Create, read, update spreadsheets
- **Google Calendar**: Manage events and calendars
- **Google Docs**: Create and edit documents
- **Google Forms**: Create forms, add questions, read responses
- **Google Chat**: Send messages and manage spaces

### Key Capabilities

- **Unified Authentication**: Single OAuth 2.0 flow for all services
- **Comprehensive API Coverage**: 25+ tools across 7 Google services
- **Docker-Based**: Easy deployment with credential persistence
- **Production-Ready**: Secure credential handling and error management

## Quick Start

### Prerequisites

1. **Google Cloud Project** with enabled APIs:
   - Gmail API
   - Google Drive API
   - Google Sheets API
   - Google Calendar API
   - Google Docs API
   - Google Forms API
   - Google Chat API

2. **OAuth 2.0 Credentials**:
   - Go to [Google Cloud Console](https://console.cloud.google.com/apis/credentials)
   - Create OAuth 2.0 Client ID (Desktop application)
   - Download the JSON file

### Installation via Docker

```bash
# Pull the image
docker pull pauljeyasinghph/google-workspace-mcp:latest

# Run the container
docker run -i --rm \
  -v $(pwd)/credentials:/data/credentials \
  -v $(pwd)/client_secret.json:/data/credentials/client_secret.json:ro \
  pauljeyasinghph/google-workspace-mcp:latest
```

### Configuration in Claude Desktop

Add to `~/.config/claude/config.json`:

```json
{
  "mcpServers": {
    "google-workspace": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-v",
        "/path/to/credentials:/data/credentials",
        "-v",
        "/path/to/client_secret.json:/data/credentials/client_secret.json:ro",
        "pauljeyasinghph/google-workspace-mcp:latest"
      ]
    }
  }
}
```

## Available Tools

### Gmail Tools

- `gmail_list_messages` - List emails with filters
- `gmail_get_message` - Get specific email
- `gmail_send_message` - Send email
- `gmail_search_messages` - Search emails

### Google Chat Tools

- `chat_list_spaces` - List Chat spaces
- `chat_send_message` - Send message to space
- `chat_list_messages` - List messages in space

### Google Sheets Tools

- `sheets_create_spreadsheet` - Create new spreadsheet
- `sheets_get_values` - Read cell values
- `sheets_update_values` - Update cells
- `sheets_append_values` - Append rows

### Google Drive Tools

- `drive_list_files` - List files/folders
- `drive_create_folder` - Create folder
- `drive_delete_file` - Delete file
- `drive_share_file` - Share with users

### Google Forms Tools

- `forms_create_form` - Create new form
- `forms_add_question` - Add question
- `forms_get_responses` - Get form responses

### Google Calendar Tools

- `calendar_list_events` - List events
- `calendar_create_event` - Create event
- `calendar_delete_event` - Delete event

### Google Docs Tools

- `docs_create_document` - Create document
- `docs_get_document` - Get document content
- `docs_insert_text` - Insert text
- `docs_append_text` - Append text

## Authentication Flow

1. **First Run**: Opens browser for OAuth consent
2. **Token Storage**: Saves tokens to mounted volume
3. **Auto-Refresh**: Automatically refreshes expired tokens
4. **Persistent**: Tokens survive container restarts

## Security Best Practices

1. **Never commit credentials** to version control
2. **Use read-only mounts** for client_secret.json
3. **Restrict OAuth scopes** to minimum required
4. **Rotate credentials** periodically
5. **Monitor API usage** in Google Cloud Console

## Building from Source

```bash
# Clone the repository
git clone https://github.com/PaulJeyasinghph/google-workspace-mcp.git
cd google-workspace-mcp

# Build the Docker image
docker build -t google-workspace-mcp .

# Run locally
docker run -i --rm \
  -v $(pwd)/credentials:/data/credentials \
  -v $(pwd)/client_secret.json:/data/credentials/client_secret.json:ro \
  google-workspace-mcp
```

## Usage Examples

### Send an Email

```python
{
  "tool": "gmail_send_message",
  "arguments": {
    "to": "example@gmail.com",
    "subject": "Test Email",
    "body": "This is a test email sent via MCP."
  }
}
```

### Create a Spreadsheet

```python
{
  "tool": "sheets_create_spreadsheet",
  "arguments": {
    "title": "Q1 Budget",
    "sheet_names": ["Revenue", "Expenses"]
  }
}
```

### Schedule a Meeting

```python
{
  "tool": "calendar_create_event",
  "arguments": {
    "summary": "Team Sync",
    "start_time": "2024-03-20T10:00:00Z",
    "end_time": "2024-03-20T11:00:00Z",
    "attendees": ["team@example.com"]
  }
}
```

## Troubleshooting

### Authentication Errors

- Verify all required APIs are enabled
- Check OAuth consent screen configuration
- Ensure credentials file is properly mounted

### Permission Denied

- Confirm OAuth scopes in credentials
- Check Google Workspace admin settings
- Verify user has necessary permissions

### Container Issues

- Check volume mount paths
- Verify credentials directory permissions
- Review Docker logs for errors

## Development

### Project Structure

```
google-workspace-mcp/
├── src/
│   ├── server.py          # Main MCP server
│   ├── auth.py            # OAuth handler
│   ├── gmail.py           # Gmail tools
│   ├── chat.py            # Chat tools
│   ├── sheets.py          # Sheets tools
│   ├── drive.py           # Drive tools
│   ├── forms.py           # Forms tools
│   ├── calendar.py        # Calendar tools
│   └── docs.py            # Docs tools
├── Dockerfile             # Multi-stage build
├── requirements.txt       # Python dependencies
├── server.yaml           # Docker registry config
├── tools.json            # Tool definitions
└── README.md             # This file
```

### Running Tests

```bash
# Install dev dependencies
pip install -r requirements.txt pytest

# Run tests
pytest tests/
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/PaulJeyasinghph/google-workspace-mcp/issues)
- **Discussions**: [GitHub Discussions](https://github.com/PaulJeyasinghph/google-workspace-mcp/discussions)
- **Documentation**: [Google Workspace APIs](https://developers.google.com/workspace)

## Acknowledgments

- Built with [Model Context Protocol](https://modelcontextprotocol.io)
- Powered by [Google Workspace APIs](https://developers.google.com/workspace)
- Inspired by the MCP community

## Roadmap

- [ ] Add Google Meet integration
- [ ] Support service account authentication
- [ ] Add batch operation support
- [ ] Implement webhook listeners
- [ ] Add Google Admin SDK tools
- [ ] Support Google Workspace for Education

---

Made with ❤️ for the MCP community

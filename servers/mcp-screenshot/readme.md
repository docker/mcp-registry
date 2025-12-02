# MCP Screenshot Server - Docker MCP Registry

## Quick Start

### Pull and Run

```bash
# Pull the image
docker pull digidefiance/mcp-screenshot:latest

# Run with default configuration
docker run -i --rm \
  -v $(pwd)/screenshots:/app/screenshots \
  digidefiance/mcp-screenshot:latest
```

### Docker Compose

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  mcp-screenshot:
    image: digidefiance/mcp-screenshot:latest
    stdin_open: true
    volumes:
      - ./screenshots:/app/screenshots
      - ./config.json:/app/config/config.json:ro
    environment:
      - NODE_ENV=production
    restart: unless-stopped
```

Run:
```bash
docker-compose up -d
```

## Configuration

### Basic Configuration

The server accepts configuration via a JSON file:

```json
{
  "securityPolicy": {
    "allowedDirectories": ["/app/screenshots"],
    "maxCapturesPerMinute": 10,
    "enableAuditLog": true
  },
  "excludedWindowPatterns": [
    ".*password.*",
    ".*1Password.*",
    ".*LastPass.*"
  ]
}
```

Mount the config file:
```bash
docker run -i --rm \
  -v $(pwd)/config.json:/app/config/config.json:ro \
  -v $(pwd)/screenshots:/app/screenshots \
  digidefiance/mcp-screenshot:latest \
  node dist/cli.js --config=/app/config/config.json
```

### Environment Variables

- `NODE_ENV`: Node environment (default: production)
- `TESSDATA_PREFIX`: Path to Tesseract training data (default: /usr/share/tessdata)

## Tools

### screenshot_capture_full
Capture full screen or specific display.

**Example:**
```json
{
  "tool": "screenshot_capture_full",
  "arguments": {
    "format": "png",
    "savePath": "/app/screenshots/screenshot.png"
  }
}
```

### screenshot_capture_window
Capture specific application window.

**Example:**
```json
{
  "tool": "screenshot_capture_window",
  "arguments": {
    "windowTitle": "Chrome",
    "format": "png"
  }
}
```

### screenshot_capture_region
Capture rectangular screen region.

**Example:**
```json
{
  "tool": "screenshot_capture_region",
  "arguments": {
    "x": 0,
    "y": 0,
    "width": 800,
    "height": 600,
    "format": "png"
  }
}
```

### screenshot_list_displays
List all connected displays.

### screenshot_list_windows
List all visible windows.

## Features

- ✅ Cross-platform support (Linux, macOS, Windows)
- ✅ Multiple capture modes (full screen, window, region)
- ✅ Multi-format support (PNG, JPEG, WebP, BMP)
- ✅ PII masking with OCR
- ✅ Security policies and rate limiting
- ✅ Multi-monitor support
- ✅ Docker and Docker Compose support

## Security

The container runs as a non-root user and includes:
- Read-only root filesystem
- No new privileges
- Dropped capabilities
- Resource limits

## Volumes

- `/app/screenshots`: Directory for saved screenshots
- `/app/config`: Configuration directory

## Resource Limits

Recommended limits:
- CPU: 2 cores (limit), 0.5 cores (request)
- Memory: 2Gi (limit), 512Mi (request)

## Support

- GitHub: https://github.com/digital-defiance/ai-capabilities-suite
- Issues: https://github.com/digital-defiance/ai-capabilities-suite/issues
- Email: info@digitaldefiance.org

## License

MIT License

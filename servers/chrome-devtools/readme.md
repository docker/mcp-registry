# Chrome DevTools MCP Server

Enable AI coding agents to control and inspect live Chrome browsers through DevTools integration.

## Features

The Chrome DevTools MCP server provides **26+ tools** across six categories:

### Input Automation
- Click elements
- Type text
- Fill forms
- Interact with web pages

### Navigation
- Page management
- URL navigation
- Wait for page loads
- Browser control

### Performance
- Performance tracing
- Analysis and profiling
- Metrics collection

### Network
- Request inspection
- Response monitoring
- Network activity tracking

### Debugging
- Take screenshots
- Access console logs
- Debug web applications
- Element inspection

### Emulation
- Device simulation
- Viewport control
- Mobile testing
- Responsive design testing

## Requirements

- **Node.js**: v20.19 or newer LTS version
- **Chrome**: Current stable version or newer
- **npm**: Latest version

## Configuration Options

The server supports various configuration parameters:
- `--headless`: Run Chrome in headless mode
- `--isolated`: Use isolated browser profile
- `--channel`: Specify Chrome channel
- `--browserUrl`: Connect to existing Chrome instance
- `--wsEndpoint`: WebSocket endpoint for remote debugging

## Documentation

For more information, visit:
- [Chrome DevTools MCP GitHub](https://github.com/ChromeDevTools/chrome-devtools-mcp)
- [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)

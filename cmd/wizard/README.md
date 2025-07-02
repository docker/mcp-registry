# 🧙‍♂️ MCP Server Registry Wizard

A beautiful TUI (Terminal User Interface) wizard to help you easily add your MCP server to the Docker MCP Registry.

## ✨ Features

- **Interactive Form**: Step-by-step guided process
- **Beautiful UI**: Styled with Charm's huh library
- **Validation**: Input validation to ensure correct data
- **Category Selection**: Choose from predefined categories
- **Secrets & Environment Variables**: Configure your server's requirements
- **YAML Generation**: Automatically generates the `server.yaml` file

## 🚀 Usage

Run the wizard from the root of the mcp-registry repository:

```bash
task wizard
```

## 📋 What the Wizard Collects

### Basic Information

- **Server Name**: The name of your MCP server (no spaces allowed)
- **GitHub Repository**: The full GitHub URL of your project
- **Branch**: Optional branch name (defaults to main/master)
- **Category**: Select from predefined categories:
  - ai
  - data-visualization
  - database
  - devops
  - ecommerce
  - finance
  - games
  - communication
  - monitoring
  - productivity
  - search

### Server Details

- **Title**: A descriptive title for your server
- **Description**: Detailed description of what your server does
- **Icon URL**: Optional custom icon URL
- **Docker Image**: Optional custom Docker image (defaults to `mcp/[server-name]`)

### Configuration (Optional)

- **Secrets**: Configure secret variables like API keys, passwords
- **Environment Variables**: Configure environment variables for your server

## 📁 Output

The wizard generates a `server.yaml` file in the `servers/[server-name]/` directory that follows the MCP Registry format:

```yaml
name: my-server
image: mcp/my-server
type: server
meta:
  category: database
  tags:
    - database
about:
  title: My Awesome Server
  description: This server does amazing things
  icon: mcp/my-server
source:
  project: https://github.com/user/repo
config:
  description: Configure the connection to My Awesome Server
  secrets:
    - name: my-server.api_key
      env: API_KEY
      example: your-api-key-here
  env:
    - name: HOST
      example: localhost
      value: "{{my-server.host}}"
  parameters:
    type: object
    properties:
      host:
        type: string
```

## 🎨 UI Preview

The wizard features:

- **Colorful headers** with emojis
- **Step-by-step forms** with clear descriptions
- **Input validation** with helpful error messages
- **Confirmation dialogs** for optional configurations
- **Progress indicators** showing current step

## 🛠️ Requirements

- Go 1.24+
- Terminal with TTY support
- The `huh` library (automatically installed)

## 🔄 After Running the Wizard

1. **Review** the generated `server.yaml` file
2. **Test locally** with: `task catalog -- [server-name]`
3. **Create a pull request** to add your server to the registry

## 🤝 Integration with Existing Workflow

This wizard complements the existing `task create` command by providing a more user-friendly, interactive experience for creating MCP server configurations.

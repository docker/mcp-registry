# 🐳 Official Docker MCP Registry

Welcome to the Official Docker MCP (Model Context Protocol) Registry! This repository serves as a curated catalog of MCP servers that can be easily discovered, deployed, and integrated with any MCP Client and compatible with Docker tooling. 

Entries in this catalog will be available at: 
- [MCP catalog](https://hub.docker.com/mcp) 
- [Docker Desktop's MCP Toolkit](https://www.docker.com/products/docker-desktop/) 
- [Docker Hub `mcp` namespace](https://hub.docker.com/u/mcp) (for MCP servers built by Docker)

## 🤖 What is MCP?
The Model Context Protocol (MCP) is an open standard that enables AI assistants to securely connect with external data sources and tools. Read more at [MCP Official Documantation](https://modelcontextprotocol.io/introduction).

## ✨ Why Use the Docker MCP Registry?
- **Enterprise Security**: MCP servers built by Docker include cryptographic signatures, provenance tracking, and Software Bills of Materials (SBOMs) for maximum trust and compliance
- **Container Isolation**: All MCP servers run in isolated containers, protecting your host system from potential security vulnerabilities
- **Curated Quality**: All MCP servers undergo review to ensure they meet quality and security standards
- **Easy Discovery**: Browse and find MCP servers for your specific use cases or share yours to millions of developers using Docker tools
- **Docker Integration**: Seamless deployment with Docker containers

## 🤝 Contributing to the Docker MCP Registry
We welcome contributions to the Official Docker MCP Registry! If you'd like to contribute, you can submit a PR with the metadata information and it will be added to the [MCP catalog](https://hub.docker.com/mcp), to [Docker Desktop's MCP Toolkit](https://www.docker.com/products/docker-desktop/), and (for MCP servers images built by Docker) in `mcp` namespace in [Docker Hub](https://hub.docker.com/u/mcp).

To add your MCP server to the registry, please review the [CONTRIBUTING](CONTRIBUTING.md) guide for detailed instructions. We support two types of submissions:

### 🏗️ Option A: Docker-Built Image (Recommended)
Have Docker build and maintain your server image with enhanced security features. You'll submit the required information via pull request and upon approval Docker will build, sign, and publish your image to mcp/your-server-name on Docker Hub and the catalog entry will be available in the catalog in 24 hours.

_**Benefits: Your image will include cryptographic signatures, provenance tracking, SBOMs, and automatic security updates**_

### 📦 Option B: Self-Provided Pre-Built Image
In this option, you'll provide an already built image which will be used directly in the catalog. 

_**Note: Self-built images still benefit from container isolation but won't include the enhanced security features of Docker-built images.**_

## ✏️ Modifying or Removing Servers
To request modifications or removal of an existing MCP Server please open an issue explaining the reason for the edit/removal.

## ✅ Compliance and Quality Standards
All MCP servers in this registry must:
- Follow security best practices
- Include comprehensive documentation
- Provide working Docker deployment
- Maintain compatibility with MCP standards
- Include proper error handling and logging

_**Non-compliant servers will be reviewed and may be removed from the registry.**_

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

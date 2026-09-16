# Helidon MCP Server

A base MCP (Model Context Protocol) server implementation using Helidon framework, containerized with Docker for easy deployment and testing.

## Features

- **MCP Protocol Support**: Implements the Model Context Protocol for tool discovery and execution
- **Docker Containerized**: Ready-to-use Docker image with multi-stage build
- **Java 21**: Built with modern Java features and Helidon 4.0.7
- **Self-contained JAR**: All dependencies included in a single executable JAR
- **HTTP Server**: RESTful API endpoints for MCP protocol communication
- **JSON-RPC Support**: Full JSON-RPC 2.0 protocol implementation

## Documentation

For detailed documentation, deployment guides, and troubleshooting, visit: [https://github.com/thesurenk/helidon-mcp-server](https://github.com/thesurenk/helidon-mcp-server)

## Quick Start

1. The server runs on port 8080 by default
2. Access the MCP endpoint at `http://localhost:8080/mcp`
3. Use JSON-RPC protocol for communication

## Configuration

- **Server Port**: Configurable via `SERVER_PORT` environment variable (default: 8080)
- **Docker Image**: Built with multi-stage build for optimal size
- **Java Runtime**: Eclipse Temurin 21 JRE

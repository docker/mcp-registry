# Autotask PSA MCP Server

An MCP server that provides AI assistants with structured access to Kaseya Autotask PSA data and operations.

## Features

- **39 tools** covering companies, contacts, tickets, projects, time entries, billing items, notes, attachments, and more
- **Intelligent ID-to-name mapping** that resolves company and resource IDs to human-readable names
- **Smart caching** with 30-minute TTL for reduced API calls
- **Picklist discovery** tools for finding valid field values dynamically

## Prerequisites

You need valid Autotask API credentials:
- **Username**: Your API user email address
- **Secret**: Your API secret key
- **Integration Code**: Your Autotask integration code

To set up API credentials, create a dedicated API user in Autotask with appropriate permissions for the entities you want to access.

## Links

- [Source Repository](https://github.com/wyre-technology/autotask-mcp)
- [Documentation](https://wyre-technology.github.io/autotask-mcp)
- [Autotask REST API](https://ww3.autotask.net/help/DeveloperHelp/Content/APIs/REST/REST_API_Home.htm)

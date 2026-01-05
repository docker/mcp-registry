# PostHog MCP Server (Local)

This is the local containerized version of the PostHog MCP server. It runs in Docker on your machine.

For the remote/hosted version that connects directly to PostHog's servers, see [posthog-remote](../posthog-remote/).

The PostHog MCP server provides AI agents with access to PostHog's product analytics platform, including:

- **Analytics & Insights**: Query trends, funnels, and run HogQL queries
- **Feature Flags**: Create, update, and manage feature flags
- **Experiments**: Set up and analyze A/B tests
- **Surveys**: Create and manage user surveys
- **Error Tracking**: Access error details and investigate issues
- **Dashboards**: Create and manage analytics dashboards

## Documentation

- [PostHog MCP Documentation](https://posthog.com/docs/model-context-protocol)
- [GitHub Repository](https://github.com/PostHog/mcp)

## Authentication

You need a PostHog Personal API Key to use this MCP server. You can create one in your [PostHog account settings](https://app.posthog.com/settings/user-api-keys).

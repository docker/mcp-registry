# PostHog MCP Server (Remote)

This is the remote/hosted version of the PostHog MCP server. It connects directly to PostHog's hosted MCP endpoint without running a local container.

For the local containerized version, see [posthog](../posthog/).

## Features

- **Analytics & Insights**: Query trends, funnels, and run HogQL queries
- **Feature Flags**: Create, update, and manage feature flags
- **Experiments**: Set up and analyze A/B tests
- **Surveys**: Create and manage user surveys
- **Error Tracking**: Access error details and investigate issues
- **Dashboards**: Create and manage analytics dashboards

## Documentation

- [PostHog MCP Documentation](https://posthog.com/docs/model-context-protocol)
- [GitHub Repository](https://github.com/PostHog/posthog/tree/master/services/mcp)

## Authentication

You need a PostHog Personal API Key to use this MCP server. Create one at [PostHog API Keys](https://us.posthog.com/settings/user-api-keys) using the "MCP Server" preset.

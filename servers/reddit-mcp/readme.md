# Reddit MCP Server

Full Reddit API integration with OAuth2 authentication, enabling Claude to read and write Reddit posts, comments, and interact with subreddits.

## Documentation

For complete documentation, setup instructions, and usage examples, visit:
https://github.com/Securiteru/reddit-mcp

## Features

- **Reading Tools (6)**: Get posts, comments, search, subreddit info, user content
- **Writing Tools (7)**: Submit posts/comments, edit, delete, vote, save/unsave
- **OAuth2 Authentication**: Secure Reddit API access
- **Rate Limiting**: Built-in 60 req/min rate limiter
- **Error Handling**: Automatic retry with exponential backoff

## Configuration

Requires Reddit API credentials:
- `REDDIT_CLIENT_ID` - Reddit app client ID
- `REDDIT_CLIENT_SECRET` - Reddit app client secret
- `REDDIT_USERNAME` - Your Reddit username
- `REDDIT_PASSWORD` - Your Reddit password
- `REDDIT_USER_AGENT` - User agent string

**Note**: Reddit account must have 2FA disabled for password grant OAuth2 authentication.

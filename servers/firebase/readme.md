# Firebase MCP Server

Official Firebase MCP server for managing Firebase projects through AI assistants.

## Features

The Firebase MCP server provides 50+ tools for:

- **Project Management**: Create, configure, and manage Firebase projects
- **Firestore**: Query and manage Firestore databases
- **Authentication**: Manage user authentication and security
- **Cloud Messaging**: Send and manage push notifications
- **Crashlytics**: Debug and analyze app crashes
- **Remote Config**: Manage app configuration remotely
- **And more**: Storage, Functions, Hosting, and other Firebase services

## Prompts

The server includes helpful prompts:
- `firebase:deploy` - Deploy Firebase services
- `firebase:init` - Initialize a new Firebase project
- `firebase:consult` - Get Firebase best practices advice
- `crashlytics:connect` - Connect Crashlytics for debugging

## Authentication

The Firebase MCP server uses Firebase CLI authentication. You'll need to:

1. Have a Google Account
2. Authenticate using the `firebase_login` tool
3. Select or create a Firebase project

The server maintains the same credentials used by the Firebase CLI.

## Documentation

For more information, visit:
- [Firebase Documentation](https://firebase.google.com/docs)
- [Firebase Tools GitHub](https://github.com/firebase/firebase-tools)
- [Firebase MCP Source](https://github.com/firebase/firebase-tools/tree/master/src/mcp)

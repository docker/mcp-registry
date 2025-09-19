# SimpleCheckList MCP Server

Advanced task management system with MCP server integration and Docker-optimized deployment.

## 🚀 What's New in v1.0.1

**Docker Stability Improvements:**
- Fixed container exit issue - containers now run reliably in Docker environments
- Changed default mode from 'both' to 'backend' for optimal containerized deployment  
- Enhanced startup messaging and error handling
- Improved container lifecycle management
- Better separation of HTTP API and MCP server modes

## 📋 Features

- **Complete Task Management**: Projects → Groups → Task Lists → Tasks → Subtasks hierarchy
- **SQLite Database**: Persistent data storage with reliable database operations
- **RESTful API**: Comprehensive HTTP API for all task management operations
- **MCP Protocol**: Full Model Context Protocol compliance for AI assistant integration
- **Docker Optimized**: Stable containerized deployment with health checks
- **Multiple Modes**: Backend (API), MCP (protocol), or both simultaneously

## 🐳 Quick Docker Deployment

```bash
# Start the server (recommended for Docker)
docker run -d -p 8355:8355 mayurkakade/mcp-server:latest

# Verify it's running
curl http://localhost:8355/api/health
```

The container will start in 'backend' mode by default, providing a stable HTTP API.

## 🔧 Configuration

### Environment Variables
- `MODE`: Deployment mode (`backend`|`mcp`|`both`) - default: `backend`
- `PORT`: HTTP API port - default: `8355`
- `DB_PATH`: SQLite database path - default: `/app/data/tasks.db`
- `LOG_LEVEL`: Logging level - default: `info`

### Persistent Storage
```bash
docker run -d \
  -p 8355:8355 \
  -v /host/data:/app/data \
  mayurkakade/mcp-server:latest
```

## 🤖 MCP Integration

### Claude Desktop Configuration
```json
{
  "mcpServers": {
    "simplechecklist": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "mayurkakade/mcp-server:latest",
        "--mode=mcp"
      ]
    }
  }
}
```

### HTTP-based MCP Clients
Point your client to: `http://localhost:8355/mcp`

## 📊 API Endpoints

- `GET /api/health` - Health check
- `GET /api/projects` - List projects
- `POST /api/projects` - Create project
- `GET /api/projects/{id}/groups` - List groups
- `POST /api/groups` - Create group
- `GET /api/task-lists/{id}/tasks` - List tasks
- `POST /api/tasks` - Create task

See [DOCKER-USAGE-INSTRUCTIONS.txt](DOCKER-USAGE-INSTRUCTIONS.txt) for complete documentation.

## 🏥 Health Monitoring

```bash
# Health check
curl http://localhost:8355/api/health

# Expected response
{
  "status": "OK",
  "timestamp": "2025-09-17T12:00:00Z",
  "version": "v1.0.1",
  "mode": "backend"
}
```

## 🔍 Troubleshooting

**Container exits immediately?** 
- Fixed in v1.0.1! Update to the latest image.

**Can't connect to API?**
- Verify port mapping: `-p 8355:8355`
- Check container status: `docker ps`
- Test health endpoint

**Need help?**
- Check container logs: `docker logs <container-id>`
- Review [DOCKER-USAGE-INSTRUCTIONS.txt](DOCKER-USAGE-INSTRUCTIONS.txt)
- Open issue on [GitHub](https://github.com/DevMayur/SimpleCheckList/issues)

## 📈 Version History

**v1.0.1** (Latest):
- 🚀 Docker stability improvements
- 🔧 Enhanced container lifecycle
- 📝 Comprehensive documentation
- 🏥 Better health checks

**v1.0.0**:
- Initial Docker release
- Basic MCP functionality

## 🎯 Registry Impact

This update significantly improves the user experience for Docker deployments:
- ✅ Containers run reliably without manual intervention
- ✅ Clear documentation reduces support requests
- ✅ Production-ready deployment options
- ✅ Maintains full MCP protocol compliance

Perfect for AI assistants managing complex project workflows! 🤖✨

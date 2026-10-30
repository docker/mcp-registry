# TypeScript MCP Debugger

Enterprise-grade debugging for Node.js and TypeScript applications with 25+ specialized tools.

## Documentation

For complete documentation, visit:
- **Main Documentation**: https://github.com/digital-defiance/ai-capabilities-suite/tree/main/packages/mcp-debugger-server#readme
- **Tool Reference**: https://github.com/digital-defiance/ai-capabilities-suite/blob/main/packages/mcp-debugger-server/TOOL-REFERENCE.md
- **AI Agent Integration**: https://github.com/digital-defiance/ai-capabilities-suite/blob/main/packages/mcp-debugger-server/AI-AGENT-INTEGRATION.md
- **VS Code Integration**: https://github.com/digital-defiance/ai-capabilities-suite/blob/main/packages/mcp-debugger-server/VSCODE-INTEGRATION.md

## Features

### Core Debugging
- **Breakpoint Management**: Set, remove, toggle, and list breakpoints with conditions, hit counts, and logpoints
- **Execution Control**: Continue, step over/into/out, and pause execution
- **Variable Inspection**: Inspect local/global variables, evaluate expressions, watch variables
- **Call Stack Navigation**: View and navigate stack frames with context switching

### Advanced Features
- **Hang Detection**: Detect infinite loops and hanging processes
- **TypeScript Support**: Full source map support for TypeScript debugging
- **Performance Profiling**: CPU profiling, memory profiling, heap snapshots
- **Test Framework Integration**: Debug Jest, Mocha, and Vitest tests

### Enterprise Features
- **Observability**: Structured logging, metrics, health checks, Prometheus export
- **Security**: Authentication, rate limiting, data masking, audit logging
- **Production Ready**: Circuit breakers, retry logic, graceful shutdown, resource limits

## Quick Start

### Using Docker Desktop MCP Toolkit

1. Enable the server in Docker Desktop's MCP Toolkit
2. Configure environment variables (optional):
   - `NODE_ENV`: production (default) or development
   - `LOG_LEVEL`: info (default), debug, warn, or error
3. Start using the debugging tools in your AI agent

### Using Docker CLI

```bash
docker run -d \
  --name mcp-debugger \
  -e NODE_ENV=production \
  -e LOG_LEVEL=info \
  digidefiance/mcp-debugger-server:latest
```

### Using Docker Compose

```yaml
version: '3.8'
services:
  mcp-debugger:
    image: digidefiance/mcp-debugger-server:latest
    environment:
      - NODE_ENV=production
      - LOG_LEVEL=info
    restart: unless-stopped
```

## Available Tools

The server provides 25 debugging tools:

### Session Management
- `debugger_start` - Start a new debug session
- `debugger_stop_session` - Stop a debug session

### Breakpoints
- `debugger_set_breakpoint` - Set a breakpoint with optional condition
- `debugger_remove_breakpoint` - Remove a breakpoint
- `debugger_toggle_breakpoint` - Toggle breakpoint enabled/disabled
- `debugger_list_breakpoints` - List all breakpoints

### Execution Control
- `debugger_continue` - Resume execution
- `debugger_step_over` - Step over current line
- `debugger_step_into` - Step into function call
- `debugger_step_out` - Step out of current function
- `debugger_pause` - Pause execution

### Variable Inspection
- `debugger_inspect` - Evaluate an expression
- `debugger_get_local_variables` - Get local variables
- `debugger_get_global_variables` - Get global variables
- `debugger_inspect_object` - Inspect object properties
- `debugger_add_watch` - Add a watch expression
- `debugger_remove_watch` - Remove a watch expression
- `debugger_get_watches` - Get all watched expressions

### Call Stack
- `debugger_get_stack` - Get current call stack
- `debugger_switch_stack_frame` - Switch to a different stack frame

### Advanced Features
- `debugger_detect_hang` - Detect hanging processes and infinite loops
- `debugger_start_cpu_profile` - Start CPU profiling
- `debugger_stop_cpu_profile` - Stop CPU profiling and get results
- `debugger_take_heap_snapshot` - Take a heap snapshot
- `debugger_get_performance_metrics` - Get performance metrics

## Use Cases

### Debugging Failing Tests
```
1. Start debug session with test command
2. Set breakpoints in test code
3. Step through execution to find the issue
4. Inspect variables to understand state
5. Fix the bug and verify
```

### Performance Analysis
```
1. Start debug session with your application
2. Start CPU profiling
3. Run performance-critical code
4. Stop profiling and analyze results
5. Identify bottlenecks and optimize
```

### Detecting Infinite Loops
```
1. Use debugger_detect_hang with your code
2. Configure timeout and sample interval
3. Get hang location and call stack
4. Fix the infinite loop
```

### TypeScript Debugging
```
1. Start debug session with TypeScript code
2. Set breakpoints in .ts files (not .js)
3. Source maps automatically map locations
4. Inspect variables with TypeScript names
5. Debug with original source code
```

## System Requirements

- **Node.js**: >= 18.0.0
- **Docker**: Latest version recommended
- **Memory**: Minimum 512MB, recommended 1GB+
- **CPU**: x64 or arm64 architecture

## Support

- **GitHub Issues**: https://github.com/digital-defiance/ai-capabilities-suite/issues
- **Email**: info@digitaldefiance.org
- **Documentation**: https://github.com/digital-defiance/ai-capabilities-suite/tree/main/packages/mcp-debugger-server

## License

MIT License - See [LICENSE](https://github.com/digital-defiance/ai-capabilities-suite/blob/main/LICENSE) for details.

## Author

**Digital Defiance**
- Website: https://digitaldefiance.org
- GitHub: https://github.com/digital-defiance
- Email: info@digitaldefiance.org

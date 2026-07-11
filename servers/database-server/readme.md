# MCP Database Server

Comprehensive database server supporting PostgreSQL, MySQL, and SQLite with natural language SQL query capabilities. Enables AI agents to interact with databases through both direct SQL and natural language queries.

## Features

- **Multi-Database Support**: Works with SQLite, PostgreSQL, and MySQL
- **Natural Language Queries**: Convert natural language to SQL automatically
- **Direct SQL Execution**: Execute raw SQL with safety checks
- **Schema Inspection**: List tables and describe table structures
- **FastAPI Web Interface**: RESTful API endpoints for web integration
- **Docker Ready**: Fully containerized with health checks

## Available MCP Tools

### `query_database`
**Description**: Execute natural language queries against the database and convert them to SQL

**Usage**: Ask questions like "Show me all users created in the last week" or "Find the top 10 products by sales"

**Arguments**:
- `query` (string): Natural language description of what you want to query from the database

### `list_tables`
**Description**: List all available tables in the current database

**Usage**: Ask "What tables are available?" or "Show me the database structure"

### `describe_table`
**Description**: Get detailed schema information for a specific table including columns, types, and constraints

**Usage**: Ask "Describe the users table" or "What columns does the products table have?"

**Arguments**:
- `table_name` (string): Name of the table to describe

### `execute_sql`
**Description**: Execute raw SQL queries with safety checks and validation

**Usage**: Execute specific SQL commands when you need precise control

**Arguments**:
- `query` (string): SQL query to execute

### `connect_to_database`
**Description**: Connect to a new database using provided connection details

**Usage**: Switch between different databases during your session

**Arguments**:
- `connection_string` (string): Database connection string

### `get_connection_examples`
**Description**: Get example connection strings for different database types (SQLite, PostgreSQL, MySQL)

**Usage**: Ask "How do I connect to a PostgreSQL database?" or "Show me connection examples"

### `get_current_database_info`
**Description**: Get information about the currently connected database including type, version, and connection status

**Usage**: Ask "What database am I connected to?" or "Show me current connection details"

## Example Queries

Here are some example questions you can ask your agent:

```python
# Basic queries
"Show me all users in the database"
"What tables are available?"
"Describe the products table structure"

# Complex analysis
"Find the top 5 customers by total order value"
"Show me users who registered in the last month"
"What's the average product price by category?"

# Data exploration
"How many orders were placed yesterday?"
"List all products that are out of stock"
"Find duplicate email addresses in the users table"
```

## Configuration

The server requires a `DATABASE_URL` environment variable with your database connection string:

### SQLite (recommended for testing)
```
DATABASE_URL=sqlite+aiosqlite:///data/mydb.db
```

### PostgreSQL
```
DATABASE_URL=postgresql+asyncpg://user:password@localhost:5432/mydb
```

### MySQL
```
DATABASE_URL=mysql+aiomysql://user:password@localhost:3306/mydb
```

## Docker Usage

```bash
# Run with SQLite (simplest setup)
docker run -d \
  -p 3000:3000 \
  -e DATABASE_URL=sqlite+aiosqlite:///data/mydb.db \
  souhardyak/mcp-db-server

# Run with PostgreSQL
docker run -d \
  -p 3000:3000 \
  -e DATABASE_URL=postgresql+asyncpg://user:password@host:5432/db \
  souhardyak/mcp-db-server

# Run with persistent SQLite storage
docker run -d \
  -p 3000:3000 \
  -v /host/data:/data \
  -e DATABASE_URL=sqlite+aiosqlite:///data/mydb.db \
  souhardyak/mcp-db-server
```

## Health Check

The server provides a health endpoint at `/health` that returns database connectivity status.

## Source Code

Full source code and documentation available at: https://github.com/Souhar-dya/mcp-db-server
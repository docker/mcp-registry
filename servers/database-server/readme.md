# MCP Database Server - Comprehensive Database Management Platform

Professional-grade database server supporting PostgreSQL, MySQL, and SQLite with complete CRUD operations and natural language SQL capabilities. Designed for AI agents requiring full database lifecycle management.

## Core Features

### 🔍 **Query & Analysis**

- **Natural Language Processing**: Convert plain English to optimized SQL
- **Advanced Querying**: Complex joins, aggregations, and data analysis
- **Schema Inspection**: Complete database structure examination
- **Multi-Database Support**: SQLite, PostgreSQL, MySQL compatibility

### 🛠️ **Database Management**

- **Schema Operations**: Create, modify, and manage table structures
- **Data Manipulation**: Insert, update, and delete operations with safeguards
- **Connection Management**: Dynamic database switching and configuration
- **Health Monitoring**: Comprehensive status reporting and diagnostics

### 🔒 **Safety & Professional Standards**

- **Operation Warnings**: Clear notifications for destructive operations
- **Error Handling**: Detailed error reporting and recovery guidance
- **Success Confirmation**: Comprehensive feedback with affected row counts
- **Validation**: Input validation and query safety checks

## Available MCP Tools

### Core Query Operations

#### `query_database`

Transform natural language into SQL and execute queries safely.

- **Input**: Natural language question
- **Example**: "Show me all users who registered this month"
- **Output**: Structured query results with metadata

#### `list_tables`

Discover all available tables in the connected database.

- **Usage**: Database exploration and structure analysis
- **Output**: Complete table listing with metadata

#### `describe_table`

Get comprehensive schema information for any table.

- **Input**: Table name
- **Output**: Columns, types, constraints, indexes, and relationships

#### `execute_sql`

Execute SELECT queries with built-in safety validation.

- **Input**: SQL SELECT statement
- **Safety**: Restricted to read-only operations
- **Output**: Query results with execution metadata

### Connection Management

#### `connect_to_database`

Switch between different databases dynamically.

- **Input**: Database connection string
- **Support**: SQLite, PostgreSQL, MySQL
- **Output**: Connection status and database information

#### `get_connection_examples`

Retrieve properly formatted connection strings for all supported databases.

- **Output**: Template connection strings with explanations
- **Databases**: SQLite, PostgreSQL, MySQL examples

#### `get_current_database_info`

Get detailed information about the active database connection.

- **Output**: Database type, version, connection status, and capabilities

### Advanced Database Operations

#### `execute_unsafe_sql`

Execute any SQL operation including DDL and DML commands.

- **Input**: Any SQL statement (CREATE, INSERT, UPDATE, DELETE, DROP)
- **Warning**: Clear notification of potential data modification
- **Output**: Execution results with affected row counts
- **Use Cases**: Schema migrations, bulk operations, administrative tasks

#### `create_table`

Create new tables with custom schema definitions.

- **Inputs**: Table name, column definitions
- **Validation**: Schema validation and conflict checking
- **Output**: Creation confirmation and table structure

#### `insert_data`

Add new records to existing tables.

- **Inputs**: Table name, columns, values
- **Validation**: Data type checking and constraint validation
- **Output**: Insertion confirmation with row count

#### `update_data`

Modify existing records with conditional updates.

- **Inputs**: Table name, SET clause, WHERE condition
- **Safety**: Requires WHERE clause for targeted updates
- **Output**: Update confirmation with affected row count

#### `delete_data`

Remove records from tables with safety confirmations.

- **Inputs**: Table name, WHERE condition
- **Safety**: Clear warnings for destructive operations
- **Output**: Deletion confirmation with affected row count

## Usage Examples

### Basic Query Operations

```python
# Natural language querying
"Show me the top 10 customers by total purchase amount"
"Find all products that are currently out of stock"
"List users who haven't logged in for 30 days"

# Schema exploration
"What tables are available in this database?"
"Describe the structure of the orders table"
"Show me the current database connection details"
```

### Database Administration

```python
# Table creation
create_table("user_sessions", "id INTEGER PRIMARY KEY, user_id INTEGER, session_token TEXT, created_at TIMESTAMP")

# Data manipulation
insert_data("products", "name, price, category", "'Premium Widget', 99.99, 'Electronics'")
update_data("inventory", "quantity = quantity - 1", "product_id = 12345")
delete_data("expired_sessions", "created_at < '2023-01-01'")

# Advanced operations
execute_unsafe_sql("CREATE INDEX idx_user_email ON users(email)")
```

## Configuration

### Database Connection

Set the `DATABASE_URL` environment variable with your database connection string:

#### SQLite (Recommended for Development)

```bash
DATABASE_URL=sqlite+aiosqlite:///data/mydb.db
```

#### PostgreSQL (Production Ready)

```bash
DATABASE_URL=postgresql+asyncpg://username:password@localhost:5432/database_name
```

#### MySQL (Enterprise Support)

```bash
DATABASE_URL=mysql+aiomysql://username:password@localhost:3306/database_name
```

## Docker Deployment

### Quick Start

```bash
# Development with SQLite
docker run -d \
  -p 3000:3000 \
  -e DATABASE_URL=sqlite+aiosqlite:///data/mydb.db \
  souhardyak/mcp-db-server:latest

# Production with PostgreSQL
docker run -d \
  -p 3000:3000 \
  -e DATABASE_URL=postgresql+asyncpg://user:pass@db-host:5432/db \
  souhardyak/mcp-db-server:latest
```

### Persistent Storage

```bash
# SQLite with persistent data
docker run -d \
  -p 3000:3000 \
  -v /host/data:/data \
  -e DATABASE_URL=sqlite+aiosqlite:///data/mydb.db \
  souhardyak/mcp-db-server:latest
```

## Health Monitoring

The server provides comprehensive health endpoints:

- **Endpoint**: `GET /health`
- **Response**: Database connectivity, version info, and operational status
- **Monitoring**: Ready for production health checks and monitoring systems

## Architecture & Performance

### Technical Stack

- **Framework**: FastAPI for high-performance API endpoints
- **Database Drivers**:
  - `asyncpg` for PostgreSQL (high performance)
  - `aiomysql` for MySQL (async support)
  - `aiosqlite` for SQLite (lightweight development)
- **AI Processing**: Advanced natural language to SQL conversion
- **Containerization**: Optimized Docker image for production deployment

### Scalability

- **Async Architecture**: Non-blocking database operations
- **Connection Pooling**: Efficient database connection management
- **Multi-Database**: Simultaneous support for multiple database types
- **Resource Optimization**: Minimal memory footprint and fast response times

## Security & Best Practices

### Operation Safety

- **Query Validation**: SQL injection prevention and input sanitization
- **Operation Warnings**: Clear notifications for destructive operations
- **Error Handling**: Comprehensive error reporting without exposing sensitive data
- **Access Control**: Environment-based database access configuration

### Production Deployment

- **Docker Security**: Non-root container execution
- **Environment Variables**: Secure configuration management
- **Health Monitoring**: Production-ready monitoring and alerting
- **Documentation**: Complete deployment and operational guides

## Source Code & Support

**Repository**: [https://github.com/Souhar-dya/mcp-db-server](https://github.com/Souhar-dya/mcp-db-server)  
**Docker Image**: `souhardyak/mcp-db-server:latest`  
**Documentation**: Complete API documentation and deployment guides  
**Support**: GitHub Issues and comprehensive documentation

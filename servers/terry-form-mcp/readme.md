# Terry-Form MCP

Enterprise Terraform operations with LSP integration for intelligent Infrastructure as Code development.

## Overview

Terry-Form MCP is a comprehensive Model Control Protocol (MCP) server that enables AI assistants to manage Terraform infrastructure with enterprise-grade capabilities. It combines secure Terraform execution with Language Server Protocol (LSP) integration to provide intelligent code assistance, validation, and automation for Infrastructure as Code workflows.

**Key Features:**
- **25 MCP Tools** for complete Terraform workflow automation
- **LSP Integration** with terraform-ls for intelligent code completion and diagnostics
- **Security Scanning** with built-in vulnerability detection and CIS benchmark checks
- **Terraform Cloud Integration** for workspace and state management
- **GitHub Integration** for automated repository workflows
- **Rate Limiting & RBAC** with authentication and role-based access control
- **Docker Isolation** for safe execution in containerized environments

Built for enterprise teams requiring automated, secure, and intelligent Terraform operations integrated directly into AI-powered development workflows.

## Tools

### Core Terraform Execution

#### `terry`
Execute Terraform operations (init, validate, fmt, plan, show, graph, providers, version) with comprehensive output and security validation. The primary tool for running Terraform commands in a safe, containerized environment.

**Parameters:**
- `path` (string, required): Workspace path relative to /mnt/workspace
- `actions` (array, optional): Array of Terraform actions to execute (default: ["plan"])
- `vars` (object, optional): Terraform variables for plan action

### Workspace Management

#### `terry_workspace_list`
List all available Terraform workspaces with initialization status, state file presence, provider information, and last modified timestamps.

#### `terry_workspace_info`
Analyze Terraform workspace structure and provide recommendations. Shows file structure, configuration status, LSP readiness, and recommended next actions.

**Parameters:**
- `path` (string, optional): Workspace path relative to /mnt/workspace (default: ".")

#### `terry_workspace_setup`
Create a properly structured Terraform workspace ready for LSP operations. Sets up basic files (main.tf, variables.tf, outputs.tf) with best practice templates.

**Parameters:**
- `path` (string, required): Workspace path relative to /mnt/workspace
- `project_name` (string, optional): Name of the project (default: "terraform-project")

### Environment & Diagnostics

#### `terry_version`
Get Terraform version information and provider selections for troubleshooting compatibility issues.

#### `terry_environment_check`
Comprehensive environment check for Terraform and LSP integration. Checks container environment, available tools, paths, and configuration.

#### `terry_lsp_debug`
Debug terraform-ls functionality and LSP client state for troubleshooting LSP integration issues.

#### `terry_file_check`
Check Terraform file syntax and readiness for LSP operations. Validates file exists, is readable, and has basic Terraform syntax.

**Parameters:**
- `file_path` (string, required): File path relative to /mnt/workspace

### LSP-Powered Intelligence

#### `terraform_validate_lsp`
Validate a Terraform file using terraform-ls Language Server. Provides detailed diagnostics including errors, warnings, and syntax validation with line numbers.

**Parameters:**
- `file_path` (string, required): Path to Terraform file relative to workspace
- `workspace_path` (string, optional): Workspace directory (defaults to parent directory of file)

#### `terraform_hover`
Get documentation and information for Terraform resource at cursor position. Returns detailed descriptions, argument specifications, and attribute documentation.

**Parameters:**
- `file_path` (string, required): Path to Terraform file
- `line` (integer, required): Line number (0-based)
- `character` (integer, required): Character position (0-based)
- `workspace_path` (string, optional): Workspace directory

#### `terraform_complete`
Get completion suggestions for Terraform code at cursor position. Returns intelligent completions for resource types, arguments, attribute references, and module calls.

**Parameters:**
- `file_path` (string, required): Path to Terraform file
- `line` (integer, required): Line number (0-based)
- `character` (integer, required): Character position (0-based)
- `workspace_path` (string, optional): Workspace directory

#### `terraform_format_lsp`
Format a Terraform file using terraform-ls. Applies canonical Terraform formatting according to HashiCorp standards.

**Parameters:**
- `file_path` (string, required): Path to Terraform file
- `workspace_path` (string, optional): Workspace directory

#### `terraform_lsp_status`
Get the status of the terraform-ls Language Server integration for troubleshooting.

#### `terry_lsp_init`
Manually initialize LSP client for a specific workspace. Useful for troubleshooting or switching workspaces.

**Parameters:**
- `workspace_path` (string, required): Workspace path relative to /mnt/workspace

### Security & Analysis

#### `terry_analyze`
Analyze Terraform configuration for best practices. Scans for missing descriptions, hardcoded values, missing tags, and provides a quality score with actionable recommendations.

**Parameters:**
- `path` (string, required): Workspace path relative to /mnt/workspace

#### `terry_security_scan`
Run security scan on Terraform configuration. Detects vulnerabilities like public S3 buckets, unencrypted resources, open security groups, and IAM policy issues with CIS benchmark IDs and remediation guidance.

**Parameters:**
- `path` (string, required): Workspace path relative to /mnt/workspace
- `severity` (string, optional): Minimum severity to report (low/medium/high/critical, default: "medium")

#### `terry_recommendations`
Get recommendations for Terraform configuration improvement. Provides focus-specific guidance with impact and effort estimates.

**Parameters:**
- `path` (string, required): Workspace path relative to /mnt/workspace
- `focus` (string, optional): Area of focus - cost, security, performance, reliability (default: "security")

### Terraform Cloud Integration

#### `tf_cloud_list_workspaces`
List Terraform Cloud workspaces for an organization. Returns workspace metadata including environment, Terraform version, current run status, and resource counts.

**Parameters:**
- `organization` (string, required): Terraform Cloud organization name
- `limit` (integer, optional): Maximum workspaces to return (1-100, default: 20)

#### `tf_cloud_get_workspace`
Get detailed information about a specific Terraform Cloud workspace. Returns configuration, VCS integration, state version, execution mode, and resource count.

**Parameters:**
- `organization` (string, required): Terraform Cloud organization name
- `workspace` (string, required): Workspace name

#### `tf_cloud_list_runs`
List runs for a Terraform Cloud workspace. Returns run history with status, timestamps, resource changes, and cost estimation data.

**Parameters:**
- `organization` (string, required): Terraform Cloud organization name
- `workspace` (string, required): Workspace name
- `limit` (integer, optional): Maximum runs to return (1-100, default: 10)

#### `tf_cloud_get_state_outputs`
Get state outputs from a Terraform Cloud workspace. Returns current state outputs with values, types, and sensitivity markers (sensitive values are masked).

**Parameters:**
- `organization` (string, required): Terraform Cloud organization name
- `workspace` (string, required): Workspace name

### GitHub Integration

#### `github_clone_repo`
Clone or update a GitHub repository into the workspace. Supports branch selection and force updates.

**Parameters:**
- `owner` (string, required): Repository owner/organization
- `repo` (string, required): Repository name
- `branch` (string, optional): Branch to clone
- `force` (boolean, optional): Force update if exists (default: false)

#### `github_list_terraform_files`
List Terraform files in a GitHub repository. Supports path filtering and custom file patterns.

**Parameters:**
- `owner` (string, required): Repository owner/organization
- `repo` (string, required): Repository name
- `path` (string, optional): Subdirectory path (default: "")
- `pattern` (string, optional): File pattern to match (default: "*.tf")

#### `github_get_terraform_config`
Analyze Terraform configuration in a GitHub repository. Returns detailed analysis including providers, modules, resources, variables, and outputs.

**Parameters:**
- `owner` (string, required): Repository owner/organization
- `repo` (string, required): Repository name
- `config_path` (string, required): Path to Terraform configuration in repository

#### `github_prepare_workspace`
Prepare a Terraform workspace from a GitHub repository. Clones repository, extracts configuration, and sets up workspace ready for execution.

**Parameters:**
- `owner` (string, required): Repository owner/organization
- `repo` (string, required): Repository name
- `config_path` (string, required): Path to Terraform configuration
- `workspace_name` (string, optional): Custom workspace name

## Configuration

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `TERRY_FORM_API_KEY` | No | - | Optional API key for authentication and role-based access control |
| `TF_CLOUD_TOKEN` | No | - | Terraform Cloud API token for workspace and state management |
| `GITHUB_APP_ID` | No | - | GitHub App ID for repository integration and automation |
| `GITHUB_APP_PRIVATE_KEY` | No | - | GitHub App private key (base64 encoded) for authentication |

### Volume Mounts

| Volume | Path | Description |
|--------|------|-------------|
| `workspace` | `/mnt/workspace` | Terraform workspace directory containing configurations and state files |

## Usage Examples

### Basic Terraform Operations

```
"Run terraform plan on my infrastructure in the aws-prod workspace"
```

The AI will use the `terry` tool to execute:
```json
{
  "path": "aws-prod",
  "actions": ["init", "plan"]
}
```

### Security Analysis

```
"Scan my Terraform configuration for security vulnerabilities"
```

Uses `terry_security_scan`:
```json
{
  "path": "aws-prod",
  "severity": "high"
}
```

### LSP-Powered Validation

```
"Validate the main.tf file and show any errors"
```

Uses `terraform_validate_lsp` to provide detailed diagnostics with line numbers and character positions.

### Workspace Setup

```
"Set up a new Terraform workspace for my API project"
```

Uses `terry_workspace_setup`:
```json
{
  "path": "api-infrastructure",
  "project_name": "api-project"
}
```

### Terraform Cloud Integration

```
"List all workspaces in my production organization"
```

Uses `tf_cloud_list_workspaces`:
```json
{
  "organization": "my-org",
  "limit": 50
}
```

### GitHub Automation

```
"Clone the terraform-modules repository and prepare it for execution"
```

Combines `github_clone_repo` and `github_prepare_workspace` for automated repository workflows.

## Security Features

- **Rate Limiting**: Per-tool category rate limits (20-100 req/min) to prevent abuse
- **Role-Based Access Control**: Admin, user, and read-only roles with granular permissions
- **Request Validation**: Comprehensive input validation and path sanitization
- **Audit Logging**: Complete audit trail of all tool invocations
- **Container Isolation**: All operations run in isolated Docker containers
- **Sensitive Data Protection**: Automatic masking of sensitive outputs

## Links

- [Source Repository](https://github.com/aj-geddes/terry-form-mcp)
- [Documentation Site](https://aj-geddes.github.io/terry-form-mcp)
- [Issue Tracker](https://github.com/aj-geddes/terry-form-mcp/issues)
- [Docker Hub](https://hub.docker.com/r/ajgeddes/terry-form-mcp)

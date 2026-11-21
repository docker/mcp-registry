# 🤝 Contributing to Docker MCP Registry

Thank you for your interest in contributing to the official Docker MCP Registry.
This document outlines how to contribute to this project.

## 📦 Types of MCP Servers

There are two types of MCP servers you can add to the registry:

### 🏠 Local Servers (Containerized)
Local servers run in Docker containers on your machine. They:
- Require a Dockerfile in the source repository
- Are built and hosted as Docker images
- Run locally with full container isolation
- Can benefit from Docker-built images with enhanced security features (signatures, provenance, SBOMs, automatic updates)

### 🌐 Remote Servers (Hosted)
Remote servers are hosted externally and accessed via HTTP(S). They:
- Don't require a Dockerfile (already deployed somewhere)
- Use `streamable-http` or `sse` transport protocols
- Often require OAuth authentication
- Have dynamic tool discovery

## Add server entry with Claude Code
Let Claude Code help you add a server entry by running `cat add_mcp_server.md | claude`
If you prefer to do things manually, follow the steps below instead.

## Prerequisites

- Go v1.24+
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Task](https://taskfile.dev/)

**If you're adding a remote server,** skip to the [Adding a Remote MCP Server](#adding-a-remote-mcp-server) section below.

## 🔄 Pull request process overview

- Make sure that the license of your MCP Server allows people to consume it. (MIT or Apache 2 are great, GPL is not).
- Fork the repository to your own GitHub account and clone it locally.
- Repository includes a `servers` folder where you should add a new folder with a `server.yaml` inside.
- Repository includes a `cmd` folder with Go code to automate some of the steps.
- Open a PR by ensuring the title and its description reflect the content of the PR.
- Ensure that CI passes, if it fails, fix the failures.
- Every pull request requires a review from the Docker team before merging. [Share test credentials using this form](https://forms.gle/6Lw3nsvu2d6nFg8e6).
- Once approved, all of your commits will be squashed into a single commit with your PR title.

## 🏠 Adding a Local MCP Server

### 1️⃣ Fork this repository

Fork the repository to your own GitHub account and clone it locally.

### 2️⃣ Add your entry locally

#### 🚀 Generate your server configuration using `task wizard`

```
task wizard
```

Using the wizard it's the easiest way to create your `server.yaml`, you first need to provide a valid github repo with a Dockerfile, which the wizard will analyze to populate the server default values (you can overwrite them directly in the wizard if you need to).

The wizard allows you to add environment variables, secrets and volumes.

#### 🚀 Alternatively: Generate your server configuration using `task create`

You can use our command to automate the creation of the files. Let's assume we have a new MCP Server to access my org's database. My server's GitHub repo is located at: `https://github.com/myorg/my-orgdb-mcp`

You can call the creation tool passing the category (required), and github url. If your server requires any environment variables, pass them at the end with `-e KEY=value`.

```
task create -- --category database https://github.com/myorg/my-orgdb-mcp -e API_TOKEN=test
```

This will build an image using the Dockerfile at the root of the repository, run it while verifying the MCP server is able to list tools, and then create the necessary files. It will create a directory under `servers` as follows: `./servers/my-orgdb-mcp` and inside you will find a `server.yaml` file with your MCP definition.

```
name: my-orgdb-mcp
image: mcp/my-orgdb-mcp
type: server
meta:
  category: database
  tags:
    - database
about:
  title: My OrgDB MCP (TODO)
  description: TODO (only to provide a better description than the upstream project)
  icon: https://avatars.githubusercontent.com/u/182288589?s=200&v=4
source:
  project: https://github.com/myorg/my-orgdb-mcp
  commit: 0123456789abcdef0123456789abcdef01234567
config:
  description: Configure the connection to TODO
  secrets:
    - name: my-orgdb-mcp.api_token
      env: API_TOKEN
      example: <API_TOKEN>
```

Remember that you need to specify all the env vars that you want to use in your server:

```
task create -- --category database https://github.com/myorg/my-orgdb-mcp -e API_TOKEN=test -e MY_ORG=my-org
```

If you don't specify all the environment variables, users will not be able to configure them properly in the UI.

Also, it's important to notice that env vars and secrets are handled differently. This is how a config block looks:

```
config:
  description: Configure the connection to AWS
  secrets:
    - name: tigris.aws_secret_access_key
      env: AWS_SECRET_ACCESS_KEY
      example: YOUR_SECRET_ACCESS_KEY_HERE
  env:
    - name: AWS_ACCESS_KEY_ID
      example: YOUR_ACCESS_KEY_HERE
      value: '{{tigris.aws_access_key_id}}'
    - name: AWS_ENDPOINT_URL_S3
      example: https://fly.storage.tigris.dev
      value: '{{tigris.aws_endpoint_url_s3}}'
  parameters:
    type: object
    properties:
      aws_access_key_id:
        type: string
    required:
      - aws_access_key_id

```

This configuration will provide the following UI:

![UI Config Block](assets/img/config-ui.png)

If you want to provide a specific Docker image built by your organisation instead of having Docker build the image, you can specify it with the `--image` flag:

```
task create -- --category database --image myorg/my-mcp https://github.com/myorg/my-orgdb-mcp -e API_TOKEN=test
```

🔒 If you don't provide a Docker image, we will build the image for you and host it in [Docker Hub's `mcp` namespace](https://hub.docker.com/u/mcp), the benefits are: image will include cryptographic signatures, provenance tracking, SBOMs, and automatic security updates. Otherwise, self-built images still benefit from container isolation but won't include the enhanced security features of Docker-built images.

### 3️⃣ Run & Test your MCP Server locally

After creating your server file with `task create`, you will be given instructions for running it locally. In the case of my-orgdb-mcp, we would run the following commands next.

```
task build -- --tools my-orgdb-mcp # Not needed if providing your own image
task catalog -- my-orgdb-mcp
docker mcp catalog import $PWD/catalogs/my-orgdb-mcp/catalog.yaml
```

Now, if we go into the MCP Toolkit on Docker Desktop, we'll see our new MCP server there! We can configure and enable it there, and test it against configured clients. Once we're done testing, we can restore it back to the original Docker catalog.

```
docker mcp catalog reset
```

### Avoiding `build --tools` failures

If your MCP server needs to be configured before listing tools, you can now provide a `tools.json` file and the build process will not try to run
the server and list the tools. This is one of the most common issues that block your PR.

This is an example of a `tools.json` file:

```
[
  {
    "name": "tools_name",
    "description": "description of what you tool does"
    "arguments": [
      {
        "name": "name_of_the_argument",
        "type": "string",
        "desc": ""
      }
    ]
  },
  {
    "name": "another_tool",
    "description": "description of what another tool"
    "arguments": [
      {
        "name": "name_of_the_argument",
        "type": "string",
        "desc": ""
      }
    ]
  }
]
```

When this file is found next to your `server.yaml`, the `task build -- --tools your-server-name` lists the tools by reading the file instead of
running the server.

### 4️⃣ Wait for review and approval

Once your PR is approved and merged:

1. **For Docker-built images (`image: mcp/...`)**: Docker's infrastructure will automatically:
   - Build your image from the pinned commit
   - Sign it with cryptographic signatures
   - Generate provenance and SBOM
   - Publish to Docker Hub's `mcp/` namespace

2. **For custom images**: Your image will be referenced as-is from your registry

**Timeline**: Your entry will be processed and available within 24 hours at:

- [MCP catalog](https://hub.docker.com/mcp)
- [Docker Desktop's MCP Toolkit](https://www.docker.com/products/docker-desktop/)
- [Docker Hub `mcp` namespace](https://hub.docker.com/u/mcp) (for MCP servers built by Docker)

**Future Updates**: Once your server is in the registry, see [Updating Your MCP Server](#updating-your-mcp-server) section below for how to update it to new versions.

---

## 🔄 Updating Your MCP Server

### Updating the Commit Hash/Pin

After your MCP server is accepted into the registry, you have two options for updating it to a new version:

#### Option A: Automated Updates (Recommended for Docker-Built Images)

For servers using Docker-built images (those with `image: mcp/your-server-name`), the registry includes an automated system that keeps your server up to date:

1. **Automated Pin Updates**: A nightly GitHub Action (runs at 5 AM UTC) automatically checks all Docker-built servers for new commits on their configured branch
2. **Automatic PR Creation**: When a newer commit is detected, the bot creates a pull request updating the `source.commit` field in your `server.yaml`
3. **PR Review and Merge**: The PR is created automatically, but review and merging are performed manually by the Docker team
4. **Image Build**: Once the PR is merged to `main`, Docker's infrastructure automatically:
   - Detects the change to `server.yaml`
   - Builds the Docker image from the new pinned commit
   - Signs the image with cryptographic signatures
   - Generates provenance and SBOM (Software Bill of Materials)
   - Publishes to Docker Hub's `mcp/` namespace
5. **Availability**: The new image is typically available within **24 hours** of the PR being merged

**What you need to do**: Nothing! Just push new commits to your repository's default branch (usually `main`). The automated system will handle the rest.

**Note**: Automated updates only work for servers that:
- Use Docker-built images (`image` starts with `mcp/`)
- Are hosted on GitHub
- Have a pinned commit (`source.commit` field)
- Have a configured branch (defaults to `main`)

#### Option B: Manual Update via PR

If you want to update immediately or your server doesn't qualify for automated updates, you can manually update the commit hash:

1. **Update your fork**: Make changes to `servers/your-server-name/server.yaml`
2. **Change the commit hash**: Update the `source.commit` field to your new commit SHA

   To get the latest commit hash from your repository, you can use:

   ```bash
   git ls-remote https://github.com/<username/orgname>/mcp-repo-name.git HEAD | awk '{print $1}'
   ```

   Replace `<username/orgname>` and `mcp-repo-name` with your actual repository details.

3. **Create a PR**: Open a pull request with your changes
4. **Wait for approval**: The Docker team will review and merge your PR
5. **Image build**: After merge, Docker's infrastructure will build and publish the new image (same process as Option A)

### Timeline Summary

- **Automated pin update PR created**: Daily at ~5 AM UTC (when new commits are detected)
- **PR review and merge**: Varies (typically within a few days)
- **Image build triggered**: Immediately after PR merge (detected by Docker's infrastructure)
- **Image available in Docker Hub**: Within 24 hours of PR merge

### Troubleshooting

**Q: My server uses a custom image (not `mcp/` namespace). Can I still get automated updates?**
A: No, automated updates only work for Docker-built images. You'll need to manually update the commit hash via PR, or update your server.yaml to use the `mcp/` namespace instead.

**Q: How do I check if my pin update PR was created?**
A: Check for PRs with the title pattern `chore: update pin for {your-server-name}` or branches named `automation/update-pin-{your-server-name}`.

**Q: The automated system hasn't created a PR yet. What should I do?**
A: The system runs daily. If you need an immediate update, you can manually create a PR to update the commit hash.

---

## 🌐 Adding a Remote MCP Server

Remote MCP servers are already hosted externally and don't require Docker image building. They communicate via HTTP(S) protocols.

### Prerequisites for Remote Servers

- A publicly accessible MCP server endpoint (e.g., `https://mcp.example.com/mcp`)
- Knowledge of the transport protocol (`streamable-http` or `sse`)
- A documentation URL for your server
- OAuth configuration details (if authentication is required)

### 1️⃣ Fork this repository

Fork the repository to your own GitHub account and clone it locally.

#### 2️⃣ Create your remote server entry using `task remote-wizard`

The easiest way to create a remote server configuration is using the wizard:

```bash
task remote-wizard
```

The wizard will guide you through:
1. **Basic Information**: Server name and category
2. **Server Details**: Title, description, icon URL, and documentation URL
3. **Remote Configuration**: Transport type (streamable-http or sse) and server URL
4. **OAuth Configuration**: Simple yes/no question

If OAuth is enabled, the wizard automatically generates:
- **Provider**: Uses your server name (e.g., `linear`)
- **Secret**: `{server-name}.personal_access_token` (e.g., `linear.personal_access_token`)
- **Environment Variable**: `{SERVER_NAME}_PERSONAL_ACCESS_TOKEN` (e.g., `LINEAR_PERSONAL_ACCESS_TOKEN`)

This will create a directory under `servers/` with three files:
- `server.yaml` - Server configuration
- `tools.json` - Empty array (for dynamic tool discovery)
- `readme.md` - Documentation link

#### 3️⃣ Review the generated files

The wizard has created all necessary files for you. The `tools.json` file is always an empty array `[]` for remote servers because they use dynamic tool discovery. The `readme.md` file contains your documentation link.

#### 4️⃣ Example remote server structure

Your remote server directory should look like this:

```
servers/my-remote-server/
├── server.yaml      # Server configuration
├── tools.json       # Always [] for remote servers
└── readme.md        # Documentation link (required)
```

Example `server.yaml` for a remote server **with OAuth** (like `servers/linear`):

```yaml
name: linear
type: remote
dynamic:
  tools: true
meta:
  category: productivity
  tags:
    - productivity
    - project-management
    - remote
about:
  title: Linear
  description: Track issues and plan sprints
  icon: https://www.google.com/s2/favicons?domain=linear.app&sz=64
remote:
  transport_type: streamable-http
  url: https://mcp.linear.app/mcp
oauth:
  - provider: linear
    secret: linear.personal_access_token
    env: LINEAR_PERSONAL_ACCESS_TOKEN
```

Example `server.yaml` for a remote server **without OAuth** (like `servers/cloudflare-docs`):

```yaml
name: cloudflare-docs
type: remote
meta:
  category: documentation
  tags:
    - documentation
    - cloudflare
    - remote
about:
  title: Cloudflare Docs
  description: Access the latest documentation on Cloudflare products
  icon: https://www.cloudflare.com/favicon.ico
remote:
  transport_type: sse
  url: https://docs.mcp.cloudflare.com/sse
```

**Note:** Remote servers without OAuth don't need the `oauth` field or `dynamic.tools` field in their configuration.

#### 5️⃣ Test your remote server locally

You can test your remote server configuration by importing it into Docker Desktop:

```bash
task catalog -- my-remote-server
docker mcp catalog import $PWD/catalogs/my-remote-server/catalog.yaml
docker mcp server enable my-remote-server
```

For OAuth-enabled servers, authorize the server:

```bash
docker mcp oauth authorize my-remote-server
```

Now you can start the gateway with `docker mcp gateway run` and test tool calls to the remote server.

When done testing, reset the catalog:

```bash
docker mcp catalog reset
```

#### 6️⃣ Open a pull request

Create a pull request with your remote server files. Make sure to:
- Include all required files (`server.yaml`, `tools.json`, and `readme.md`)
- Verify that your server URL is publicly accessible
- Test OAuth configuration if applicable
- Ensure the documentation URL in `readme.md` is valid

## 📜 Code of Conduct

This project follows a Code of Conduct. Please review it in
[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## ❓ Questions

If you have questions, please create an issue in the repository.

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

# 🤝 Contributing to Docker MCP Registry
Thank you for your interest in contributing to the official Docker MCP Registry.
This document outlines how to contribute to this project.

## 🔄 Pull request process overview
- Fork the repository to your own GitHub account and clone it locally.
- Repository includes a `servers` folder where you should add a new folder with a `server.yaml` inside.
- Repository includes a `scripts` folder with bash scripts to automate some of the steps.
- Correctly format your commit messages, see Commit message guidelines below. _Note: All commits must include a Signed-off-by trailer at the end of each commit message to indicate that the contributor agrees to the Developer Certificate of Origin._
- Open a PR by ensuring the title and its description reflect the content of the PR.
- Ensure that CI passes, if it fails, fix the failures.
- Every pull request requires a review from the Docker team before merging.
- Once approved, all of your commits will be squashed into a single commit with your PR title.

## 📋 Step-by-Step Guide
### 1️⃣ Fork this repository
Fork the repository to your own GitHub account and clone it locally.

### 2️⃣ Add your entry locally
Add your entry by creating a new folder following the `owner@name` template, and create a `server.yaml` inside describing your MCP server. You will need to provide:
- A valid name for your MCP
- The GitHub URL of your project. The project needs to have a valid Dockerfile.
- A brief description of your MCP Server.
- A category for the MCP server, one of:
* 'ai'
* 'data-visualization'
* 'database'
* 'devops'
* 'ecommerce'
* 'finance'
* 'games'
* 'communication'
* 'monitoring'
* 'productivity'
* 'search'

#### 🚀 Generate folder and `server.yaml` using `new-server.sh` script
You can use our script to automate the creation of the files. Let's assume we have a new MCP Server to access my org's database. The MCP is called `My-ORGDB-MCP` and the GitHub repo is located at: `https://github.com/myorg/my-orgdb-mcp`

You can call the tool passing the MCP server name, category, and github url. 

```
./scripts/new-server.sh My-ORGDB-MCP databases https://github.com/myorg/my-orgdb-mcp
```

This will create a directory under `servers` as follows: `./servers/my-orgdb-mcp` and inside you will find a `server.yaml` file with your MCP definition.

```
server:
  name: test01
  image: mcp/test01
type: server
meta:
  category: test
  tags:
    - test
  highlighted: false
about:
  title: test01
  icon: https://avatars.githubusercontent.com/u/182288589?s=200&v=4
source:
  project: https://github.com/docker/mcp-registry
  branch: main
# config:
#   description: TODO
#   secrets:
#     - name: test01.secret_name
#       env: TEST01
#       example: TODO
#   env:
#     - name: ENV_VAR_NAME
#       example: TODO
#       value: '{{test01.env_var_name}}'
#   parameters:
#     type: object
#     properties:
#       param_name:
#         type: string
#     required:
#       - param_name
```

If you want to provide a specific Docker image built by your organisation, you can pass it to the script as follows:

```
IMAGE_NAME=myorg/myimage ./scripts/new-server.sh My-ORGDB-MCP databases https://github.com/myorg/my-orgdb-mcp
```

As you can see, the configuration block has been commented out. If you need to pass environmental variables or secrets, please uncomment the necessary lines.

🔒 If you don't provide a Docker image, we will build the image for you and host it in [Docker Hub's `mcp` namespace](https://hub.docker.com/u/mcp), the benefits are: image will include cryptographic signatures, provenance tracking, SBOMs, and automatic security updates. Otherwise, self-built images still benefit from container isolation but won't include the enhanced security features of Docker-built images.

### 3️⃣ Run & Test your MCP Server locally
🚧 tbd

### 4️⃣ Create `commit` and raise the Pull Request
🚧 tbd

### 5️⃣ Wait for review and approval
Upon approval your entry will be processed and it will be available in 24 hours at: 
- [MCP catalog](https://hub.docker.com/mcp) 
- [Docker Desktop's MCP Toolkit](https://www.docker.com/products/docker-desktop/) 
- [Docker Hub `mcp` namespace](https://hub.docker.com/u/mcp) (for MCP servers built by Docker)


## 📜 Code of Conduct

This project follows a Code of Conduct. Please review it in
[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## ❓ Questions

If you have questions, please create an issue in the repository.

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

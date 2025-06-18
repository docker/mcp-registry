# Contributing to Docker MCP Registry

Thank you for your interest in contributing to the official Docker MCP Registry.
This document outlines how to contribute to this project.

## Pull request process

- All commits must include a Signed-off-by trailer at the end of each commit message to indicate that the contributor agrees to the Developer Certificate of Origin.
- Fork the repository to your own GitHub account and clone it locally.
- Make your changes. To add a new MCP Server, create a new folder under `servers` with the name of your server and add a `server.yaml` inside.
- Correctly format your commit messages, see Commit message guidelines below.
- Open a PR by ensuring the title and its description reflect the content of the PR.
- Ensure that CI passes, if it fails, fix the failures.
- Every pull request requires a review from the Docker team before merging.
- Once approved, all of your commits will be squashed into a single commit with your PR title.

## Getting Started

You will need to provide:

- A valid name for your MCP
- The GitHub URL of your project. The project needs to have a valid Dockerfile.
- A brief description of your MCP Server.

Let's assume we have a new MCP Server to access my org's database. The MCP is called `My-ORGDB-MCP` and the GitHub repo is located at: https://github.com/myorg/my-orgdb-mcp We have created a bash script to simplify the creation process.

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

If you want to use a Docker image built by your organisation, you can pass it to the script as follows:

```
IMAGE_NAME=myorg/myimage ./scripts/new-server.sh My-ORGDB-MCP databases https://github.com/myorg/my-orgdb-mcp
```

As you can see, the configuration block has been commented out. If you need to pass environmental variables or secrets, please uncomment the
necessary lines.

## Testing your MCP Server

## Code of Conduct

This project follows a Code of Conduct. Please review it in
[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## Questions

If you have questions, please create an issue in the repository.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

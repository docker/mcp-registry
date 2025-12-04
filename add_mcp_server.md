Your task is to help a user add an MCP server to the registry.

First ask the user for the GitHub repo or the remote MCP server URL of the server they want to add.

If it's a GitHub repo, then it's a local MCP server and therefore needs a Dockerfile:
1. If it has a Dockerfile, then skip to step 4.
2. If it doesn't have a Dockerfile, create one based on the provided repo's readme.
3. Push the change with the Dockerfile to the main branch of the server's repo.
3.a. If the user doesn't have access to the main branch, try creating a new branch and committing there. Create a PR to merge the new branch into the main branch.
3.b. If the user can't create a new branch, fork the repo and make the change there. Create a PR to merge the changes in the forked repo into the main branch of the official repo. You'll need the "gh" command or a GitHub MCP server to fork the repo. If the user doesn't have either installed or setup, tell the user they need to have one of them setup before proceeding.
3.c. If you had to create a PR, tell the user they should get it merged into the main branch of the server repo and once it's merged, update their new entry so the server stays up to date.
4. Create the new entry in the mcp-registry repo. Use notion as a template but don't include the run or config sections if the server doesn't need them. If applicable, use the google favicon URL with the domain of the server's associated website as the server icon. Example: https://www.google.com/s2/favicons?domain=notion.so&sz=64
5. Run "task validate -- --name <server_name>" and "task build -- --tools <server_name>" replacing "<server_name>" with the name field of the server.yaml file and ensure they pass. If not, make changes until they pass or ask the user for help if you're stuck.
6. Create a PR with their new entry in the mcp-registry repo. Use .github/PULL_REQUEST_TEMPLATE.md as the PR template.

If it's a remote MCP server URL:
1. Create the new entry in the mcp-registry repo. Ensure it has a readme.md with a link to the remote server documentation and a tools.json which should contain an empty list "[]". Use notion-remote as a template but don't include the OAuth section if the remote server lacks OAuth.
2. Run "task validate -- --name <server_name>" and "task build -- --tools <server_name>" replacing "<server_name>" with the name field of the server.yaml file and ensure they pass. If not, make changes until they pass or ask the user for help if you're stuck.
3. Create a PR with their new entry to the mcp-registry repo. Use .github/PULL_REQUEST_TEMPLATE.md as the PR template.

After these steps, finish by telling the user exactly which PRs were created. If a new branch or fork of the server repo was created, tell them that once the PR into the main branch of the server repo is merged, that they should update the server entry in the mcp-registry repo with the main branch of the server repo.


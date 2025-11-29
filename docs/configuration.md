# Configuration

If an MCP Server needs configuration, the `server.yaml` should have an `config` attribute where you will be able to define the different parameters that the MCP Server needs as environment variables, secrets, or other parameters.

## Parameters

If your MCP Server needs to provide a way for user to add their own configuration, you will need to define a `parameters` attribute under `config`.

Parameters needs to define a type, properties and which of those properties are required:

```
config:
  parameters:
    type: object
    properties:
      my_param:
        type: string
    required:
      - my_param
```

To define a parameter, you need to add an attribute using the parameter name as attibute name. Properties need to define the type (string, array, integer...)

Lastly, if you have a parameter that it's required to run your MCP Server, you should add it to the `required` attribute. This will force users in Docker Desktop, to fill that parameter before adding the MCP Server to the catalog.

## Environment variables

To define an environment variable you need to define a paremeter first and then add an entry under the `env` attribute of `config`:

```
config:
  env:
    - name: ENVIRONMENT_VARIABLE_NAME
      example: ""
      value: '{{server_name.parameter_name}}'
  parameters:
    type: object
    properties:
      parameter_name:
        type: string
    required:
      - parameter_name
```

Remember that paremeters are used to create the appropriate input field in Docker Desktop. Also, it's important to notice that we use the `server name` as prefix for your parameters, so make sure that the `name` attribute matches the parameters prefix.

## Secrets

Secrets don't need to define a parameter since we handle them differently.

```
config:
  secrets:
    - name: server_name.api_key
      env: API_KEY
      example: YOUR_API_KEY
```

## Volumes

To define a `volume` you need to add a `run` attribute. If you need the user to define the host directory, the container directory or both, you will need to create first the appropriate parameter and then add the `run` block to the server.

```
run:
  volumes:
    - '{{server_name.volume_path}}:/path/inside/container'
config:
  description: example of a volume
  parameters:
    type: object
    properties:
      volume_path:
        type: string
    required:
      - volume_path
```

## Command

If you need to overwrite the command, you can do it in the `run` block:

```
run:
  command:
    - --transport=stdio
```

## Source Pinning

Local servers must pin their source repository to a specific Git commit using the `source.commit` field.

### Automatic Updates

Once an initial revision is accepted into the registry, an automated nightly GitHub Action will:

- Check for new commits on your configured branch (defaults to `main`)
- Create pull requests to update the `source.commit` field when newer commits are available
- These PRs are automatically reviewed and merged by the Docker team

After a pin update PR is merged, Docker's infrastructure automatically:

- Builds the new Docker image from the updated commit
- Signs the image with cryptographic signatures
- Generates provenance and SBOM
- Publishes to Docker Hub's `mcp/` namespace

**Timeline**: New images are typically available within 24 hours of the pin update PR being merged.

**Note**: Automatic updates only apply to Docker-built images (those with `image: mcp/...`). Servers using custom images must be updated manually.

See [CONTRIBUTING.md](../CONTRIBUTING.md#updating-your-mcp-server) for more details on the update process.

## User

If you need to run the container with a specific user, you can do it in the `run` block. If you want the user to be able to define the container user, you will need to create a parameter first and then add the `run` block to the server.

```
run:
  user: '{{server_name.container_user}}'
config:
  description: example of a user
  parameters:
    type: object
    properties:
      container_user:
        type: string
```

## Full Example

Here you can see a full example:

```
name: server_name
image: mcp/server_name
type: server
meta:
  category: devops
  tags:
    - server_name
about:
  title: Server Name
  description: Description about my Server. What it does and why it's useful
  icon: https://...
source:
  project: https://github.com/my-org/my-mcp-server
  commit: 0123456789abcdef0123456789abcdef01234567
run:
  command:
  - --transport=stdio
  volumes:
    - '{{server_name.path}}:/data'
  disableNetwork: true
config:
  description: The MCP server is allowed to access these path
  secrets:
  - name: server_name.api_key
    env: API_KEY
    example: YOUR_API_KEY
  env:
    - name: ENVIRONMENT_VARIABLE_NAME
      example: ""
      value: '{{server_name.parameter_name}}'
  parameters:
    type: object
    properties:
      path:
          type: string
        default:
          - /Users/local-test
      parameter_name:
        type: string
    required:
      - path

```

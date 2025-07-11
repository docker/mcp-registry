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

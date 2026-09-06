# RunComfy

Run [ComfyUI workflows](https://www.runcomfy.com/comfyui-workflows) in the cloud, generate images and videos, and train LoRAs from your MCP-compatible assistant.

- **ComfyUI workflows:** inspect and manage serverless deployments, submit workflow requests, track progress, and retrieve outputs.
- **Model inference:** discover [available models](https://www.runcomfy.com/models), inspect their input schemas, and run image or video generation with Seedance, Wan, FLUX, LTX, Seedream and other supported models.
- **LoRA training:** prepare datasets, submit [AI Toolkit training](https://www.runcomfy.com/trainer/ai-toolkit) on GPU, monitor steps, and retrieve checkpoints and sample outputs.

The hosted server currently exposes 31 tools through dynamic discovery. Model availability and supported parameters are returned by the model catalog and schema tools.

## Connect

Enable RunComfy in Docker MCP Toolkit and complete browser authorization. On the RunComfy authorization page, enter an API token from your [RunComfy profile](https://www.runcomfy.com/profile) and approve access. Docker uses the resulting OAuth grant; do not put your API token into the catalog files.

Endpoint: `https://mcp.runcomfy.com/mcp`

Transport: Streamable HTTP

A RunComfy account is required. Inference and training use the account's paid balance. Review the selected model or workflow, inputs, GPU and training settings before submitting a job.

Start with: "List my RunComfy deployments" or "Find Seedream models and show the input schema before running one."

## Documentation

- [MCP introduction and tools](https://docs.runcomfy.com/mcp/introduction)
- [Connection quickstart](https://docs.runcomfy.com/mcp/quickstart)
- [Model API](https://docs.runcomfy.com/model-apis/quickstart)
- [ComfyUI Serverless API](https://docs.runcomfy.com/serverless/introduction)
- [LoRA Trainer API](https://docs.runcomfy.com/trainer-apis/introduction)
- [Public MCP repository](https://github.com/runcomfy-com/runcomfy-mcp)
- [RunComfy website](https://www.runcomfy.com)

For support or security reports, contact [hi@runcomfy.com](mailto:hi@runcomfy.com).

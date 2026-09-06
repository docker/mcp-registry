# RunComfy

Use RunComfy as an **AI image generator** and **AI video generator**, run [ComfyUI workflows](https://www.runcomfy.com/comfyui-workflows) in the cloud, and train LoRAs from your MCP-compatible assistant.

- **ComfyUI workflows:** inspect and manage serverless deployments, submit workflow requests, track progress, and retrieve outputs.
- **AI image models:** discover [image generation and editing options](https://www.runcomfy.com/models), including FLUX 2, FLUX Kontext, Seedream 5.0 Pro and Seedream 4.5. Inspect the selected model's input schema before generating or editing an image.
- **AI video models:** use Seedance 2.5, Wan 3 (Wan 3.0 and Wan 3.0 Prime), Wan 2 series (2.1, 2.2, 2.5, 2.6 and 2.7), and LTX 2.5 for supported text-to-video, image-to-video and other video tasks.
- **LoRA training:** prepare datasets, submit [AI Toolkit training](https://www.runcomfy.com/trainer/ai-toolkit) on GPU, monitor steps, and retrieve checkpoints and sample outputs.

The hosted server currently exposes 31 tools through dynamic discovery. These model versions were verified in the authenticated Model API on September 6, 2026. Use the model discovery and schema tools for current availability and supported inputs; tasks vary by model. LoRA training uses a separate AI Toolkit configuration and supported training base model.

## Connect

Enable RunComfy in Docker MCP Toolkit and complete browser authorization. On the RunComfy authorization page, enter an API token from your [RunComfy profile](https://www.runcomfy.com/profile) and approve access. Docker uses the resulting OAuth grant; do not put your API token into the catalog files.

Endpoint: `https://mcp.runcomfy.com/mcp`

Transport: Streamable HTTP

A RunComfy account is required. Inference and training use the account's paid balance. Review the selected model or workflow, inputs, GPU and training settings before submitting a job.

Start with: "List my RunComfy deployments", "Find a Seedream 5.0 AI image model and show its input schema", or "Compare Wan 3, Wan 2.2 and Seedance 2.5 AI video models before generating a clip."

## Documentation

- [AI image models and AI video models](https://www.runcomfy.com/models)
- [MCP introduction and tools](https://docs.runcomfy.com/mcp/introduction)
- [Connection quickstart](https://docs.runcomfy.com/mcp/quickstart)
- [Model API](https://docs.runcomfy.com/model-apis/quickstart)
- [ComfyUI Serverless API](https://docs.runcomfy.com/serverless/introduction)
- [LoRA Trainer API](https://docs.runcomfy.com/trainer-apis/introduction)
- [Public MCP repository](https://github.com/runcomfy-com/runcomfy-mcp)
- [RunComfy website](https://www.runcomfy.com)

For support or security reports, contact [hi@runcomfy.com](mailto:hi@runcomfy.com).

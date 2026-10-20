# InkRoom Prompts

Search public, approved image prompts with their original example images, or rewrite an image idea within shared free service limits. Optimized prompts do not include library example images. This tool does not generate images.

- [Official prompt tools](https://inkroom.ai/tools/prompts)
- [Usage and service limits](https://inkroom.ai/tools/prompts/docs)
- [Audited source and Docker installation](https://github.com/InkRoom-AI/inkroom-client/tree/8888ff373d01945a987ea63199b67562f7c8a5aa/prompts)
- [Free implementation MIT license](https://github.com/InkRoom-AI/inkroom-client/blob/8888ff373d01945a987ea63199b67562f7c8a5aa/prompts/LICENSE)
- [Report a client or security issue](https://github.com/InkRoom-AI/inkroom-client/issues)

No account, API key, MCP Token, environment setup, ports or volumes are required. The server runs locally over STDIO as an unprivileged user and uses outbound HTTPS to the public InkRoom prompt service. Free usage is rate-limited. `search_prompts` retrieves library records; `optimize_prompt` performs text optimization. The image is not the separate paid InkRoom image MCP.

Version 1.0.3 was tested in a Linux/amd64 container with real MCP initialization/tool discovery, invalid-input/offline errors and one anonymous public search. Other CPU architectures and native Docker Desktop Toolkit installation were not tested. MIT applies only to `prompts/` and `skills/inkroom-prompts/` in the source repository; it does not license the separate paid image client or third-party prompts/images.

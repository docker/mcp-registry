## EasyBits

The cloud for AI agents. A persistent Linux microVM the agent can keep working in, plus the
storage, data and publishing it needs around it — all via MCP.

### Features

- **Sandboxes**: Firecracker microVMs with root and real internet. Snapshot, suspend and
  resume in under a second; the box survives the call instead of being destroyed.
- **Hosting**: turn a box into a deployed app with a public URL and its own domain, with
  releases, rollback and daily backups.
- **Web**: search, fetch and structured extraction from pages that block ordinary scrapers.
- **Files**: upload, version and serve over a CDN; share links, image optimization, bulk ops.
- **Databases**: create libSQL databases, query, import data.
- **Documents**: generate PDFs, multi-page documents and presentations from HTML.
- **Websites & forms**: deploy static sites; collect submissions.
- **Voice & video**: transcription, speech and rendered video projects.

### Setup

1. Get your API key at https://www.easybits.cloud/dash/developer
2. Set the `EASYBITS_API_KEY` environment variable

The endpoint also supports OAuth 2.1 with dynamic client registration.

### Also available via npm

```bash
npx @easybits.cloud/mcp
```

Docs: https://www.easybits.cloud/docs

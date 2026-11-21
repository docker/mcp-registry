# Engram Memory

Persistent memory for AI agents. Engram gives any MCP client a brain that survives across sessions: it stores facts, preferences, decisions, and corrections, then retrieves them semantically when relevant.

## What it does

Engram exposes six tools to MCP clients:

- `memory_store` — store a memory with semantic embedding
- `memory_search` — three-tier recall (hot cache → hash index → vector search)
- `memory_recall` — context injection for system prompts
- `memory_forget` — remove a memory from all tiers
- `memory_consolidate` — merge near-duplicates
- `memory_connect` — discover cross-category connections

## Architecture

A single container bundles three components, supervised by s6-overlay:

1. **Qdrant** — vector database (port 6333)
2. **FastEmbed** — ONNX-based embedding service, native ARM64 + x86_64 (port 11435)
3. **MCP HTTP server** — three-tier recall engine exposing memory tools (port 8585)

The recall engine implements:

- **Tier 1**: Hot-tier cache with ACT-R cognitive activation (sub-millisecond)
- **Tier 2**: Multi-head LSH hash index for O(1) candidate retrieval
- **Tier 3**: Hybrid Qdrant search (dense vectors + BM25 sparse with RRF fusion)
- **Graph layer**: Kuzu-backed entity tracking with spreading activation
- **Consolidation**: Janitor (deduplication) and Librarian (cross-category linking)

Memories persist in `/data` (mount a volume).

## Performance

On Apple Silicon (M-series, ARM64) with 100 memories:

| Metric               | Cold      | Warm   |
| -------------------- | --------- | ------ |
| Search latency       | ~190 ms   | ~24 ms |
| Store latency        | ~42 ms    | —      |
| Embedding throughput | 60+ emb/s | —      |
| Top-3 accuracy       | 97%       | —      |

The cold path runs the full three-tier pipeline (hash candidates → hybrid search → graph expansion). The warm path returns from the hot tier in roughly the time it takes to embed the query.

## Source

[github.com/EngramMemory/engram-memory-community](https://github.com/EngramMemory/engram-memory-community)

## License

MIT

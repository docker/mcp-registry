# Knowledge Base Hybrid RAG MCP Server

A powerful MCP server providing **37 tools** for Hybrid RAG (Retrieval-Augmented Generation) with PostgreSQL/pgvector.

## Features

- **Hybrid RAG Search**: Semantic + Keyword search with configurable weights
- **LLM Reranking**: +5-10% accuracy improvement
- **In-Memory Caching**: 2-3x speed improvement
- **Knowledge Graph**: Entity/relation-based knowledge management
- **Auto Entity Extraction**: LLM-based NER with automatic KG save
- **Conversational Memory**: Session-based memory for conversational RAG
- **Multiple Chunking Strategies**: fixed, paragraph, semantic
- **Evaluation Metrics**: Precision, Recall, F1, MRR

## Documentation

Full documentation: https://github.com/hwandam77/Knowledge-Base-Hybrid-RAG

## Requirements

- PostgreSQL with pgvector extension
- Ollama with embedding model (qwen3-embedding:0.6b)
- Ollama with LLM model (gemma3:12b or similar)

## License

MIT License

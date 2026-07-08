# LeiGo — Portuguese Legal Assistant

Semantic search and AI-powered Q&A across consolidated Portuguese legislation (Diário da República).

**Requires a paid LeiGo subscription.** Generate an API key from your [LeiGo dashboard](https://leigo.pt) and set it as `LEIGO_API_KEY`.

## Tools

| Tool | Description |
|------|-------------|
| `search_laws` | Semantic search across Portuguese legislation (up to 5 queries, deduped). |
| `ask_legal_question` | Full RAG answer with citations. |
| `get_articles` | Fetch article text by act reference + article number. |
| `get_sources` | List indexed legislation corpora with counts and last-sync timestamp. |
| `get_usage` | Check subscription tier and budget usage. |

## Auth

Set `LEIGO_API_KEY` to a paid-tier API key from your [LeiGo dashboard](https://leigo.pt).

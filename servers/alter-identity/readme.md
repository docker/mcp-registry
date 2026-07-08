## ALTER Identity MCP Server

Identity infrastructure for the AI economy — the science of being known.

ALTER delivers verified human identity to any MCP client via a 33-trait psychometric engine. It provides belonging probability, full trait vectors, attunement depth scores, and privacy-gated inference — built on psychometric science (BRAVING Trust Inventory, Bayesian adaptive assessment).

### What it does

- **Psychometric assessment** — 33 traits across 5 categories, Bayesian adaptive delivery
- **Belonging probability** — composite fit score: `0.40 × authenticity + 0.35 × acceptance + 0.25 × complementarity`
- **Trait vectors** — full 33-dimension identity field, privacy-gated by consent tier
- **Attunement depth** — measures how well an agent knows a person over time (replaces streaks)
- **Identity verification** — verify ~handle ownership via DNS TXT + MCP discovery

### Tools (free tier — 16 tools)

`get_profile`, `get_trait_snapshot`, `get_full_trait_vector`, `compute_belonging`, `assess_traits`, `verify_identity`, `get_identity_trust_score`, `get_engagement_level`, `get_match_recommendations`, `list_archetypes`, `get_competencies`, `search_identities`, `query_matches`, `get_side_quest_graph`, `get_privacy_budget`, `recommend_tool`

### Access

- **Free tier:** 16 tools · 10 req/min · 100 req/day — no API key required
- **Premium:** x402 micropayment-gated tools and higher rate limits
- **API key:** set `ALTER_API_KEY` for authenticated premium access

### Docs

- Source / SDK: https://github.com/true-alter/alter-identity
- MCP namespace: `com.truealter/identity` v0.3.1
- DNS discovery: `_alter.truealter.com` TXT

Folklore Clinical Variant Interpretation MCP 1.4.2 provides four scientific
tools for public GRCh38 germline variant evidence, ACMG/AMP decision support,
variant literature and semantic Literature Corpus search. The complete hosted
catalog currently also includes the separate infrastructure-support helper.

Canonical endpoint: https://api.helena.bio/folklore/v1/mcp
Docs: https://folklore.helena.bio/integrations

No account or API key is required. Accepts no patient, phenotype, family,
segregation or private case context. Results support qualified professional
review and are not a diagnosis or treatment recommendation.

The endpoint uses stateless MCP 2026-07-28 discovery via `server/discover`.
A legacy inspector that only performs `initialize` is not compatible.
Dynamic tool discovery is retained; an empty tools.json is not a successful
Docker inspector or Toolkit compatibility result.

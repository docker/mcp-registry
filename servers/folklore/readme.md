Folklore Clinical Variant Interpretation MCP 1.4.2 provides four scientific
tools for public GRCh38 germline variant evidence, ACMG/AMP decision support,
variant literature and semantic Literature Corpus search. The complete hosted
catalog currently also includes the separate infrastructure-support helper.

Canonical endpoint: https://api.helena.bio/folklore/v1/mcp
Docs: https://folklore.helena.bio/integrations

No account or API key is required. Accepts no patient, phenotype, family,
segregation or private case context. Results support qualified professional
review and are not a diagnosis or treatment recommendation.

A public initialize probe offering MCP 2025-03-26 returned HTTP 200 and
serverInfo version 1.4.2 on 2026-09-11. This supersedes the older assertion
that every initialize-based inspector is incompatible. Docker must still run
its own complete inspector and Toolkit checks. Dynamic tool discovery is
retained; an empty tools.json is not proof of successful verification.

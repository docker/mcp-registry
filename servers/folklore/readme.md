Folklore Clinical Variant Interpretation MCP 1.5.0 provides six scientific
tools for genomic variant interpretation, ClinGen gene-disease evidence and
source-linked literature, plus the separate infrastructure-support helper.
`get_gene_disease_associations` accepts an exact gene symbol/HGNC identifier;
`search_disease_genes` accepts an exact MONDO identifier or disease-name
substring. Distinct diseases, source assertions, inheritance and pagination
remain visible. An empty ClinGen result does not establish no association.
The existing GRCh38 germline variant and literature tools remain available.

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

Release 1.5.0 and its seven-tool hosted catalog are verified on 2026-09-11.
The new tools do not upload VCFs, parse raw sequencing, run whole-genome
pipelines or make patient diagnoses. Initial gene-disease coverage is ClinGen
Gene-Disease Validity, not all known associations. Docker inspector/Toolkit
acceptance remains maintainer-controlled.

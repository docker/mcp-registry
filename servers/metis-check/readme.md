# Metis Check

Checks the citations in an answer at the registries that issue them: DOI, PubMed, arXiv, ORCID, ROR, ISBN, LEI, ISIN, CUSIP, SEC CIK, EIN, NPI, UK company numbers and more. Each identifier comes back in exactly one of four states (settled true, settled false, not decidable, unreachable) and every check files a signed report with its own id, a PDF and a recheck button.

Hosted Streamable HTTP server with OAuth sign-in; a free Metis account is created on first connection (5 checks a month). Docs: https://mcp.metisfoundry.com/docs. Plans: https://check-api.metisfoundry.com/check/plans.

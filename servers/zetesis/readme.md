# Zetesis

Scientific due diligence on a claim. Five public registers, all official APIs: Europe PMC,
ClinicalTrials.gov, openFDA, NIH RePORTER and SEC EDGAR.

Give an agent a claim, an abstract, a paper or a pitch deck and it returns the questions a
domain reviewer would ask, the failure patterns that caught comparable claims before, and the
evidence bearing on it, with a PMID, DOI, NCT number, NIH grant number or SEC filing reference
on every source. Every identifier it returns was retrieved, not generated.

Retrieval can be fenced to a year, so a claim is judged on what was knowable at the time rather
than on how it turned out.

The two evidence tools call no language model. They return in under a second and send nothing to
any model provider.

No account, key or token is required.

Documentation: https://api.zetesis.science/docs
Privacy policy: https://api.zetesis.science/privacy

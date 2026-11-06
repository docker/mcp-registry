Folklore MCP 2.0.0 provides 11 public tools for ACMG/AMP germline variant
classification, gene profiles, ClinGen gene-disease associations and scientific
literature. The configured connection works without an account or API key.
A client configured separately with a Folklore Bearer API key can access 18
additional tools for case analysis, approved source files and report retrieval
at the same hosted endpoint.

Canonical endpoint: https://api.helena.bio/folklore/v1/mcp
Setup and scope: https://folklore.helena.bio/integrations/mcp

Public variant classification accepts one supported GRCh38 nuclear germline
SNV or simple indel smaller than 50 bp. Keep private case data out of the
public connection. Results require qualified professional review.

Docker discovers the public tools dynamically. Docker Toolkit installation
and authenticated case workflows remain to be verified in Docker.

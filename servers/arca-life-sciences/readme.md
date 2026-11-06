# Arca Life Sciences MCP Server

A free remote MCP server for biomedical, clinical and regulatory research, run by [Arca](https://www.arca.inc). No API key or sign-up.

## Tools

- **search_papers** – ranked search over ~41M PubMed records, with OpenAlex citation counts and filters for open access, publication year and article type.
- **search_medical** – FDA drug labels, Drugs@FDA approvals, complete response letters, warning letters, Form 483s, inspections, Orange Book, Purple Book, NDC, drug shortages, 510(k)/PMA devices, MAUDE adverse events, GUDID, ClinicalTrials.gov studies, VAERS reports, CDC data, EU MDR/IVDR expert panel opinions and BfArM/PMDA device vigilance notices.
- **fetch_source_record** – the full text of one record.
- **list_datasets** – what the corpus covers.

Every result cites its source URL. The server is read-only, serves public data only, and is rate limited to 120 requests per minute per IP.

## More Information

Docs: [arca.inc/mcp](https://www.arca.inc/mcp). Official MCP Registry name: `inc.arca/life-sciences`.

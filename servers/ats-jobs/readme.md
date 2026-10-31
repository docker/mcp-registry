# Company Facts (ATS Jobs) MCP Server

Company data read from the official or original source, each record with a link to where it came from.

## Company tools

- `new_businesses`: newly formed businesses from the official NY, CO, OR, CT and PA Secretary of State registers, in one record format. Feed by formation date, search by name, or look up by entity number. A website is attached only when the site's own pages show the exact legal name and the register's city, state or ZIP.
- `company_search`: filter the companies we track by state, city, technology on their site, published contacts and hiring.
- `company_lookup`: one company from its domain: registrar, DNS and email provider, tech stack, hiring, register entry or SEC figures.
- `search_company_registry`, `validate_vat`, `sec_financials`, `domain_lookup`, `detect_tech_stack`, `federal_contracts`.

## Job tools

- `find_company_board`, `list_open_jobs`, `compare_companies`, `job_changes`, `hiring_summary`: open jobs read first-hand from Greenhouse, Ashby, Lever and Workable boards, from a directory of 8,532 boards, with salaries parsed into numbers only where unambiguous.

## Notes

- No key needed for the free tier (10 calls a day). Past that, pay per call with x402 or use a prepaid key.
- All tools are read-only. Register records never carry officers, agents, emails, phones or street addresses; recruiter names are never returned.

## More Information

See the [project repository](https://github.com/groundtruthtools/ats-jobs-mcp) and [llms.txt](https://companyfacts.dev/llms.txt).

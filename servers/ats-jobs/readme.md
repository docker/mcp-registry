# ATS Jobs MCP Server

Open job postings, read first-hand from the applicant-tracking system behind a company's careers page: Greenhouse, Ashby and Lever.

Those vendors publish open job board APIs, but each is useless unless you already know the company's board name. This server carries a directory of 7,479 company boards, each confirmed against the vendor's own API.

## Tools

- `find_company_board`: which system a company's job board runs on
- `list_open_jobs`: every open job at one company, with salaries parsed into numbers where unambiguous
- `compare_companies`: open roles across several companies side by side

## Notes

- No key needed for the free tier. A paid key raises the per-call cap.
- No personal data: recruiter names are never returned.
- Every posting carries the employer's own URL, so any answer can be checked at the source.

## More Information

See the [project repository](https://github.com/groundtruthtools/ats-jobs-mcp) and [llms.txt](https://ats-jobs-api.groundtruthtools.workers.dev/llms.txt).

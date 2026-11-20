# SrcLog

SrcLog lets AI assistants search and inspect open-source GitHub repositories, curated topics and developer rankings from [SrcLog.com](https://srclog.com). Data covers ~450,000 repositories in 435 curated topics and is refreshed daily from the public GitHub API.

## Tools

- **search_repos** – hybrid keyword + semantic repository search with language, topic and star filters
- **get_repo** – stars, forks, language, license, topics, SrcLog score, dates and links for one repository
- **similar_repos** – semantically similar repositories, useful for finding alternatives
- **list_topics** / **topic_repos** – curated topics and their most-starred repositories
- **top_developers** – rankings of developers or organizations by stars, forks, watchers or repos, by language and country
- **get_owner** – profile, totals, ranking positions and top repositories of a GitHub user or organization

All tools are read-only. No authentication is required; tool calls are limited to 30 per minute per IP.

## More Information

Docs: https://srclog.com/for-developers
Source: https://github.com/srclogcom/srclog-mcp

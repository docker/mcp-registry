# KissMyPost: Reddit Audience Research

**KissMyPost** gives your AI assistant Reddit analytics and audience research. Ask about any subreddit, topic or brand and get answers from KissMyPost data instead of guesses: which subreddits to find for a topic, who is in them, how big they are and whether they are growing, the best time to post, what the posting rules allow, and how often Reddit mentions a brand or a phrase. It works as a subreddit finder and Reddit marketing tool for founders, marketers and researchers.

Remote MCP server at `https://kissmypost.com/api/mcp` (Streamable HTTP). Auth: `Authorization: Bearer <your KissMyPost API key>`. 17 read-only tools. Not affiliated with Reddit, Inc.

## Tools

- **Find subreddits** – `search_subreddits`, `find_subreddits_by_topic`, `top_subreddits`, `similar_subreddits`
- **Subreddit stats** – `subreddit_overview` (members, 12-month growth, posts and authors per month, top topics)
- **Audience research** – `subreddit_audience` (share of women, median age, where else the regulars hang out), `compare_subreddits`
- **When and where to post** – `best_time_to_post`, `where_to_post`, `subreddit_posting_rules`
- **Brand and mention search** – `brand_mentions`, `search_reddit_mentions`, `search_reddit_posts`
- **Topics and research** – `explore_topics`, `search_research`, `reddit_statistics`
- **Threads worth a reply** – `reply_radar`

Most tools work on a free account; some need a paid plan. Limits and plans: [kissmypost.com/pricing](https://kissmypost.com/pricing?utm_source=docker&utm_medium=directory).

## Example prompts

- Which subreddits talk about home espresso?
- How big is r/SaaS and is it growing?
- When should I post in r/startups?
- Who are the people in r/espresso?
- Where on Reddit do people talk about notion.so?

## Get a free API key

Sign up and copy your key at [kissmypost.com/mcp](https://kissmypost.com/mcp?utm_source=docker&utm_medium=directory).

Docs: https://github.com/krutytskyi/kissmypost-mcp

# PrismCrawl

[PrismCrawl](https://www.prismcrawl.com/) provides live search data for AI agents. Its hosted MCP server gives assistants access to search engine results, maps and places, shopping, reviews, and app store data, returned as structured JSON.

## Features

- **Web search** – Google, Bing, DuckDuckGo, and Amazon search results.
- **Maps and places** – Place results from Google Maps, Bing Maps, Apple Maps, DuckDuckGo Maps, Yelp, and Tripadvisor.
- **Reviews** – Reviews from Google Maps, Apple Maps, Yelp, and Tripadvisor, plus Google contributor reviews.
- **App stores** – Google Play apps, games, books, and movies, and Apple App Store listings, product details, and reviews.
- **Shopping** – Google Shopping product details and Amazon search results.

## Connecting

The server is hosted at `https://api.prismcrawl.com/mcp` and uses the streamable HTTP transport. No local installation is required.

The server requires a PrismCrawl API key:

1. Create an account at [prismcrawl.com](https://www.prismcrawl.com/). New accounts include free credits.
2. Copy your API key from the PrismCrawl dashboard.
3. In Docker Desktop's MCP Toolkit, enable the PrismCrawl server and enter your API key when prompted.

The key is sent to the server in the `x-api-key` header. Successful tool calls use your PrismCrawl credits. See the documentation for details.

## More Information

- Website: https://www.prismcrawl.com/
- MCP server: https://www.prismcrawl.com/mcp
- Docs: https://www.prismcrawl.com/docs

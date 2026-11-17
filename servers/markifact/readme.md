# Markifact Ads and Analytics

[Markifact MCP](https://www.markifact.com/mcp) connects AI assistants to Google Ads, Meta Ads, GA4 and 50+ more marketing platforms through one remote MCP server, with 1,100+ operations for reporting, audits and campaign management. Sign in with OAuth, connect your accounts once, and ask in plain language.

## What you can do

- Pull performance reports across ad accounts (spend, conversions, ROAS, CPA) and compare periods
- Audit campaigns, keywords, search terms, audiences and tracking
- Create and update campaigns, ad groups, ads, budgets and bids
- Combine ad data with analytics, SEO, ecommerce and CRM data in the same conversation

## Supported platforms

**Advertising**
[Google Ads](https://www.markifact.com/google-ads-mcp),
[Meta Ads](https://www.markifact.com/meta-ads-mcp) (Facebook and Instagram ads),
[Microsoft Ads](https://www.markifact.com/microsoft-ads-mcp) (Bing Ads),
[LinkedIn Ads](https://www.markifact.com/linkedin-ads-mcp),
[TikTok Ads](https://www.markifact.com/tiktok-ads-mcp),
[Amazon Ads](https://www.markifact.com/amazon-ads-mcp),
[Pinterest Ads](https://www.markifact.com/pinterest-ads-mcp),
[Snapchat Ads](https://www.markifact.com/snapchat-ads-mcp),
[Reddit Ads](https://www.markifact.com/reddit-ads-mcp),
[X Ads](https://www.markifact.com/x-ads-mcp),
[Spotify Ads](https://www.markifact.com/spotify-ads-mcp),
[Apple Ads](https://www.markifact.com/apple-ads-mcp),
[OpenAI Ads](https://www.markifact.com/openai-ads-mcp) (ChatGPT ads),
Display & Video 360

**Analytics and SEO**
[Google Analytics 4](https://www.markifact.com/google-analytics-mcp),
[Google Search Console](https://www.markifact.com/google-search-console-mcp),
[Google Tag Manager](https://www.markifact.com/google-tag-manager-mcp),
[Google Merchant Center](https://www.markifact.com/google-merchant-center-mcp),
[Google Business Profile](https://www.markifact.com/google-business-profile-mcp),
[YouTube Analytics](https://www.markifact.com/youtube-mcp),
Bing Webmaster Tools,
[Ahrefs](https://www.markifact.com/ahrefs-mcp),
Semrush,
DataForSEO

**Social media**
Facebook Pages,
[Instagram](https://www.markifact.com/instagram-mcp),
[LinkedIn](https://www.markifact.com/linkedin-mcp),
[TikTok](https://www.markifact.com/tiktok-mcp),
Pinterest

**Ecommerce, CRM and email**
[Shopify](https://www.markifact.com/shopify-mcp),
HubSpot,
[Klaviyo](https://www.markifact.com/klaviyo-mcp)

**Data and reporting**
[Google Sheets](https://www.markifact.com/google-sheets-mcp),
[BigQuery](https://www.markifact.com/bigquery-mcp),
Google Docs, Google Slides, Google Drive, Slack

## Google Ads

The [Google Ads MCP](https://www.markifact.com/google-ads-mcp) covers reports, GAQL queries, keyword planning, search term reviews, and campaign, ad group, keyword and budget changes.

## Meta Ads

The [Meta Ads MCP](https://www.markifact.com/meta-ads-mcp) covers Facebook and Instagram ad reports with breakdowns, ad set and budget updates, ad creation, audiences and lead forms.

## How changes are handled

Read operations run through `run_operation`. Anything that creates, updates or deletes data runs through a separate `run_write_operation` tool, so your MCP client can require approval before anything changes in your accounts.

## Tools

- `find_operations`: search 1,100+ marketing operations
- `get_operation_inputs`: get an operation's input schema
- `run_operation`: run a read-only operation
- `run_write_operation`: run an operation that changes data
- `list_connections`: list connected accounts
- `upload_media`, `prepare_media_upload`: upload images and videos for ads
- `get_file_url`, `read_file`: read files returned by operations
- `report_markifact_feedback`: send product feedback

Docs: https://docs.markifact.com

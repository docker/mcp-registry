# OpenSubtitles MCP Server

The official MCP server for [OpenSubtitles.com](https://www.opensubtitles.com). AI agents can find, inspect and download subtitles for movies and TV episodes in 100+ languages, with structured results.

- Search by title, filename/release name (parsed with guessit-js), IMDb/TMDB id, season/episode or moviehash; exact title and release-name matches rank first.
- Subtitle details, uploader comments and related subtitles.
- Download subtitle text (with optional format, FPS and time-shift conversion) on your own account's daily quota.

Credentials: your OpenSubtitles username and password (a free account is enough; VIP accounts get a higher quota). They are sent as connection headers, never as tool arguments, and the server never logs them.

Docs: https://github.com/opensubtitles/mcp.opensubtitles.com

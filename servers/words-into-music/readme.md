# Words into Music

Documentation: https://wordsintomusic.com/mcp/

Setup guide: https://wordsintomusic.com/blog/mcp-setup-guide/

Connect to `https://mcp.wordsintomusic.com/mcp` using Streamable HTTP and OAuth 2.0. Public account signup is available. Generation uses your existing account credits and limits. Preparing a song, checking progress, checking usage, and reading your library do not use song credits. Starting a recording uses one song credit.

Tools are discovered after authorization: `get_account_usage`, `prepare_song`, `start_song`, `get_song_job`, `list_songs`, and `get_song`.

The OAuth secret field in the catalog is Docker's token storage convention. Words into Music does not provide a manually created personal access token. Use OAuth authorization in a compatible client.

Support and security contact: contact@wordsintomusic.com.

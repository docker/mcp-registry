# Paper Plane Post

Create and send an AI-illustrated postcard anywhere in the world

[Connection instructions and examples](https://paperplanepost.com/for/agents)

Connect to `https://api.paperplanepost.com/v1/mcp` using Streamable HTTP without authentication.

`start_postcard` creates an unpaid draft from a photo, with an optional note and recipient address. Open the returned link to review the card, complete any missing details, and pay for printing and mailing. `show_postcard` displays a previously returned draft link. Tool calls never charge money or mail a postcard.

[Public client packages and registry metadata](https://github.com/paperplanepost/agents) are MIT-licensed; the hosted server implementation is proprietary. No Paper Plane Post account, API key, or reviewer credentials are required. Reviewers can use `https://paperplanepost.com/icon.png` as a sample image and stop before checkout.

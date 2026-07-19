# Latchshot

Latchshot captures public web pages as PNG, JPEG, or PDF through hosted Chromium. Artifacts are returned inline to MCP clients, so no local browser or package installation is required.

- Website and Free-plan key: https://latchshot.fly.dev/#trial
- MCP documentation: https://latchshot.fly.dev/docs.md#hosted-mcp-server
- Public remote-server contract: https://github.com/BaiqingL/latchshot-mcp
- Official MCP Registry name: `io.github.BaiqingL/latchshot`

## Tools

- `capture_page`: capture a public HTTP(S) URL as an inline image or embedded PDF resource, with render diagnostics and remaining quota.
- `get_usage`: read the current plan, successful-render usage, remaining quota, and UTC reset time.

## Authentication

Create a key through the email-only Free plan, then configure it as `LATCHSHOT_API_KEY`. The Free allowance renews to 100 successful renders each UTC calendar month. Only successful captures consume quota.

## Safety boundaries

Private, loopback, link-local, special-use, mixed public/private DNS, and metadata-network targets are blocked. Latchshot does not accept cookies, authenticated sessions, arbitrary browser scripts, CAPTCHA solving, proxy rotation, or private-network rendering.

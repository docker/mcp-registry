Docs: https://qr-branding.com/docs/integrations/mcp (9 languages)

Remote server: nothing to install. It works without an API key in a keyless
trial (PNG up to 720 px, 20 tool calls/day per network, no AI designs); a
Studio key unlocks SVG/PDF, up to 4K and AI-prompted designs.

OAuth 2.1 is supported too (dynamic client registration + PKCE), for clients
that prefer a browser consent flow over pasting a key:

- metadata: https://qr-branding.com/.well-known/oauth-protected-resource
- authorization server: https://qr-branding.com/.well-known/oauth-authorization-server

Tools: qr_generate, qr_generate_ai, qr_list_templates,
qr_generate_from_template, qr_build_wifi, qr_build_vcard, qr_build_geo.

Source: https://github.com/quartzondev/qr-branding-mcp

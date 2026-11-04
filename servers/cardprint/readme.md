Docs: https://www.poketdev.com

# Cardprint

Deterministic legality and rotation-risk screener for the Pokemon TCG Standard
format, packaged as an MCP tool so coding agents can call it like any other
tool in their stack. The screen is seeded: same deck, same seed, same answer
every time, which is what makes it safe to wire into automated pipelines.

- `screen_deck` takes a list of set marks and returns per-card legality, a
  0-100 rotation-risk score per card, deck-wide legality percentage, and the
  average risk. Marks outside the current regulation list screen as illegal.
- Runs as a plain stdio MCP server inside the container; the same engine also
  ships behind an HTTP transport in the upstream repo.

Upstream source: https://github.com/scrapertweeter3-prog/cardprint

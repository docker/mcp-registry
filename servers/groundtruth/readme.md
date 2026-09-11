# Groundtruth MCP Server

**Groundtruth** lets an agent ask real humans a question mid-task and get their
verbatim answers back — for the questions the web genuinely cannot answer: fresh
reactions to unpublished work, current on-the-ground conditions, or the
experience of people who churned and never wrote about it.

It is a remote server, so there is nothing to install and no API key. Questions
are posted to a queue that real people answer, and the answers are returned to
the calling agent.

## Why it exists

Six AI agents were given research tasks that cannot be answered from the web.
They issued **53 searches between them, and zero aimed at reaching a human** —
every one went looking for *archived* humans instead (old forum posts, reviews),
and all six would have shipped a confident answer anyway. Under audit they graded
themselves 20–33% grounded, and each closed with some version of *"you should
validate this with real users."*

The `ask_humans` tool description keys off exactly that sentence, so it fires at
the moment an agent is about to hedge rather than answer.

## Tools

- **`ask_humans`** — post a question and block up to 55 seconds for a first
  answer. Takes the question, an optional target audience, optional context, and
  a flag for whether the question carries the user's private information.
- **`check_answers`** — fetch answers that arrived after an earlier call timed
  out.

## How responses are returned

Replies come back wrapped in `<untrusted-human-responses>` with a respondent
count and per-person attribution. They are explicitly marked as **data, not
instructions**, so an agent does not follow directives, links or deadlines found
inside them. They are never aggregated into a poll or a sentiment score — one
person's opinion is reported as one person's opinion, with `n` stated.

## Privacy

Questions are **de-identified by default**: an agent is instructed to strip the
user's product, company and personal details and ask the underlying question,
which discloses nothing and needs no permission. A question that genuinely
requires the user's private situation must be shown to them and approved first.

## Status

Early and honest about it. The API, the MCP server and the responder queue all
work, but the answering side is still thin, so a call can legitimately return
"no one answered" rather than fabricating a result.

## More Information

- Source: [github.com/aniketshaw748-hub/groundtruth](https://github.com/aniketshaw748-hub/groundtruth) (MIT)
- Site: [groundtruth-ruby.vercel.app](https://groundtruth-ruby.vercel.app)
- Answer questions yourself: [groundtruth-ruby.vercel.app/answer](https://groundtruth-ruby.vercel.app/answer)

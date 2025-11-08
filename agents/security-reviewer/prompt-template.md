# Docker MCP Security Review Instructions

$MODE_SUMMARY

Security review metadata:
- Mode: $MODE_LABEL
- Repository name: $TARGET_LABEL
- Repository path: $REPOSITORY_PATH
- Head commit: $HEAD_COMMIT
- Base commit: $BASE_COMMIT
- Commit range: $COMMIT_RANGE

Core analysis directives:
- $GIT_DIFF_HINT
- Hunt aggressively for intentionally malicious behavior (exfiltration,
  persistence, destructive payloads) in addition to accidental security bugs.
- Evaluate credential handling, network access, privilege changes, supply chain
  touch points, and misuse of sensitive APIs.
- Use only the tools provided (git, ripgrep, jq, etc.); outbound network access
  is unavailable.
- Keep any files you create within /workspace or $REPOSITORY_PATH.

Mode-specific focus:
- Differential: Map every risky change introduced in the commit range. Call out
  suspicious files, shell commands, or configuration shifts. Note beneficial
  security hardening too.
- Full: Examine the entire repository snapshot. Highlight persistence vectors,
  secrets handling, dependency risk, and opportunities for escalation.

Report expectations:
- Reproduce every heading, section order, and field exactly as written in
  `$REPORT_TEMPLATE_PATH`; replace bracketed placeholders with concrete
  content but do not add or remove sections.
- Save the final report to $REPORT_PATH.
- Articulate severity, impact, evidence, and remediation for each issue.

Labeling guidance:
- Write labels to $LABELS_PATH, one per line.
- Emit exactly one overall risk label in the form `security-risk:<level>` where
  `<level>` is one of `critical`, `high`, `medium`, `low`, or `info`.
- Align the chosen label with the overall risk level declared in the report.
- If you identify blocking or critical issues that must halt release, also
  include the label `security-blocked` on a separate line.
- Leave $LABELS_PATH empty only if the review cannot be completed.

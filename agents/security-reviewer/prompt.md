# Docker MCP Security Review Instructions

Mode: $MODE_LABEL

$MODE_SUMMARY

Repository metadata:
- Target label: $TARGET_LABEL
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
- Structure your findings using `/opt/security-reviewer/report-template.md`.
- Save the final report to $REPORT_PATH.
- Articulate severity, impact, evidence, and remediation for each issue.

Labeling guidance:
- Write labels to $LABELS_PATH, one per line.
- Emit exactly one overall risk label in the form `security-risk:<level>` where
  `<level>` is one of `critical`, `high`, `medium`, `low`, or `info`.
- Align the chosen label with the overall risk level declared in the report.
- Leave $LABELS_PATH empty only if the review cannot be completed.

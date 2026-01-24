# Security Review Report

## Scenario Classification
- **Scenario Type:** [Pin Upgrade PR / New Server Onboarding / Material Change]
- **Justification:** [Explain which scenario applies based on concrete evidence from the diff. For pin upgrades, state if changes are limited to version/commit fields. For material changes, list what execution scope changed (permissions, env vars, network config, etc.)]

## Scope Summary
- **Review Mode:** [Differential/Full]
- **Repository:** [Repository name]
- **Head Commit:** [SHA]
- **Base Commit:** [SHA or N/A]
- **Commit Range:** [base...head or single commit]
- **Check Conclusion:** [success/failure]
- **Overall Risk Level:** [CRITICAL/HIGH/MEDIUM/LOW/INFO]

## Scope Context
This review focuses on **runtime MCP server security** for Docker registry distribution:
- **IN SCOPE:** Runtime MCP server code (tools, handlers, business logic exposed via MCP interface)
- **OUT OF SCOPE:** CI/CD workflows, pre-commit hooks, contributor build scripts, test infrastructure
- **CONTEXT:** Docker rebuilds images from source - contributor build hygiene issues do not affect distribution security

## Executive Summary
[One to two paragraphs summarizing:
1. What changed (for differential) or what the server does (for full)
2. Security Bar assessment: Does this violate any blocking categories (A-E)?
3. Key security findings if any
4. Overall security posture and recommendation]

## Detailed Findings

[ONLY include findings that meet evidence requirements for their severity level]

### [Finding Title] — [CRITICAL/HIGH/MEDIUM/LOW/INFO]

- **Security Bar Category:** [A: Malicious Behavior / B: Credential Handling / C: Network Exfiltration / D: Privilege Escalation / E: Supply Chain / F: Access Control / N/A: Not a Security Bar violation]

- **Impact:** [Describe the risk to MCP server users or Docker ecosystem. Focus on runtime behavior, not build-time issues.]

- **Evidence:** [REQUIRED - Provide ALL of the following:
  - File path and line numbers (e.g., `src/tools/execute.py:45-52`)
  - Code snippet or command showing the issue
  - For pin upgrades: upstream diff lines that introduce the issue
  - Proof this is reachable via MCP tools or executes automatically at runtime
  - For CRITICAL/HIGH: Demonstration of Security Bar violation]

- **MCP Boundary Analysis:** [REQUIRED - Answer:
  - Is this reachable via MCP tool calls? [Yes/No - specify which tool]
  - Does it execute automatically when server runs? [Yes/No]
  - Is it in test/CI directories? [Yes/No - if yes, explain why it's still flagged]
  - Is it documented and opt-in? [Yes/No - link to documentation]]

- **Recommendation:** [Specific, actionable steps to remediate. For blocking issues, this is mandatory before merge.]

[Repeat for each finding]

## Precision Rule Filters Applied

[Document which precision rules were applied to avoid false positives:]
- Rule 1 (Trace MCP Boundary): [Verified all findings are reachable via MCP or auto-run]
- Rule 2 (Documentation Override): [Checked if flagged behaviors are documented and opt-in]
- Rule 3 (Test/CI Code): [Excluded findings in test/, .github/workflows/, etc.]
- Rule 7 (Credential Placeholders): [Verified credentials are real, not examples]
- [List other rules applied]

## Defense-In-Depth Observations

**Positive Security Patterns:**
- [List security controls present in the codebase]
- [Note hardening measures, validation, least privilege, etc.]

**Remaining Gaps (Non-Blocking):**
- [List defense-in-depth improvements that don't violate Security Bar]
- [These are recommendations, not blockers]

## Recommended Labels

[REQUIRED - Exactly ONE risk label + blocking label if applicable]

**Risk Level Label:**
- `security-risk:[critical/high/medium/low/info]` — [Justify based on highest severity finding]

**Blocking Label (if applicable):**
- `security-blocked` — [ONLY if ANY finding violates Security Bar categories A-E with required evidence]

**Reasoning:**
[Explain why this risk level was chosen and whether it blocks merge per the Security Bar]

## Check Conclusion Determination

**Conclusion:** `[success/failure]`

**Reasoning:**
- [IF failure] One or more findings violate Security Bar categories [list which: A/B/C/D/E]
- [IF success] No Security Bar violations detected. Report may contain INFO/LOW/MEDIUM recommendations.

**Auto-Merge Eligibility:**
- [IF success] Eligible for auto-merge (if server in allowlist and other requirements met)
- [IF failure] Requires manual review and remediation before merge

## Next Steps

[ONLY for blocking issues]
1. [Specific remediation action with file/line references]
2. [Verification steps]
3. [Timeline or priority if known]

[For non-blocking issues]
- [Optional recommendations for future improvements]

## Conclusion

[Final summary including:
1. Overall security assessment
2. Whether this passes or fails Security Bar
3. For pin upgrades: whether upstream changes are safe
4. For new servers: whether server is safe for Docker registry
5. Any critical context for reviewers]

---

## Security Bar Reference

**This review evaluated the following blocking categories:**
- **A. Malicious or Deceptive Behavior:** Exfiltration, backdoors, obfuscation, hidden behavior
- **B. Credential & Secret Handling:** Hardcoded secrets, credential leaks, insecure transmission
- **C. Network Access & Exfiltration:** Undocumented calls, unauthorized data uploads
- **D. Privilege & Host Interaction:** Escalation, host filesystem access, Docker socket access
- **E. Supply Chain Risk:** curl|bash, untrusted sources, self-updating code

**Non-blocking (reported but not blocking):**
- **F. Access Control & Data Exposure:** Metadata exposure, missing rate limits, overly broad access

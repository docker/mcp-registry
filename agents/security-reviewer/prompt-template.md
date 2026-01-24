# Docker MCP Security Review Instructions

$MODE_SUMMARY

## Security Review Metadata

- Mode: $MODE_LABEL
- Repository name: $TARGET_LABEL
- Repository path: $REPOSITORY_PATH
- Head commit: $HEAD_COMMIT
- Base commit: $BASE_COMMIT
- Commit range: $COMMIT_RANGE

---

## 1. MCP Security Design Principles

### Principle 1: MCP servers are untrusted by default

Docker does not assume:
- The MCP server is well-maintained
- The author is security-savvy
- The environment is hardened

The bar is therefore about **preventing ecosystem harm**, not enforcing perfect application security.

### Principle 2: Docker is responsible for distribution safety, not application correctness

**Docker's responsibility:**
- Preventing malicious behavior
- Preventing accidental credential leakage
- Preventing supply-chain abuse
- Preventing unexpected network or privilege behavior

**NOT Docker's responsibility:**
- Business logic vulnerabilities
- Performance issues
- Misconfiguration risks that are clearly documented and opt-in

### Principle 3: Intent matters more than theoretical exploitability

- **Intentionally malicious or deceptive behavior** → hard stop
- **Clearly accidental but dangerous patterns** → warn/block depending on severity
- **Theoretical risks without realistic exploitation paths** → informational

---

## 2. MCP Threat Model

### In Scope (Blocking Issues)

- Credential exfiltration (tokens, SSH keys, cloud creds)
- Backdoors / persistence mechanisms
- Hidden network callbacks
- Supply-chain abuse
- Privilege escalation on host
- Abuse of MCP client trust (tools doing unexpected things)

### Out of Scope (Non-Blocking)

- CI/CD workflow hygiene issues
- Pre-commit hook configuration
- Dockerfile issues (Docker rebuilds from source)
- SQL injection in internal queries (unless exploitable via MCP interface)
- Denial of service from expensive queries
- Overly permissive internal APIs
- Poor secrets documentation (vs actual leaks)
- Test credentials in CI/test environments

---

## 3. The MCP Security Bar (Normative Policy)

An MCP server **MUST NOT** be included in the Docker MCP Registry if it violates any requirement below.

### A. Malicious or Deceptive Behavior (BLOCKING)

❌ **Disallowed:**
- Code that intentionally exfiltrates data (files, env vars, tokens, prompts)
- Establishes persistence outside the MCP process lifecycle
- Executes destructive actions unrelated to declared functionality
- Obfuscated logic intended to hide behavior
- Undocumented background services, daemons, cron jobs
- Hidden network destinations or encrypted payloads without explanation

✅ **Allowed:**
- Explicit, documented network calls required for functionality
- Telemetry only if clearly documented and opt-in

### B. Credential & Secret Handling (BLOCKING / HIGH)

❌ **Disallowed:**
- Hardcoded real credentials, API keys, tokens
- Code that reads secrets and sends them over the network without explicit user intent
- Logging secrets or including them in error messages
- Using credentials implicitly (e.g., auto-loading cloud creds without user opt-in)

⚠️ **Conditionally Allowed (Test credentials):**
- Clearly labeled as test-only
- Scoped to local/CI usage
- Not transmitted externally
- Not copy-pasteable into production without modification

✅ **Allowed:**
- Use of env vars / secret stores
- Mock or placeholder credentials (example, changeme, etc.)
- Standard SDK credential chains (AWS SDK default chain, GCP ADC, Azure DefaultCredential) when the tool's stated purpose requires cloud access

### C. Network Access & Data Exfiltration (BLOCKING / HIGH)

❌ **Disallowed:**
- Undocumented outbound network calls
- Calls to endpoints derived from runtime data, attacker-controlled DNS, or downloaded content
- Uploading user data without explicit user command

⚠️ **Conditionally Allowed:**
- Calls to third-party APIs only if:
  - Core to MCP server functionality
  - Clearly documented
  - User-initiated

✅ **Allowed:**
- Localhost-only communication
- Read-only API access for declared purposes
- Endpoints explicitly configured by user via environment variables or config files

### D. Privilege & Host Interaction (BLOCKING)

❌ **Disallowed:**
- Requires root or escalated privileges without justification
- Modifies host filesystem outside declared scope
- Executes shell commands unrelated to MCP server purpose
- Accesses Docker socket, SSH agent, kubeconfig without documentation

⚠️ **Conditionally Allowed (with justification):**
- Root/escalation only if:
  - Core functionality strictly requires it (e.g., Docker management tool)
  - Clearly documented with security warnings
  - Alternatives discussed/dismissed
  - Scope is minimized

✅ **Allowed:**
- Explicit, documented host interactions
- Read-only inspection tools

### E. Supply Chain & Dependency Risk (HIGH / MEDIUM)

❌ **Disallowed:**
- Executing remote code during install without user awareness (curl | bash)
- Pulling dependencies from untrusted or obfuscated sources
- Self-updating code

⚠️ **Conditionally Allowed:**
- Large dependency trees (normal)
- Native extensions (if documented)

✅ **Allowed:**
- Pinned dependencies
- Standard package managers (npm, PyPI, Maven Central, crates.io)

### F. Access Control & Data Exposure (MEDIUM / LOW)

⚠️ **Non-blocking but should be flagged:**
- Overly broad access to internal metadata
- Exposing sensitive data via MCP tools without auth checks
- Missing rate limits

✅ **Allowed:**
- Tools intended for monitoring/debugging
- Read-only introspection

---

## 4. Severity Mapping

| Finding Type | Severity | Blocks Registry? |
|---|---|---|
| Malicious behavior | Critical | ✅ Yes |
| Real credential leaks | Critical | ✅ Yes |
| Backdoor/persistence | Critical | ✅ Yes |
| Obfuscated logic (malicious intent) | Critical | ✅ Yes |
| Undocumented network exfiltration | High | ✅ Yes |
| Privilege escalation (host) | High | ✅ Yes |
| Supply chain abuse (curl\|bash) | High | ✅ Yes |
| Self-updating code | High | ✅ Yes |
| Docker socket access (unauthorized) | High | ✅ Yes |
| Test creds in CI only | Medium | ❌ No |
| Metadata exposure (documented) | Medium | ❌ No |
| Missing rate limits | Low | ❌ No |
| Theoretical injection (sanitized) | Info | ❌ No |

---

## 5. Precision Rules (Minimize False Positives)

Apply these filters BEFORE classifying any finding as medium or higher:

### Rule 1: Trace the MCP Boundary

**DO NOT FLAG** if the code path is not reachable via MCP tools.

A vulnerability is only in-scope if:
- User can trigger it via an MCP tool call, OR
- It executes automatically when the server starts/runs

Internal functions, admin scripts, dev tooling, migration scripts → **OUT OF SCOPE**

### Rule 2: Documentation Overrides Suspicion

**DO NOT FLAG** if behavior is:
- Explicitly documented in README/docs
- Matches the stated purpose of the MCP server
- User explicitly opted-in via configuration

**Exception:** Documentation does NOT override Security Bar categories A-C violations. Documenting malicious behavior (exfiltration, persistence, hidden callbacks) does not make it acceptable.

### Rule 3: Test/CI Code is Presumed Safe

**DO NOT FLAG** issues in:
- Directories: `test/`, `tests/`, `__tests__/`, `spec/`, `.github/workflows/`
- Files: `*_test.*`, `*_spec.*`, `*.test.*`, `*.spec.*`
- CI configs: `.github/workflows/*`, `.gitlab-ci.yml`, `.circleci/*`

**Exception:** Block if test code could be executed in production

### Rule 4: Localhost/Loopback is Presumed Safe

**DO NOT FLAG** network calls to:
- `127.0.0.1`, `localhost`, `0.0.0.0` (when binding)
- `::1` (IPv6 loopback)
- Unix sockets

**Exception:** Flag if localhost access targets cloud metadata endpoints (`169.254.169.254`, `metadata.google.internal`, `metadata.aws.internal`) or Docker socket (`/var/run/docker.sock`) without documentation.

### Rule 5: Read-Only ≠ Exfiltration

**DO NOT FLAG** code that reads sensitive data UNLESS it also sends it.

Reading alone is not a vulnerability if:
- Data stays in-process or goes to user's client via MCP response
- No network transmission to third parties
- Explicitly part of the tool's purpose

### Rule 6: Standard Dependency Patterns are Safe

**DO NOT FLAG:**
- Large dependency trees
- Well-known packages from standard registries with established history
- Pinned versions in package manifests
- Post-install scripts from well-established packages

**DO FLAG (as MEDIUM, non-blocking):**
- Packages with names very similar to popular packages (potential typosquatting)
- Dependencies on packages with minimal adoption or no visible community
- Post-install scripts that execute network calls or download additional code

### Rule 7: Credential Placeholders are Not Leaks

**DO NOT FLAG** obviously fake credentials:
- `example`, `changeme`, `your-token-here`, `<YOUR_KEY>`, `TODO`, `xxxxxxxx`
- Empty strings, null, undefined
- Test values clearly labeled (`MOCK_*`, `TEST_*`, `EXAMPLE_*`)

### Rule 8: Framework Patterns are Not Backdoors

**DO NOT FLAG** standard framework features:
- Express/Flask/FastAPI request handlers
- Django admin interfaces
- Standard logging/metrics libraries
- Health check endpoints

### Rule 9: Dynamic Code Execution (Context Matters)

**DO NOT FLAG** eval/exec if:
- Used for config parsing (YAML, JSON evaluation)
- Template rendering (Jinja, handlebars)
- Sandboxed REPL tools (where that's the entire purpose)
- Code analysis tools (linters, formatters)

**DO FLAG** eval/exec if:
- Takes unsanitized user input via MCP tools
- Executes downloaded code
- No sandboxing when accepting arbitrary code

### Rule 10: Privilege Escalation Requires Host Impact

**DO NOT FLAG:**
- In-container privilege drops (running as non-root inside container)
- Capability requests if properly documented
- Standard container entrypoint patterns

### Rule 11: Theoretical vs Exploitable

**DO NOT FLAG** if ALL of these are true:
- No attacker-controlled input via MCP tools
- Would require chaining multiple vulnerabilities
- Requires specific environment configuration unlikely in practice
- No evidence of malicious intent

---

## 6. Evidence Requirements

To classify a finding at a given severity, provide ALL required evidence:

### Critical (security-blocked)
1. Exact file path and line numbers
2. Proof of malicious intent OR credential compromise OR unauthorized exfiltration
3. Demonstration that it's reachable via MCP interface or runs automatically
4. Explanation of impact (what gets stolen/destroyed/accessed)

### High (security-blocked)
1. Exact file path and line numbers
2. Violation of Security Bar categories A-E
3. Trace from MCP tool to vulnerable code path
4. Explanation of realistic exploitation

### Medium (NOT blocking)
1. File path
2. Description of the issue
3. Why it doesn't rise to High

### Low/Info
1. Brief description
2. Recommendation (optional)

---

## 7. Scenario Detection (REQUIRED)

This check may be running in one of two scenarios:

**A) Pin upgrade PR (most common):**
- Docker/mcp-registry changes only a pinned commit/version reference for an upstream MCP server
- Expect the diff to be small and mainly update a server entry (e.g., source.commit / version pin)

**B) New MCP server onboarding (manual/full):**
- A new server entry is being added, or a server definition is changing materially (permissions, env, image/build settings, entrypoint/cmd, network behavior)

**You MUST determine which scenario applies by inspecting the diff contents:**
- If changes are limited to pin/version/commit fields in an existing server entry and do not change execution scope, classify as scenario (A) Pin upgrade PR
- If the PR adds a new server entry OR changes execution scope (permissions, env vars, secrets, image/build, commands, network, volumes, host access), classify as scenario (B) New onboarding/material change

**In your report, explicitly state which scenario you determined and why, using concrete evidence from the diff.**

### Pin Upgrade Exception Rule

For PRs that ONLY update a version/commit pin:
- Apply Security Bar to the DIFF, not the entire codebase
- Block only if the pin update introduces a Security Bar violation
- If upstream diff cannot be fetched, default to MEDIUM risk (cannot verify)
- Changes to permissions, env vars, network config, or entrypoint → treat as NEW server

**Evidence required:**
- Show the specific lines from the upstream diff that violate the bar
- Don't block based on pre-existing issues unless they're being worsened

---

## 8. Scope Context

**Docker MCP Registry Scope:**
- **IN SCOPE:** Runtime MCP server code (tools, handlers, business logic)
- **OUT OF SCOPE:** CI/CD workflows, pre-commit hooks, build scripts, test infrastructure
- **CONTEXT:** Docker rebuilds all images from source - focus on runtime security, not build hygiene

### The "Docker Rebuilds Anyway" Principle

Many findings in Dockerfiles and build scripts are moot because:
- Docker builds images from source
- Docker controls the final Dockerfile
- Upstream build issues ≠ distribution issues
- Focus should be on **runtime MCP server behavior**

---

## 9. False Positive Decision Tree

For EVERY finding, ask these questions IN ORDER:

1. Is it reachable via MCP tools or auto-runs?
   → NO = Not a finding (internal code)

2. Is it documented and matches stated purpose?
   → YES = Not a finding (expected behavior)

3. Is it in test/CI directories or files?
   → YES = Not a finding (unless provably runs in production)

4. Is it reading data but not sending anywhere?
   → YES = Not a finding (unless it exfiltrates)

5. Is it using standard libraries/patterns correctly?
   → YES = Not a finding (framework feature)

6. Is the credential obviously fake/placeholder?
   → YES = Not a finding (example value)

7. Is it localhost/loopback communication?
   → YES = Not a finding (unless proxying externally)

8. Does it require multiple unlikely conditions to exploit?
   → YES = Info at most (theoretical)

9. Can you provide ALL required evidence for the severity level?
   → NO = Downgrade severity OR mark as "requires human review"

**If it passes ALL 9 filters → legitimate finding**

---

## 10. Blocking Label Rules

The `security-blocked` label determines whether a PR can proceed:
- If emitted → PR fails automated checks, requires manual review
- If not emitted → PR can proceed (may auto-merge if other criteria met)

**Emit `security-blocked` label** IF:
- ANY finding violates Security Bar categories A-E with required evidence

**Do NOT emit `security-blocked`** IF:
- No Security Bar violations found
- Only INFO/LOW/MEDIUM findings present

---

## Core Analysis Directives

$GIT_DIFF_HINT

- Hunt aggressively for intentionally malicious behavior (exfiltration, persistence, destructive payloads) in addition to accidental security bugs
- Evaluate credential handling, network access, privilege changes, supply chain touch points, and misuse of sensitive APIs
- Use only the tools provided (git, ripgrep, jq, etc.); outbound network access is unavailable
- Keep any files you create within /workspace or $REPOSITORY_PATH

### Mode-Specific Focus

**Differential:** Map every risky change introduced in the commit range. Call out suspicious files, shell commands, or configuration shifts. Note beneficial security hardening too.

**Full:** Examine the entire repository snapshot. Highlight persistence vectors, secrets handling, dependency risk, and opportunities for escalation.

---

## Report Expectations

- Reproduce every heading, section order, and field exactly as written in `$REPORT_TEMPLATE_PATH`; replace bracketed placeholders with concrete content but do not add or remove sections
- Save the final report to $REPORT_PATH
- Articulate severity, impact, evidence, and remediation for each issue
- **Explicitly state scenario classification** (Pin Upgrade vs New Server) with evidence
- Apply all precision rules and document which rules filtered out potential false positives
- Provide complete evidence per severity requirements

---

## Labeling Guidance

- Write labels to $LABELS_PATH, one per line
- Emit exactly one overall risk label in the form `security-risk:<level>` where `<level>` is one of `critical`, `high`, `medium`, `low`, or `info`
- Align the chosen label with the overall risk level declared in the report
- If you identify blocking issues that violate Security Bar categories A-E, also include the label `security-blocked` on a separate line
- Leave $LABELS_PATH empty only if the review cannot be completed
- Note: The risk label and blocking label serve different purposes. A report can have `security-risk:medium` without `security-blocked` if no Security Bar violations exist

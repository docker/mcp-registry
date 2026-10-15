Retasc is a work queue for AI coding agents that Claude Code, Codex, Cursor and other MCP clients pull from when you have several agents to coordinate.

Two agents can never hold the same issue. A claim is one server transaction. `next_issue` claims the highest effective-priority issue that is unblocked and unclaimed. When the CLI sees a dirty worktree or a commit from the last 24 hours on an unclaimed issue's branch, Retasc holds it out of dispatch for 60 minutes so a second agent does not start the same work. A claim replaces the hold.

Two of your agents try to file the same bug. Retasc refuses a create that looks like open work and names the look-alikes. The check uses issue IDs mentioned in the new title or body and rare words shared with open issue titles in the project. The agent comments on or claims the existing issue, or re-files naming the issues it read.

Blocked work is never handed out. Every issue an agent creates must declare what it is blocked by, or state in one line why it depends on nothing. Closing a blocker frees its dependents. A chore blocking an urgent feature jumps the queue. You get a message on Slack, Discord, Telegram or email when your issue's last open blocker closes.

Rate limits, closed laptops and stalled sessions can stop your agent. The Retasc CLI renews each 30-minute lease every 10 minutes without involving the model. When renewals stop, the issue returns to the queue in its owner's lane. That person's next agent, on any runtime, resumes from the checkpoint of what's done, what's next and the gotchas. If nobody returns for 7 days, it rejoins the shared pool.

Work waiting for review jumps the queue. Retasc requires a note naming the branch, PR or commit unless the agent already left a checkpoint. A different person's agent gets the note. Whoever sent the work to review cannot claim or accept its review. Review ranks at least High immediately, Urgent at 21.6 hours and above everything after 24 hours.

GitHub and GitLab issues reach Retasc through verified webhooks. No agent gets a synced issue until a member approves its exact text. A source edit voids the approval. The Dash lists every agent, its runtime, the issue it holds and the minutes left on its lease.

`npx @retasc/cli@latest bind` wires MCP for Claude Code, Codex, Cursor, OpenCode, Gemini CLI, Grok and Pi. Your folder gets an agent key after browser approval if you're signed out. Start or restart your agent there to pull work. Billing is per action, with no seats. Members are free and unlimited. A new org gets $10 of usage at signup, no card.

Docs: https://docs.retasc.com

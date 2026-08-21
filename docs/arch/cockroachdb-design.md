# CockroachDB Integration Design (MCP Registry)

This document summarizes the design choices for integrating CockroachDB as the durable state backend for MCP server artifacts.

Key points:
- Persist authoritative state (idempotency, metadata) in CRDB to leverage Raft replication and leader election per range.
- Use Serializable MVCC transactions and bounded retry loops (crdb.ExecuteTx) to provide strict serializability and safe concurrent execution.
- Partition hot metadata by owner/region to keep transactions single-range and reduce leaseholder transfers.
- Keep application servers stateless; rely on DB for durable locks/leases; implement graceful drain and health probes for rolling updates.

Operational notes:
- Migrations and schema evolution scripts belong with the server implementation repository and CI pipeline.- Run DB migrations at deploy time with a safe migrator (idempotent SQL files).
- Configure DB connection via env var COCKROACH_URLS and tune pool sizes.

For reviewers: this approach grounds in CockroachDB literature: Raft per range, MVCC + Serializable snapshot isolation, and leaseholder-aware locality are explicitly used to avoid split-brain, write skew, and high-latency cross-region transactions.


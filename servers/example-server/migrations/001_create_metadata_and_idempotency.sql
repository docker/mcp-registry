-- migrations/001_create_metadata_and_idempotency.sql
CREATE TABLE IF NOT EXISTS tool_metadata (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  owner STRING NOT NULL,
  key STRING NOT NULL,
  value JSONB,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  INDEX (owner),
  FAMILY (owner, key, value, created_at)
);

-- Idempotency table: one-row-per-invocation, use SERIALIZABLE transactions to upsert and return stored result.
CREATE TABLE IF NOT EXISTS idempotency (
  invocation_key STRING PRIMARY KEY, -- e.g. sha256(client+tool+params+timestamp_bucket)
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  status STRING, -- "in_progress","done","failed"
  response JSONB NULL,
  expire_at TIMESTAMP WITH TIME ZONE
);

-- Add TTL for cleanup (uses scheduled job or process sweep)

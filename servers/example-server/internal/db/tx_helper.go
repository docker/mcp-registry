package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
)

var ErrInvocationInProgress = errors.New("invocation in progress")

// ErrAlreadyDone is returned to surface an already-completed invocation and carry the stored response.
type ErrAlreadyDone struct{
	Response []byte
}

func (e ErrAlreadyDone) Error() string { return "already done" }

// ExecuteIdempotentInvocation runs a transaction that records an invocation marker
// and stores the result atomically. doWork should perform DB-safe work and return
// a JSON-marshallable response.
func ExecuteIdempotentInvocation(ctx context.Context, db *sql.DB, invocationKey string, doWork func(tx *sql.Tx) (json.RawMessage, error)) (json.RawMessage, error) {
	attempts := 5
	backoff := 100 * time.Millisecond
	for i := 0; i < attempts; i++ {
		err := crdb.ExecuteTx(ctx, db, nil, func(tx *sql.Tx) error {
			// Check existing marker
			var status sql.NullString
			var respBytes []byte
			err := tx.QueryRowContext(ctx, "SELECT status, response FROM idempotency WHERE invocation_key = $1", invocationKey).Scan(&status, &respBytes)
			if err == nil {
				if status.String == "done" {
					return ErrAlreadyDone{Response: respBytes}
				}
				return ErrInvocationInProgress
			}
			if err != sql.ErrNoRows && err != nil {
				return err
			}

			// Insert marker
			if _, err := tx.ExecContext(ctx, "INSERT INTO idempotency (invocation_key, status, created_at, expire_at) VALUES ($1, 'in_progress', now(), now() + interval '24 hours')", invocationKey); err != nil {
				return err
			}

			// perform the work
			resp, err := doWork(tx)
			if err != nil {
				_, _ = tx.ExecContext(ctx, "UPDATE idempotency SET status='failed' WHERE invocation_key=$1", invocationKey)
				return err
			}

			// store result
			_, err = tx.ExecContext(ctx, "UPDATE idempotency SET status='done', response=$2 WHERE invocation_key=$1", invocationKey, resp)
			return err
		})

		if err == nil {
			var out []byte
			_ = db.QueryRowContext(ctx, "SELECT response FROM idempotency WHERE invocation_key=$1", invocationKey).Scan(&out)
			return out, nil
		}
		if alreadyDone, ok := err.(ErrAlreadyDone); ok {
			return alreadyDone.Response, nil
		}
		// retryable -> sleep
		time.Sleep(backoff)
		backoff *= 2
	}
	return nil, errors.New("transaction retries exhausted")
}

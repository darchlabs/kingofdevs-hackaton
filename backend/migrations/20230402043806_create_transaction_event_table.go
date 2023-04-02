package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTransactionEventTable, downCreateTransactionEventTable)
}

func upCreateTransactionEventTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS transaction_event (
			id TEXT PRIMARY KEY,
			transaction_id TEXT NOT NULL,
			event_id TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL
		)
	`)

	return err
}

func downCreateTransactionEventTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS transaction_event
	`)

	return err
}

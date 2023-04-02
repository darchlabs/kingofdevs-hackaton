package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTransactionTable, downCreateTransactionTable)
}

func upCreateTransactionTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE transaction (
			id TEXT NOT NULL PRIMARY KEY,
			tx TEXT NOT NULL,
			from_addr TEXT NOT NULL,
			from_balance decimal(38, 0) NOT NULL,
			contract_balance decimal(38, 0) NOT NULL,
			gas_paid bigint NOT NULL,
			gas_price decimal(38, 0) NOT NULL,
			gas_cost decimal(38, 0) NOT NULL,
			from_is_whale boolean NOT NULL,
			tx_succeded boolean NOT NULL DEFAULT false,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL
		)
	`)

	return err
}

func downCreateTransactionTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS transaction
	`)

	return err
}

package transaction

import (
	"time"
)

type Transaction struct {
	ID              string    `json:"id" db:"id"`
	ContractAddr    string    `json:"contract_addr" db:"contract_addr"`
	Tx              string    `json:"tx" db:"tx"`
	FromAddr        string    `json:"from_addr" db:"from_addr"`
	FromBalance     string    `json:"from_balance" db:"from_balance"`
	ContractBalance string    `json:"contract_balance" db:"contract_balance"`
	GasPaid         string    `json:"gas_paid" db:"gas_paid"`
	GasPrice        string    `json:"gas_price" db:"gas_price"`
	GasCost         string    `json:"gas_cost" db:"gas_cost"`
	FromIsWhale     bool      `json:"from_is_whale" db:"from_is_whale"`
	TxSucceded      bool      `json:"tx_succeded" db:"tx_succeded"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

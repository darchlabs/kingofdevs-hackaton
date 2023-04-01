package transaction

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	TX              common.Hash    `json:"tx" db:"tx"`
	From            common.Address `json:"from" db:"from"`
	FromBalance     *big.Int       `json:"from_balance" db:"from_balance"`
	ContractBalance *big.Int       `json:"contract_balance" db:"contract_balance"`
	GasPaid         uint64         `json:"gas_paid" db:"gas_paid"`
	GasPrice        *big.Int       `json:"gas_price" db:"gas_price"`
	GasCost         *big.Int       `json:"gas_cost" db:"gas_cost"`
	FromIsWhale     bool           `json:"from_is_whale" db:"from_is_whale"`
	TXSucceded      bool           `json:"tx_succeded" db:"tx_succeded"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

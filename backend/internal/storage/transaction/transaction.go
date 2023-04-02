package transactionstorage

import (
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/transaction"
)

type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage{
		storage: s,
	}
}

func (s *Storage) ListTxs() ([]*transaction.Transaction, error) {
	// define events response
	txs := []*transaction.Transaction{}

	// get txs from db
	eventQuery := "SELECT * FROM transaction"
	err := s.storage.DB.Select(&txs, eventQuery)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (s *Storage) GetTxById(id string) (*transaction.Transaction, error) {
	// define events response
	tx := transaction.Transaction{}

	// get txs from db
	eventQuery := "SELECT * FROM transaction WHERE id = $1"
	err := s.storage.DB.Get(&tx, eventQuery, id)
	if err != nil {
		return nil, err
	}

	return &tx, nil

}

func (s *Storage) ListCurrentHashes() (*[]string, error) {
	// define events response
	var hashesArr []string

	// get txs from db
	eventQuery := "SELECT tx FROM transaction"
	err := s.storage.DB.Get(&hashesArr, eventQuery)
	if err != nil {
		return nil, err
	}

	return &hashesArr, nil
}

func (s *Storage) GetTVL(contractAddr string) (*int64, error) {
	// define events response
	var tvl int64

	// get txs from db
	eventQuery := "SELECT SUM(contract_balance) FROM transaction WHERE address = $1"
	err := s.storage.DB.Get(&tvl, eventQuery, contractAddr)
	if err != nil {
		return nil, err
	}

	return &tvl, nil

}

func (s *Storage) ListTotalAddresses(contractAddr string) (*int64, error) {
	// define events response
	var totalAddr int64

	query := "SELECT COUNT(DISTINCT address) FROM transaction WHERE contract_addr = $1"

	// execute query and retrieve result
	err := s.storage.DB.Get(&totalAddr, query, contractAddr)
	if err != nil {
		return nil, err
	}

	return &totalAddr, nil
}
func (s *Storage) InsertTx(t *transaction.Transaction) (*transaction.Transaction, error) {
	// check if already existe an event with the same address and name
	tx, _ := s.GetTxById(t.ID)
	if tx != nil {
		return nil, fmt.Errorf("transaction already exists with hash=%s", t.Tx)
	}

	eventQuery := "INSERT INTO transaction (id, contract_addr, tx, from_addr, from_balance, contract_balance, gas_paid, gas_price, gas_cost, from_is_whale, tx_succeded, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
	_, err := s.storage.DB.Exec(eventQuery, t.ID, t.ContractAddr, t.Tx, t.FromAddr, t.FromBalance, t.ContractBalance, t.GasPaid, t.GasPrice, t.GasCost, t.FromIsWhale, t.TxSucceded, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get created event
	created, err := s.GetTxById(t.ID)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Storage) UpdateTx(t *transaction.Transaction) (*transaction.Transaction, error) {
	// check if already existe an event with the same address and name
	ev, err := s.GetTxById(t.Tx)
	if err != nil {
		return nil, err
	}

	if ev == nil {
		return nil, fmt.Errorf("%s", "tx does not exists on the db")
	}

	// update tx on db
	query := "UPDATE transaction SET network = $1, node_url = $2, address = $3, latest_block_number = $4, abi_id = $5, status = $6, error = $7, updated_at = $8 WHERE id = $9"
	_, err = s.storage.DB.Exec(query, t.Tx, t.FromAddr, t.FromBalance, t.ContractBalance, t.GasPaid, t.GasPrice, t.GasCost, t.FromIsWhale, t.TxSucceded, t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get created event
	created, err := s.GetTxById(t.Tx)
	if err != nil {
		return nil, err
	}

	return created, nil
}

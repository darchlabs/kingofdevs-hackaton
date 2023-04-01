package transactionstorage

import (
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/api/transaction"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	"github.com/ethereum/go-ethereum/common"
)

type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage{
		storage: s,
	}
}

func (s *Storage) ListAllTXs() ([]*transaction.Transaction, error) {
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

func (s *Storage) GetTXByHash(hash common.Hash) (*transaction.Transaction, error) {
	// define events response
	tx := &transaction.Transaction{}

	// get txs from db
	eventQuery := "SELECT * FROM transaction WHERE TX = $1"
	err := s.storage.DB.Select(&tx, eventQuery)
	if err != nil {
		return nil, err
	}

	return tx, nil

}

func (s *Storage) InsertTX(t *transaction.Transaction) (*transaction.Transaction, error) {
	// check if already existe an event with the same address and name
	ev, err := s.GetTXByHash(t.TX)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if ev != nil {
		return nil, fmt.Errorf("event already exists with this hash=%s", t.TX)
	}

	// prepare db for creating a tx on it
	tx, err := s.storage.DB.Beginx()
	if err != nil {
		return nil, err
	}

	// insert new event in database
	var txID string
	eventQuery := "INSERT INTO transaction (tx, from, from_balance, contract_balance, gas_paid, gas_price, gas_cost, from_is_whale, tx_succeded, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING tx"
	err = tx.Get(&txID, eventQuery, t.TX, t.From, t.FromBalance, t.ContractBalance, t.GasPaid, t.GasPrice, t.GasCost, t.FromIsWhale, t.TXSucceded, t.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// get created event
	createdTX, err := s.GetTXByHash(common.HexToHash(txID))
	if err != nil {
		return nil, err
	}

	return createdTX, nil
}

func (s *Storage) UpdateTX(t *transaction.Transaction) (*transaction.Transaction, error) {
	// check if already existe an event with the same address and name
	ev, err := s.GetTXByHash(t.TX)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	if ev == nil {
		return nil, fmt.Errorf("%s", "tx does not exists on the db")
	}

	// prepare db for creating a tx on it
	tx, err := s.storage.DB.Beginx()
	if err != nil {
		return nil, err
	}

	// update tx on db
	query := "UPDATE transaction SET network = $1, node_url = $2, address = $3, latest_block_number = $4, abi_id = $5, status = $6, error = $7, updated_at = $8 WHERE id = $9"
	_, err = tx.Exec(query, t.TX, t.From, t.FromBalance, t.ContractBalance, t.GasPaid, t.GasPrice, t.GasCost, t.FromIsWhale, t.TXSucceded, t.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// get created event
	createdTX, err := s.GetTXByHash(t.TX)
	if err != nil {
		return nil, err
	}

	return createdTX, nil
}

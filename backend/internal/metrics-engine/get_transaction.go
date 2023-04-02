package metricsengine

import (
	"context"
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (m *Metric) GetTransaction(txHash common.Hash) (*transaction.Transaction, error) {
	tx, _, err := m.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, err
	}

	fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return nil, err
	}

	fromBalance, err := m.client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return nil, err
	}

	// get contract balance
	contractBalance, err := m.client.BalanceAt(context.Background(), *tx.To(), nil)
	if err != nil {
		return nil, err
	}

	isWhale := false
	whaleLimit := int64(10000)
	if fromBalance.Int64() > whaleLimit {
		isWhale = true
	}

	fmt.Println("555", &transaction.Transaction{
		ID:              m.idGen(),
		Tx:              txHash.Hex(),
		From:            fromAddress.Hex(),
		FromBalance:     fromBalance,
		ContractBalance: contractBalance,
		GasPaid:         tx.Gas(),
		GasPrice:        tx.GasPrice(),
		GasCost:         tx.Cost(),
		FromIsWhale:     isWhale,
		CreatedAt:       m.dateGen(),
		UpdatedAt:       m.dateGen(),
	})

	// insert transaction in db
	t, err := m.transactionstorage.InsertTx(&transaction.Transaction{
		ID:              m.idGen(),
		Tx:              txHash.Hex(),
		From:            fromAddress.Hex(),
		FromBalance:     fromBalance,
		ContractBalance: contractBalance,
		GasPaid:         tx.Gas(),
		GasPrice:        tx.GasPrice(),
		GasCost:         tx.Cost(),
		FromIsWhale:     isWhale,
		CreatedAt:       m.dateGen(),
		UpdatedAt:       m.dateGen(),
	})
	if err != nil {
		return nil, err
	}

	return t, nil
}

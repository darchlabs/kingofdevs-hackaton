package metricsengine

import (
	"context"
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (m *Metric) GetTransaction(txHash string) (*transaction.Transaction, error) {
	tx, _, err := m.client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}

	contractAddress := tx.To().String()

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

	whaleLimit := int64(10000)
	isWhale := fromBalance.Int64() > whaleLimit

	// insert transaction in db
	t, err := m.transactionstorage.InsertTx(&transaction.Transaction{
		ID:              m.idGen(),
		ContractAddr:    contractAddress,
		Tx:              txHash,
		FromAddr:        fromAddress.Hex(),
		FromBalance:     fromBalance.String(),
		ContractBalance: contractBalance.String(),
		GasPaid:         fmt.Sprint(tx.Gas()),
		GasPrice:        tx.GasPrice().String(),
		GasCost:         tx.Cost().String(),
		FromIsWhale:     isWhale,
		TxSucceded:      true,
		CreatedAt:       m.dateGen(),
		UpdatedAt:       m.dateGen(),
	})
	if err != nil {
		return nil, err
	}

	return t, nil
}

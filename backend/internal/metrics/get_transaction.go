package metrics

import (
	"context"

	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransaction(client *ethclient.Client, txHash common.Hash) *transaction.Transaction {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		panic(err)
	}

	fromBalance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}

	contractBalance, err := client.BalanceAt(context.Background(), *tx.To(), nil)
	if err != nil {
		panic(err)
	}

	isWhale := false
	whaleLimit := int64(10000)
	if fromBalance.Int64() > whaleLimit {
		isWhale = true
	}

	return &transaction.Transaction{
		TX:              txHash,
		From:            fromAddress,
		FromBalance:     fromBalance,
		ContractBalance: contractBalance,
		GasPaid:         tx.Gas(),
		GasPrice:        tx.GasPrice(),
		GasCost:         tx.Cost(),
		FromIsWhale:     isWhale,
	}
}

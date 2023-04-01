package transaction

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetSender(client *ethclient.Client, txHash *common.Hash) {

	tx, _, err := client.TransactionByHash(context.Background(), *txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx: ", tx)

	fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sender address:", fromAddress.Hex())
}

package transaction

import (
	"fmt"

	eventstorage "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTx(s *eventstorage.Storage, client *ethclient.Client, contractAddress string) *common.Hash {

	events, err := s.ListEventsByAddress(contractAddress, "", 0, 0)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// make for loop iterating by event
	eventTX, err := s.ListEventTX(contractAddress, events[0].Abi.Name, "", 0, 0)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// make for loop iterating by tx hash
	txHash := common.HexToHash(*eventTX[0])
	return &txHash
}

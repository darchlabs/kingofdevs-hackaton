package metrics

import (
	"log"

	eventstorage "github.com/darchlabs/synchronizer-v2/pkg/storage/event"
	"github.com/ethereum/go-ethereum/common"
)

func GetTXHashesByAddress(evs *eventstorage.Storage, contractAddr string) []common.Hash {

	count, err := evs.GetEventCountByAddress(contractAddr)
	if err != nil {
		log.Fatal(err)
	}

	events, err := evs.ListEventsByAddress(contractAddr, "desc", count, 0)
	if err != nil {
		log.Fatal(err)
	}

	// iterar: obtener todos los event_datas por evento
	batch := 100
	offset := 0
	txHashArr := make([]common.Hash, 0)
	for _, e := range events {
		edCount, err := evs.GetEventDataCount(contractAddr, e.Abi.Name)
		if err != nil {
			log.Fatal(err)
		}

		// get all the data from an event
		for edCount > 0 {
			eventData, err := evs.ListEventData(contractAddr, e.Abi.Name, "desc", int64(batch), int64(offset))
			if err != nil {
				log.Fatal(err)
			}

			for _, data := range eventData {
				txHash := common.HexToHash(data.Tx)
				txHashArr = append(txHashArr, txHash)
			}
			edCount -= int64(batch)
		}

	}

	return txHashArr
}

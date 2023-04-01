package metricsengine

import (
	"github.com/ethereum/go-ethereum/common"
)

func (m *Metric) GetTXHashesByAddress(contractAddr string) ([]common.Hash, error) {

	count, err := m.eventStorage.GetEventCountByAddress(contractAddr)
	if err != nil {
		return nil, err
	}

	events, err := m.eventStorage.ListEventsByAddress(contractAddr, "desc", count, 0)
	if err != nil {
		return nil, err
	}

	// iterar: obtener todos los event_datas por evento
	batch := 100
	offset := 0
	txHashArr := make([]common.Hash, 0)
	for _, e := range events {
		edCount, err := m.eventStorage.GetEventDataCount(contractAddr, e.Abi.Name)
		if err != nil {
			return nil, err
		}

		// get all the data from an event
		for edCount > 0 {
			eventData, err := m.eventStorage.ListEventData(contractAddr, e.Abi.Name, "desc", int64(batch), int64(offset))
			if err != nil {
				return nil, err
			}

			for _, data := range eventData {
				txHash := common.HexToHash(data.Tx)
				txHashArr = append(txHashArr, txHash)
			}
			edCount -= int64(batch)
		}

	}

	return txHashArr, nil
}

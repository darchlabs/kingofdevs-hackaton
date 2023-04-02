package metricsengine

import (
	"github.com/ethereum/go-ethereum/common"
)

func (m *Metric) GetTXHashesByAddresses(contractAddresses []string) ([]string, error) {
	hashesArr := make([]string, 0)
	for _, address := range contractAddresses {
		count, err := m.eventStorage.GetEventCountByAddress(address)
		if err != nil {
			return nil, err
		}

		events, err := m.eventStorage.ListEventsByAddress(address, "desc", count, 0)
		if err != nil {
			return nil, err
		}

		// iterar: obtener todos los event_datas por evento
		batch := 100
		offset := 0
		for _, e := range events {
			edCount, err := m.eventStorage.GetEventDataCount(address, e.Abi.Name)
			if err != nil {
				return nil, err
			}

			// get all the data from an event
			for edCount > 0 {
				eventData, err := m.eventStorage.ListEventData(address, e.Abi.Name, "desc", int64(batch), int64(offset))
				if err != nil {
					return nil, err
				}

				for _, data := range eventData {
					txHash := common.HexToHash(data.Tx)
					hashesArr = append(hashesArr, txHash.Hex())
				}
				edCount -= int64(batch)
			}

		}
	}

	return hashesArr, nil
}

package metricsengine

import (
	"time"

	"github.com/darchlabs/kingofdevs-hackaton/backend"
	"github.com/darchlabs/synchronizer-v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

type idGenerator func() string
type dateGenerator func() time.Time

type Metric struct {
	client             *ethclient.Client
	eventStorage       synchronizer.EventStorage
	transactionstorage backend.TransactionStorage
	idGen              idGenerator
	dateGen            dateGenerator
}

func New(client *ethclient.Client, es synchronizer.EventStorage, ts backend.TransactionStorage, idGen idGenerator, dateGen dateGenerator) *Metric {
	return &Metric{
		client:             client,
		eventStorage:       es,
		transactionstorage: ts,
		idGen:              idGen,
		dateGen:            dateGen,
	}
}

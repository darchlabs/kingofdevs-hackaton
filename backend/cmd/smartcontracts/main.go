package main

import (
	"fmt"
	"log"

	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/api/metrics"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/env"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	transactionstorage "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage/transaction"
	"github.com/darchlabs/synchronizer-v2"
	eventdb "github.com/darchlabs/synchronizer-v2/pkg/storage"
	eventstorage "github.com/darchlabs/synchronizer-v2/pkg/storage/event"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
)

var (
	eventStorage synchronizer.EventStorage
)

func main() {
	// load env values
	var env env.Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	// initialize event storage
	edb, err := eventdb.New(env.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	// initialize event storage
	evs := eventstorage.New(edb)

	// TODO: is this ok?
	// TODO: How to create the table?
	s, err := storage.New(env.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	eventStorage = eventstorage.New(edb)
	txs := transactionstorage.New(s)

	// // run migrations
	// err = goose.Up(s.DB.DB, env.MigrationDir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Instance client
	client, err := ethclient.Dial(env.ClientURL)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	// Get tx hash
	txHashArr := metrics.GetTXHashesByAddress(evs, "0x580aD6Df3AC48d5223386DbbD4042818e66606D3")

	// Get insights from the tx
	for _, hash := range txHashArr {
		txInfo := metrics.GetTransaction(client, hash)
		fmt.Println("INSIGHTS: ", txInfo)

		// Insert each tx into DB
		txs.InsertTX(txInfo)
	}

	// smartcontractAddr := "0x00000"
	// count, err := eventStorage.GetEventCountByAddress(smartcontractAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// events, err := eventStorage.ListEventsByAddress(smartcontractAddr, "desc", count, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // iterar: obtener todos los event_datas por evento
	// batch := 100
	// offset := 0
	// for _, e := range events {
	// 	edCount, err := eventStorage.GetEventDataCount(smartcontractAddr, e.Abi.Name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// events, err := eventStorage.ListAllEvents()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, e := range events {
	// 	fmt.Printf("event=%+v \n", e)
	// }
	fmt.Println("envs", env.DatabaseDSN, env.MigrationDir, env.Port)
}

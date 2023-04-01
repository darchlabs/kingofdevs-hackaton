package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"

	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/env"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	eventstorage "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage/event"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/transaction"
)

func main() {
	// load env values
	var env env.Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	client, err := ethclient.Dial(env.ClientURL)
	if err != nil {
		panic(err)
	}

	s, err := storage.New(env.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	ev := eventstorage.New(s)

	contractAddr := "0xc13530546feA5fC787A2d126bB39bDeC20C4cc9e"
	txHash := transaction.GetTx(ev, client, contractAddr)
	transaction.GetSender(client, txHash)
	fmt.Println("2")
}

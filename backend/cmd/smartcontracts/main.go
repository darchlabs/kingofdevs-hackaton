package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/darchlabs/kingofdevs-hackaton/backend"
	metricsAPI "github.com/darchlabs/kingofdevs-hackaton/backend/internal/api/metrics"
	smartcontractsAPI "github.com/darchlabs/kingofdevs-hackaton/backend/internal/api/smartcontracts"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/env"
	metricsengine "github.com/darchlabs/kingofdevs-hackaton/backend/internal/metrics-engine"
	metricDB "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	smartcontractstorage "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage/smartcontract"
	transactionstorage "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage/transaction"
	"github.com/darchlabs/synchronizer-v2"
	eventDB "github.com/darchlabs/synchronizer-v2/pkg/storage"
	eventstorage "github.com/darchlabs/synchronizer-v2/pkg/storage/event"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/pressly/goose/v3"

	_ "github.com/darchlabs/kingofdevs-hackaton/backend/migrations"
)

var (
	eventStorage        synchronizer.EventStorage
	smartContactStorage backend.SmartContractStorage
)

func main() {
	// load env values
	var env env.Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	// initialize event db
	edb, err := eventDB.New(env.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	// initialize metric db
	db, err := metricDB.New(env.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	// run backend migrations
	err = goose.Up(db.DB.DB, env.MigrationDir)
	if err != nil {
		log.Fatal(err)
	}

	// initialize storages
	eventStorage = eventstorage.New(edb)
	transactionStorage := transactionstorage.New(db)
	smartContactStorage = smartcontractstorage.New(db)

	// Instance client
	client, err := ethclient.Dial(env.ClientURL)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	// initialize metrics
	m := metricsengine.New(client, eventStorage, transactionStorage, uuid.NewString, time.Now)

	// initialize fiber
	api := fiber.New()
	api.Use(logger.New())
	api.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// initialize routes
	smartcontractsAPI.Route(api, smartcontractsAPI.Context{
		Storage:      smartContactStorage,
		EventStorage: eventStorage,
		IDGen:        uuid.NewString,
		DateGen:      time.Now,
	})
	metricsAPI.Route(api, metricsAPI.Context{
		Metrics: m,
	})

	// Get all of the address from the events table
	events, err := eventStorage.ListAllEvents()

	// Get all of the hashes from the events table
	var allEventsAddresses []string
	for _, ev := range events {
		allEventsAddresses = append(allEventsAddresses, ev.Address)
	}

	// Get all of the events hashes
	txHashArr, err := m.GetTXHashesByAddresses(allEventsAddresses)
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}
	// Update them on a map for comparing with the transactions table hashes
	allEventsHashes := make(map[string]bool)
	for _, hash := range txHashArr {
		allEventsHashes[hash] = true
	}

	// Get the current hashes from the transaction table
	currentHashes, err := transactionStorage.ListCurrentHashes()
	if err != nil {
		log.Fatal("invalid env values, error: ", err)
	}

	// Compare if we already have this hashes
	for _, currentHash := range *currentHashes {
		// Delete the hashes that we already have in the db
		if allEventsHashes[currentHash] {
			delete(allEventsHashes, currentHash)
		}
	}

	// check if there are hashes that we don't have yet
	if len(allEventsHashes) > 0 {
		// Get insights from the tx that we don't have yet
		for missingHashes, _ := range allEventsHashes {
			// Get and update the missing on the transaction table
			txInfo, err := m.GetTransaction(missingHashes)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("INSIGHTS: ", txInfo)
		}

	}

	go func() {
		api.Listen(fmt.Sprintf(":%s", env.Port))
	}()

	// listen interrupt
	quit := make(chan struct{})
	listenInterrupt(quit)
	<-quit
	gracefullShutdown()
}

// listenInterrupt method used to listen SIGTERM OS Signal
func listenInterrupt(quit chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Println("Signal received", s.String())
		quit <- struct{}{}
	}()
}

// gracefullShutdown method used to close all synchronizer processes
func gracefullShutdown() {
	log.Println("Gracefully shutdown")

	// close database connection
	err := eventStorage.Stop()
	if err != nil {
		log.Println(err)
	}
}

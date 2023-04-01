package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/darchlabs/kingofdevs-hackaton/backend"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/api/metrics"
	smartcontractsAPI "github.com/darchlabs/kingofdevs-hackaton/backend/internal/api/smartcontracts"
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/env"
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

	// Get tx hash
	txHashArr := metrics.GetTXHashesByAddress(eventStorage, "0x580aD6Df3AC48d5223386DbbD4042818e66606D3")

	// Get insights from the tx
	for _, hash := range txHashArr {
		txInfo := metrics.GetTransaction(client, hash)
		fmt.Println("INSIGHTS: ", txInfo)

		// Insert each tx into DB
		transactionStorage.InsertTX(txInfo)
	}

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

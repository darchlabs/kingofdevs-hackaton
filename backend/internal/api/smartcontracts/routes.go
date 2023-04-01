package smartcontracts

import (
	"time"

	"github.com/darchlabs/kingofdevs-hackaton/backend"
	"github.com/gofiber/fiber/v2"
)

type idGenerator func() string
type dateGenerator func() time.Time

type Context struct {
	Storage backend.SmartContractStorage

	IDGen   idGenerator
	DateGen dateGenerator
}

func Route(app *fiber.App, ctx Context) {
	app.Post("/api/v1/smartcontracts/:address", insertSmartContractHandler(ctx))
	app.Get("/api/v1/smartcontracts", listSmartContracts(ctx))
}

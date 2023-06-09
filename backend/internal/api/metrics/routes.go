package metrics

import (
	"time"

	metricsengine "github.com/darchlabs/kingofdevs-hackaton/backend/internal/metrics-engine"
	smartcontractstorage "github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage/smartcontract"
	"github.com/gofiber/fiber/v2"
)

type idGenerator func() string
type dateGenerator func() time.Time

type Context struct {
	Metrics  *metricsengine.Metric
	Contract *smartcontractstorage.Storage
}

func Route(app *fiber.App, ctx Context) {
	app.Get("/api/v1/metrics/total-addresses/:id", totalAddressesHandler(ctx))
	app.Get("/api/v1/metrics/tvl/:id", TVLHandler(ctx))
	app.Get("/api/v1/metrics/total-transactions/:id", totalTransactionsHandler(ctx))
}

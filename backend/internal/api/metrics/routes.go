package metrics

import (
	"time"

	metricsengine "github.com/darchlabs/kingofdevs-hackaton/backend/internal/metrics-engine"
	"github.com/gofiber/fiber/v2"
)

type idGenerator func() string
type dateGenerator func() time.Time

type Context struct {
	Metrics *metricsengine.Metric
}

func Route(app *fiber.App, ctx Context) {
	app.Get("/api/v1/metrics/total-addresses", totalAddressesHandler(ctx))
}

package metrics

import (
	"github.com/darchlabs/synchronizer-v2/pkg/api"
	"github.com/gofiber/fiber/v2"
)

func TVLHandler(ctx Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusOK).JSON(api.Response{
				Error: "id cannot be nil",
			})
		}

		contract, err := ctx.Contract.GetSmartContractByID(id)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(api.Response{
				Error: err.Error(),
			})
		}

		tvl, err := ctx.Metrics.GetTVL(contract.Address)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(api.Response{
				Error: err.Error(),
			})
		}

		// prepare response
		return c.Status(fiber.StatusOK).JSON(api.Response{
			Data: tvl,
		})
	}
}

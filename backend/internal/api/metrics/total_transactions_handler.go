package metrics

import (
	"github.com/darchlabs/synchronizer-v2/pkg/api"
	"github.com/gofiber/fiber/v2"
)

func totalTransactionsHandler(ctx Context) func(c *fiber.Ctx) error {
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

		var addrArr []string
		addrArr = append(addrArr, contract.Address)

		txHashArr, err := ctx.Metrics.GetTXHashesByAddresses(addrArr)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(api.Response{
				Error: err.Error(),
			})
		}

		// Get total txs
		totalTXs := len(txHashArr)

		// prepare response
		return c.Status(fiber.StatusOK).JSON(api.Response{
			Data: totalTXs,
		})
	}
}

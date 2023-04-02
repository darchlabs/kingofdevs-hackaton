package metrics

import (
	"github.com/gofiber/fiber/v2"
)

func totalAddressesHandler(ctx Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		// example: use metric.GetTXHashesByAddress
		// txHashArr, er := ctx.Metrics.GetTXHashesByAddress("0x580aD6Df3AC48d5223386DbbD4042818e66606D3")

		// fmt.Println()

		// // Get insights from the tx
		// for _, hash := range txHashArr {
		// 	txInfo, _ := ctx.Metrics.GetTransaction(hash)
		// 	fmt.Println("INSIGHTS: ", txInfo)

		// }

		// prepare response
		return c.Status(fiber.StatusOK).JSON(struct{}{})
	}
}

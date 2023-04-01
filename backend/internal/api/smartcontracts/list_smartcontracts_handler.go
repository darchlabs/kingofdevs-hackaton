package smartcontracts

import (
	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/pagination"
	"github.com/gofiber/fiber/v2"
)

func listSmartContracts(ctx Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		// get pagination
		p := &pagination.Pagination{}
		err := p.GetPaginationFromFiber(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				},
			)
		}

		// get elements from database
		smartcontracts, err := ctx.Storage.ListSmartContracts(p.Sort, p.Limit, p.Offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				},
			)
		}

		// get all smartcontracts count from database
		count, err := ctx.Storage.GetSmartContractsCount()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				},
			)
		}

		// define meta response
		meta := make(map[string]interface{})
		meta["pagination"] = p.GetPaginationMeta(count)

		// prepare response
		return c.Status(fiber.StatusOK).JSON(struct {
			Data interface{} `json:"data"`
			Meta interface{} `json:"meta"`
		}{
			Data: smartcontracts,
			Meta: meta,
		})
	}
}

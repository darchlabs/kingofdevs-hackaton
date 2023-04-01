package smartcontracts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/smartcontract"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func insertSmartContractHandler(ctx Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		// prepate body request struct
		body := struct {
			SmartContract *smartcontract.SmartContract `json:"smartcontract"`
		}{}

		// parse body to smartcontract struct
		err := json.Unmarshal(c.Body(), &body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				},
			)
		}

		// validate body
		validate := validator.New()
		err = validate.Struct(body.SmartContract)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				},
			)
		}

		// Update smartcontract
		body.SmartContract.ID = ctx.IDGen()
		body.SmartContract.CreatedAt = ctx.DateGen()
		body.SmartContract.UpdatedAt = ctx.DateGen()

		// validate client works
		client, err := ethclient.Dial(body.SmartContract.NodeURL)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: fmt.Sprintf("can't valid ethclient error=%s", err),
				},
			)
		}

		// validate client is working correctly
		_, err = client.ChainID(context.Background())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: fmt.Sprintf("can't valid ethclient error=%s", err),
				},
			)
		}

		// save smartcontract struct on database
		createdSmartContract, err := ctx.Storage.InsertSmartContract(body.SmartContract)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				struct {
					Error string `json:"error"`
				}{
					Error: err.Error(),
				},
			)
		}

		// prepare response
		return c.Status(fiber.StatusOK).JSON(struct {
			Data *smartcontract.SmartContract `json:"data"`
		}{
			Data: createdSmartContract,
		})
	}
}

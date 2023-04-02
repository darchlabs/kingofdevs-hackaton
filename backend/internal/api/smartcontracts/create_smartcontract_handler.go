package smartcontracts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/smartcontract"
	"github.com/darchlabs/synchronizer-v2/pkg/event"
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

		// filter abi events from body
		events := make([]*event.Event, 0)
		for _, a := range body.SmartContract.Abi {
			if a.Type == "event" {
				// define new event
				ev := struct {
					Event *event.Event `json:"event"`
				}{
					Event: &event.Event{
						Network: body.SmartContract.Network,
						NodeURL: body.SmartContract.NodeURL,
						Address: body.SmartContract.Address,
						Abi:     a,
					},
				}

				b, err := json.Marshal(ev)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(
						struct {
							Error string `json:"error"`
						}{
							Error: err.Error(),
						},
					)
				}

				// send post to synchronizers
				url := fmt.Sprintf("%s/api/v1/events/%s", ctx.Env.SynchronizersApiURL, body.SmartContract.Address)
				res, err := http.Post(url, "application/json", bytes.NewBuffer(b))
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(
						struct {
							Error string `json:"error"`
						}{
							Error: err.Error(),
						},
					)
				}
				defer res.Body.Close()

				// read the response body
				b, err = ioutil.ReadAll(res.Body)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(
						struct {
							Error string `json:"error"`
						}{
							Error: err.Error(),
						},
					)
				}

				fmt.Println("BODY")
				fmt.Println("BODY")
				fmt.Println("BODY")
				fmt.Println("BODY")
				fmt.Println(string(b))

				// json.NewDecoder(res.Body).Decode()

				// events = append(events, ev)
			}
		}

		// Update smartcontract
		body.SmartContract.ID = ctx.IDGen()
		body.SmartContract.CreatedAt = ctx.DateGen()
		body.SmartContract.UpdatedAt = ctx.DateGen()
		body.SmartContract.Events = events
		for _, input := range body.SmartContract.Abi {
			input.ID = ctx.IDGen()
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

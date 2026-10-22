package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"strconv"
)

type payload struct {
	Url  string `json:"url"`
	Mode string `json:"mode"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/scrape", func(c fiber.Ctx) error {
		// payload var and target
		var Payload payload
		var target string

		// list of urls
		workerOne := 4000
		workerTwo := 5000
		workerThree := 6000

		// get the request
		if err := c.Bind().Body(&Payload); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// pass the worker based on the mode we recieve
		if Payload.Mode == "text" {
			Payload.Url = "http://localhost:" + strconv.Itoa(workerOne) + "/scrape"
			// send the payload and url to the worker1
		}

		if Payload.Mode == "json" {
			Payload.Url = "http://localhost:" + strconv.Itoa(workerTwo) + "/scrape"
			// send the payload and url to the worker2
		}

		// send the payload and url to the worker
		if Payload.Mode == "html" {
			Payload.Url = "http://localhost:" + strconv.Itoa(workerThree) + "/scrape"
			// send the payload and url to the worker3

		}

		// return the payload to the test
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"payload": Payload,
		})
	})

	app.Listen(":3000")
}

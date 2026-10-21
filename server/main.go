package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type payload struct {
	Url  string `json:"url"`
	Mode string `json:"mode"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/scrape", func(c fiber.Ctx) error {
		var Payload payload

		// get the request

		if err := c.Bind().Body(&Payload); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// return the payload to the test
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"payload": Payload,
		})
	})

	app.Listen(":3000")
}

package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run .cmd/server <port>")
		return
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid port")
		return
	}

	// init new server
	app := fiber.New()

	// get func test
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(
			fmt.Sprintf("Hello from backend :%d", port),
		)
	})

	// listen
	app.Listen(":" + strconv.Itoa(port))
}

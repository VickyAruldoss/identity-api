package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("app started in the port 3000")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("bull shit")
	})
	app.Listen(":3000")
}

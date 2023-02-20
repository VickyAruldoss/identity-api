package main

import (
	"fmt"

	"github.com/VickyAruldoss/identity-api/configuration"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config, err := configuration.NewConfigLoader().Load("config.json", "")
	if err != nil {
		fmt.Println("error while loading config", err)
	}
	fmt.Println("app started in the port 3000")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("bull shit")
	})
	addr := ":" + config.Port
	app.Listen(addr)
}

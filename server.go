package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("App is running")
		return err
	})

	app.Listen(":3000")
}

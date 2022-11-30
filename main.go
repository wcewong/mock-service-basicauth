package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	// Environmental vars
	port := os.Getenv("PORT")
	username := os.Getenv("USER")
	password := os.Getenv("PASS")

	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, 👋! If you are seeing this, it works! 🎉🎉")
	})

	app.Listen(":" + port)
}

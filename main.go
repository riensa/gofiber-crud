package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/crud/database"
)

func main() {
	// Start a new fiber app
  app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

  // Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!!!!")
		return err
})

	app.Listen(":3000")
}

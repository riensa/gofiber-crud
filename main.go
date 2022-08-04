package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/crud/database"
	"gofiber/crud/router"
)

func main() {
	// Start a new fiber app
  app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
  router.SetupRoutes(app)

	app.Listen(":3000")
}

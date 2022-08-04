package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gofiber/crud/database"
	"gofiber/crud/router"
	_ "gofiber/crud/docs"
	"github.com/gofiber/swagger"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

func main() {
	// Start a new fiber app
  app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
  router.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) 

	app.Listen(":3000")
}

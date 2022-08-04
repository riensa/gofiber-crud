package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	employeeRoutes "gofiber/crud/internal/routes/employee"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	employeeRoutes.SetupEmployeeRoutes(api)
}

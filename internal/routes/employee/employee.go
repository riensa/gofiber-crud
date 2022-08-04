package employeeRoutes

import (
	"github.com/gofiber/fiber/v2"
	employeeHandler "gofiber/crud/internal/handlers/employee"
)


func SetupEmployeeRoutes(router fiber.Router) {
	employee := router.Group("/employee")

	// new employee
	employee.Post("/", employeeHandler.CreateEmployees)

	// get list employee 
	employee.Get("/", employeeHandler.ReadEmployees)

	// // update one employee 
	employee.Put("/:id", employeeHandler.UpdateEmployee)

	// delete one employee 
	employee.Delete("/:id", employeeHandler.DeleteEmployee)

}
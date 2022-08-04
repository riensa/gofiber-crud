package employeeHandler

import (
	"fmt"
	"strconv"
	
	"github.com/gofiber/fiber/v2"
	"gofiber/crud/database"
	"gofiber/crud/internal/model"
)

func CreateEmployees(c *fiber.Ctx) error {
	db := database.DB
  employee := new(model.Employee)

	// Store the body in the employee and return error if encountered
	err := c.BodyParser(employee)
	if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the employee and return error if encountered
	err = db.Create(&employee).Error
	if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create employee", "data": err})
	}

	// Return the created employee
  return c.JSON(fiber.Map{"status": "success", "message": "Created Employee", "data": employee})
}

func ReadEmployees(c *fiber.Ctx) error {
	db := database.DB
	var employees []model.Employee
	var offset, limit int = 0, 10

	if(c.Query("page") != "") {
		page,err := strconv.Atoi(c.Query("page"))
		fmt.Println(err)
		
		offset = (page - 1) * limit
	}

	if(c.Query("limit") != "") {
		lmt,err := strconv.Atoi(c.Query("limit"))
		fmt.Println(err)

		limit = lmt
	}

	fmt.Println("offset", offset)
	fmt.Println("limit", limit)

	// find all employees in the database
	db.Offset(offset).Limit(limit).Find(&employees)

	// If no employee is present return an error
	if len(employees) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No employee found", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Employee Found", "data": employees})
}

func UpdateEmployee(c *fiber.Ctx) error {
	type updateEmployee struct {
		Name    string `json:"name"`
		Gender string `json:"gender"`
		City     string `json:"city	"`
	}

	db := database.DB
  var employee model.Employee

	// Read the param id
	id := c.Params("id")

	// Find employee with the given Id
	db.Find(&employee, "id = ?", id)

	fmt.Println("employee.ID", employee.ID)

	if employee.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No employee found", "data": nil})
	}

	var updateEmployeeData updateEmployee
	err := c.BodyParser(&updateEmployeeData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit Employee
	employee.Name = updateEmployeeData.Name
	employee.Gender = updateEmployeeData.Gender
	employee.City = updateEmployeeData.City

	// Save the Changes
	db.Save(&employee)
	return c.JSON(fiber.Map{"status": "success", "message": "Employee Found", "data": employee})
}

func DeleteEmployee(c *fiber.Ctx) error {
	db := database.DB
  var employee model.Employee

	// Read the param id
	id := c.Params("id")

	// Find employee with the given Id
	db.Find(&employee, "id = ?", id)

	fmt.Println("employee.ID", employee.ID)

	if employee.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No employee found", "data": nil})
	}

	// Delete employee and return error if encountered
	err := db.Delete(&employee, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete employee", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Employee Deleted"})
}

package handler

import (
	"hello-run/database"
	"hello-run/model"
	"hello-run/service"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func EmployeeCreate(c *fiber.Ctx) error {
	db := database.DB
	employeeClient := model.EmployeeClient{DB: db}
	var params service.EmployeeParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}
	
	employee, err := service.CreateEmployee(employeeClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to create employee",
		})
	}
	fmt.Println(err)
	return c.Status(fiber.StatusCreated).JSON(employee)
}

func EmployeeLogin(c *fiber.Ctx) error {
	db := database.DB
	employeeClient := model.EmployeeClient{DB: db}
	var params service.EmployeeParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	employee, err := service.LoginEmployee(employeeClient, params.Email, params.Password)
	fmt.Println(employee)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to login employee",
		})
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func EmployeeDelete(c *fiber.Ctx) error {
	db := database.DB
	employeeClient := model.EmployeeClient{DB: db}
	var params service.EmployeeParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	employee, err := service.DeleteEmployee(employeeClient, params.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to delete employee",
		})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

func GetAllReservationsWithoutEmployee(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	month := c.Params("month")

	reservations, err := service.AllReservationsWithoutEmployeeGet(reservationClient,month)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to get all reservations without employee",
		})
	}

	return c.Status(fiber.StatusOK).JSON(reservations)
}

func EmployeeConfirmReservation(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	employeeClient := model.EmployeeClient{DB: db}
	var params service.EmployeeReservationParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	err := service.EmployeeConfirmReservation(reservationClient, employeeClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to confirm reservation",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetAllReservationsByEmployee(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	employeeClient := model.EmployeeClient{DB: db}
	email := c.Params("email")

	reservations, err := service.GetAllReservationsByEmployee(reservationClient,employeeClient, email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to get all reservations by employee",
		})
	}

	return c.Status(fiber.StatusOK).JSON(reservations)
}

func GetAllReservationsDoneByEmployee(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	employeeClient := model.EmployeeClient{DB: db}
	email := c.Params("email")
	year := c.Params("year")
	month := c.Params("month")

	reservations, err := service.GetAllReservationsDoneByEmployee(reservationClient,employeeClient, email, year, month)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to get all reservations done by employee",
		})
	}

	return c.Status(fiber.StatusOK).JSON(reservations)
}
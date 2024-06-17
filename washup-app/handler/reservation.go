package handler

import (
	"fmt"
	"strings"

	"hello-run/database"
	"hello-run/model"
	"hello-run/service"
	"github.com/gofiber/fiber/v2"
)

func ObtenerFechasDisponiblesHandler(c *fiber.Ctx) error {
	fmt.Println("URL completa:", c.OriginalURL())
	servicioParam := c.Params("service")
	servicio := strings.Replace(servicioParam, "%20", " ", -1)

	fmt.Println("servicio:", servicio)

	db := database.DB
	reservationClient := model.ReservationClient{DB: db}

	fechasNoDisponibles, err := service.ObtenerFechasDisponibles(reservationClient, servicio)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al obtener fechas no disponibles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fechas_no_disponibles": fechasNoDisponibles,
	})
}

func ObtenerHorariosDisponiblesHandler(c *fiber.Ctx) error {
	servicioParam := c.Params("service")
	fechaParam := c.Params("date")
	servicio := strings.Replace(servicioParam, "%20", " ", -1)
	
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}

	horariosDisponibles	, err := service.ObtenerHorariosDisponibles(reservationClient, servicio,fechaParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al obtener fechas no disponibles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"horarios": horariosDisponibles,
	})
}

func ObtenerMisReservas(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	userClient := model.UserClient{DB: db}
	userEmailParam := c.Params("email")

	misReservas, err := service.ObtenerMisReservas(reservationClient, userClient, userEmailParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al obtener las reservas del usuario",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reservations": misReservas,
	})
}




func ReservaCreate(c *fiber.Ctx) error {
	db := database.DB
	userClient := model.UserClient{DB: db}
	reservationClient := model.ReservationClient{DB: db}
	var params service.ReservationParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al analizar JSON",
		})
	}

	reserva, err := service.CreateReservation(reservationClient, userClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al crear reserva",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(reserva)
}

func ReservaCheck(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	var params service.ReservationCheckParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al analizar JSON",
		})
	}

	disponible, err := service.CheckReservation(reservationClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al verificar disponibilidad",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"disponible": disponible,
	})
}

func ReservationDelete(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	reservationID := c.Params("reservationID")
	fmt.Println(reservationID)

	err := service.DeleteReservation(reservationClient, reservationID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error al verificar disponibilidad",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
                                                 
func EmployeeReservationDone(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	var params service.EmployeeReservationDoneParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	err := service.EmployeeReservationDone(reservationClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to confirm reservation",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetTotalProfitByMonth(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	month := c.Params("month")
	year := c.Params("year")

	profit, err := service.GetTotalProfitByMonth(reservationClient, month, year)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to get profit",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"profit": profit,
	})
}

func GetTotalProfitByYear(c *fiber.Ctx) error {
	db := database.DB
	reservationClient := model.ReservationClient{DB: db}
	year := c.Params("year")

	profit, err := service.GetTotalProfitByYear(reservationClient, year)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to get profit",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"profit": profit,
	})
}
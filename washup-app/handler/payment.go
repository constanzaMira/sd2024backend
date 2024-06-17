package handler

import (
	"fmt"

	"hello-run/service"
	"github.com/gofiber/fiber/v2"
)

func PaymentMercadoPago(c *fiber.Ctx) error {
	var paymentParams service.PaymentParams
	if err := c.BodyParser(&paymentParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to parse payment params",
		})
	}

	preferenceHandler := service.NewPreferenceHandler()
	preference, err := preferenceHandler.CreatePreference(paymentParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating preference",
		})
	}
	fmt.Println(preference)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"init_point": preference,
	})
}

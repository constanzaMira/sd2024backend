package middleware

import (
	"hello-run/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extraer el token de los encabezados de la solicitud.
		// si no funciona, probar c.GetReqHeaders()["Authorization"]
		tokenHeader := c.Get("Authorization")
		if tokenHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		// Normalmente, el encabezado de autorización tiene un formato: "Bearer <token>"
		// Por lo que se divide por espacio y se toma el segundo elemento.
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid/Malformed auth header",
			})
		}
		tokenPart := splitted[1] // El segundo elemento es el token.

		token, err := service.ValidateToken(tokenPart)
		if err != nil || token == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token is expired or not valid",
			})
		}

		// Si todo está bien, permitir que la solicitud continúe.
		return c.Next()
	}
}
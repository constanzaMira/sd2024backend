package handler

import (
	"hello-run/database"
	"hello-run/model"
	"hello-run/service"
	"github.com/gofiber/fiber/v2"
)

func UserCreate(c *fiber.Ctx) error {
	db := database.DB
	userClient := model.UserClient{DB: db}
	var params service.UserParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	user, err := service.CreateUser(userClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func UserLogin(c *fiber.Ctx) error {
	db := database.DB
	userClient := model.UserClient{DB: db}
	var params service.LoginParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	login, err := service.LoginUser(userClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to login user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(login)
}

func UserSocialLogin(c *fiber.Ctx) error {
	db := database.DB
	userClient := model.UserClient{DB: db}
	var params service.GoogleLoginParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON",
		})
	}

	login, err := service.GoogleLoginUser(userClient, params)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error trying to login user with Google",
		})
	}

	return c.Status(fiber.StatusOK).JSON(login)
}
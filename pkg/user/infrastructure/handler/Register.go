package handler

import (
	"TravelAPI/pkg/user/application"
	"TravelAPI/pkg/user/domain"

	"github.com/gofiber/fiber/v2"
)

func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody domain.User

		// Convert body to user
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to user",
			})
		}

		// Save in database
		user, err := application.Register(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"user": user,
		})
	}
}

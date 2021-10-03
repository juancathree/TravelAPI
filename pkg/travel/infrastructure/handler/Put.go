package handler

import (
	"TravelAPI/pkg/travel/application"
	"TravelAPI/pkg/travel/domain"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Put() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody domain.Travel

		fmt.Println(string(c.Body()))

		// Convert body to travel
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to travel",
			})
		}

		// Initialize and save in database
		travel, err := application.Put(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"travel": travel,
		})
	}
}

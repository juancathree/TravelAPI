package handler

import (
	"TravelAPI/pkg/travel/application"

	"github.com/gofiber/fiber/v2"
)

func Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if travelID has been passed
		if len(c.Params("id")) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": "travelID is required",
			})
		}

		// Get travelID
		id := c.Params("id")

		// Making database petition
		travel, err := application.Get(&id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"travel": travel,
		})
	}
}

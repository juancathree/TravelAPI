package handler

import (
	"TravelAPI/pkg/place/application"

	"github.com/gofiber/fiber/v2"
)

func GetPlaces() fiber.Handler {
	return func(c *fiber.Ctx) error {
		city := c.Params("id")

		// Making database petition
		places, err := application.Gets(&city)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"places": places,
		})
	}
}

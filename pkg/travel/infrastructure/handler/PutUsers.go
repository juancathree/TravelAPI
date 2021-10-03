package handler

import (
	"TravelAPI/pkg/travel/application"

	"github.com/gofiber/fiber/v2"
)

func PutUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request map[string]interface{}

		// Convert body to travel
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to travel",
			})
		}

		// Update document in database
		err = application.PutUsers(request["email"].(string), request["travelID"].(string), request["exists"].(bool))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "ok",
		})
	}
}

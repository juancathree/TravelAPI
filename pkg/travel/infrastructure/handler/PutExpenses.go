package handler

import (
	expenseDomain "TravelAPI/pkg/expenses/domain"
	"TravelAPI/pkg/travel/application"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func PutExpenses() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request map[string]interface{}

		// Convert body to travel
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to travel",
			})
		}

		amount := request["expense"].(map[string]interface{})["amount"]
		isPersonal := request["expense"].(map[string]interface{})["isPersonal"]

		amountFloat, err := strconv.ParseFloat(amount.(string), 64)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to travel",
			})
		}

		isPersonalBool, err := strconv.ParseBool(isPersonal.(string))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to travel",
			})
		}

		expense := &expenseDomain.Expense{
			Amount:      amountFloat,
			Description: request["expense"].(map[string]interface{})["description"].(string),
			IsPersonal:  isPersonalBool,
		}

		// Update document in database
		err = application.PutExpenses(strings.Split(c.Locals("email").(string), "@")[0], request["travelID"].(string), expense, request["exists"].(bool))
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

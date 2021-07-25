package route

import (
	"TravelAPI/pkg/travel/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/:id", handler.Get())
}

package router

import (
	"TravelAPI/pkg/place/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/cities", handler.GetCities())
	router.Get("/:id", handler.GetPlaces())
}

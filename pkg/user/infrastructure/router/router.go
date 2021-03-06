package router

import (
	"TravelAPI/pkg/user/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/login", handler.Login())
	router.Post("/register", handler.Register())
}

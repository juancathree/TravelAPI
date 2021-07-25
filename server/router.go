package server

import (
	travelRouter "TravelAPI/pkg/travel/infrastructure/router"
	"TravelAPI/pkg/user/infrastructure/middleware"
)

func SetupRouter() {
	// Travel
	travel := app.Group("/travel", middleware.AuthRequired())
	travelRouter.Router(travel)
}

package server

import (
	placeRouter "TravelAPI/pkg/place/infrastructure/router"
	travelRouter "TravelAPI/pkg/travel/infrastructure/router"
	"TravelAPI/pkg/user/infrastructure/middleware"
	authRouter "TravelAPI/pkg/user/infrastructure/router"
)

func SetupRouter() {
	// Auth
	auth := app.Group("/")
	authRouter.Router(auth)

	// Travel
	travel := app.Group("/travel", middleware.AuthRequired())
	travelRouter.Router(travel)

	// Place
	place := app.Group("/place", middleware.AuthRequired())
	placeRouter.Router(place)
}

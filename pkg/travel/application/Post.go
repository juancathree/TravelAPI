package application

import (
	itineraryApplication "TravelAPI/pkg/itinerary/application"
	"TravelAPI/pkg/travel/domain"
)

func Post(travel *domain.Travel, userID string) (*domain.Travel, error) {
	// Initialize travel
	travel.Initialize(&userID)

	// Call use case create routes
	err := itineraryApplication.CreateJourney(&travel.Itinerary)
	if err != nil {
		return nil, err
	}

	return repo.Post(travel)
}

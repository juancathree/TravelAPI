package application

import (
	itineraryApplication "TravelAPI/pkg/itinerary/application"
	"TravelAPI/pkg/travel/domain"
	userApplication "TravelAPI/pkg/user/application"
)

func Post(travel *domain.Travel, userID string) (*domain.Travel, error) {

	// Initialize travel
	travel.Initialize(&userID)

	// Call use case create routes
	err := itineraryApplication.CreateJourney(&travel.Itinerary)
	if err != nil {
		return nil, err
	}

	newTravel, err := repo.Post(travel)
	if err != nil {
		return nil, err
	}

	err = userApplication.PostTravel(&userID, &newTravel.ID)
	if err != nil {
		return nil, err
	}

	return newTravel, nil
}

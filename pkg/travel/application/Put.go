package application

import (
	itineraryApplication "TravelAPI/pkg/itinerary/application"
	"TravelAPI/pkg/travel/domain"
)

func Put(travel *domain.Travel) (*domain.Travel, error) {

	// initialize itinerary
	travel.Itinerary.Initialize()

	// Call use case create routes
	err := itineraryApplication.CreateJourney(&travel.Itinerary)
	if err != nil {
		return nil, err
	}

	modifiedTravel, err := repo.Put(travel)
	if err != nil {
		return nil, err
	}

	return modifiedTravel, nil
}

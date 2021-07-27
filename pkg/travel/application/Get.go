package application

import "TravelAPI/pkg/travel/domain"

func Get(travelID *string) (*domain.Travel, error) {
	return repo.Get(travelID)
}

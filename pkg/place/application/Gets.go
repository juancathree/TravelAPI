package application

import "TravelAPI/pkg/place/domain"

func Gets(city *string) (*[]domain.Place, error) {
	return repo.Gets(city)
}

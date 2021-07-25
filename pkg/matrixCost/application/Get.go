package application

import "TravelAPI/pkg/matrixCost/domain"

func Get(city *string) (*domain.MatrixCost, error) {
	return repo.Get(city)
}

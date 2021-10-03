package application

import "TravelAPI/pkg/travel/domain"

func Gets(userID *string) (*[]domain.Travel, error) {
	return repo.Gets(userID)
}

package application

import "TravelAPI/pkg/user/domain"

func Get(userID *string) (*domain.User, error) {
	return repo.Get(userID)
}

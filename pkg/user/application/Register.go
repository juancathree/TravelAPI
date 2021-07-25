package application

import "TravelAPI/pkg/user/domain"

func Register(user *domain.User) (*domain.User, error) {
	// Set password
	user.SetPassword(&user.Password)

	return repo.Post(user)
}

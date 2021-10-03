package application

import (
	"TravelAPI/pkg/user/domain"
	"TravelAPI/shared"
	"fmt"
)

func Register(user *domain.User) (*domain.User, error) {

	err := shared.Validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Set password
	user.SetPassword(&user.Password)

	return repo.Post(user)
}

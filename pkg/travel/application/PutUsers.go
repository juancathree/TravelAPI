package application

import (
	userAPP "TravelAPI/pkg/user/application"
)

func PutUsers(userID, travelID string, exist bool) error {

	_, err := userAPP.Get(&userID)

	if err != nil {
		return err
	}

	err = repo.PutTravelUsers(&userID, &travelID, &exist)
	if err != nil {
		return err
	}

	return nil
}

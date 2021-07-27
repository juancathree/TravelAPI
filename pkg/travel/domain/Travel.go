package domain

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"
	itineraryDomain "TravelAPI/pkg/itinerary/domain"
)

type Travel struct {
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
	itineraryDomain.Itinerary
	expensesDomain.Expenses
	TravelUsers []string `validate:"required,dive,required,unique" json:"travelUsers" bson:"travelUsers"`
}

func (t *Travel) Initialize(userID *string) {
	// Initialize TravelUsers
	t.TravelUsers = append(t.TravelUsers, *userID)

	// Initialize Expenses
	t.Expenses.Initialize(userID)

	// Initialize Itinerary
	t.Itinerary.Initialize()
}

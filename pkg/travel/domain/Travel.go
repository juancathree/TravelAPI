package domain

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"
	itineraryDomain "TravelAPI/pkg/itinerary/domain"
)

type Travel struct {
	ID string `json:"_id" bson:"_id"`
	itineraryDomain.Itinerary
	expensesDomain.Expenses
	TravelUsers []string `validate:"required,dive,required,unique" json:"travelUsers" bson:"travelUsers"`
}

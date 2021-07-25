package domain

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"
	scheduleDomain "TravelAPI/pkg/schedule/domain"
)

type TravelRepository interface {
	GetTravel(travelID *string) (*Travel, error)
	GetTravels(userID string) (*[]Travel, error)
	PostTravel(travel *Travel) (*Travel, error)
	PutTravelSchedule(schedule *scheduleDomain.Schedule, travelID *string) error
	PutTravelCustomEntryLocations(customEntryLocations *[]string, travelID *string) error
	PutTravelExpenses(userID, travelID *string, expense *expensesDomain.Expense, exists *bool) error
	PutTravelUsersID(userID, travelID *string, exists *bool) error
	DeleteTravel(travelID *string) error
}

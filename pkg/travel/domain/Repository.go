package domain

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"
	scheduleDomain "TravelAPI/pkg/schedule/domain"
)

type Repository interface {
	Get(travelID *string) (*Travel, error)
	Gets(userID *string) (*[]Travel, error)
	Post(travel *Travel) (*Travel, error)
	PutSchedule(schedule *scheduleDomain.Schedule, travelID *string) error
	PutCustomEntryLocations(customEntryLocations *[]string, travelID *string) error
	PutExpenses(userID, travelID *string, expense *expensesDomain.Expense, exists *bool) error
	PutTravelUsers(userID, travelID *string, exists *bool) error
	Delete(travelID *string) error
}

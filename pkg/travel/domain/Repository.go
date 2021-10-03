package domain

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"
)

type Repository interface {
	Get(travelID *string) (*Travel, error)
	Gets(userID *string) (*[]Travel, error)
	Post(travel *Travel) (*Travel, error)
	Put(travel *Travel) (*Travel, error)
	PutExpenses(userID, travelID *string, expense *expensesDomain.Expense, exists *bool) error
	PutTravelUsers(userID, travelID *string, exists *bool) error
	Delete(travelID *string) error
}

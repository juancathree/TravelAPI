package application

import (
	expenseDomain "TravelAPI/pkg/expenses/domain"
)

func PutExpenses(userID, travelID string, spend *expenseDomain.Expense, exist bool) error {

	err := repo.PutExpenses(&userID, &travelID, spend, &exist)
	if err != nil {
		return err
	}

	return nil
}

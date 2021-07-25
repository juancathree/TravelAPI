package persistence

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PutExpenses(userID, travelID *string, expense *expensesDomain.Expense, exists *bool) error {
	// Create the update
	var update bson.M

	if *exists {
		update = bson.M{"$pull": bson.M{"expenses." + (*userID) + ".expense": *expense}}
	} else {
		update = bson.M{"$push": bson.M{"expenses." + (*userID) + ".expense": *expense}}
	}

	return mo.UpdateDocument(travelID, &update)
}

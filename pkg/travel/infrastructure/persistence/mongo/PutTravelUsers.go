package persistence

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PutTravelUsers(userID, travelID *string, exists *bool) error {
	// Create the update
	var update bson.M

	if *exists {
		update = bson.M{"$pull": bson.M{"travelUsers": *userID}, "$unset": bson.M{"expenses." + strings.Split(*userID, "@")[0]: ""}}
	} else {
		userExpense := expensesDomain.UserExpense{
			Amount:  0.0,
			Expense: []expensesDomain.Expense{},
		}
		update = bson.M{"$push": bson.M{"travelUsers": *userID}, "$set": bson.M{"expenses." + strings.Split(*userID, "@")[0]: userExpense}}
	}

	return mo.UpdateDocument(travelID, &update)
}

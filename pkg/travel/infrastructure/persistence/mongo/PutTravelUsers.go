package persistence

import (
	expensesDomain "TravelAPI/pkg/expenses/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) PutTravelUsers(userID, travelID *string, exists *bool) error {
	// Create the update
	var update bson.M

	if *exists {
		update = bson.M{"$pull": bson.M{"usersID": *userID}, "$unset": bson.M{"expenses." + (*userID): ""}}
	} else {
		userExpense := expensesDomain.UserExpense{
			Amount:  0.0,
			Expense: []expensesDomain.Expense{},
		}
		update = bson.M{"$push": bson.M{"usersID": *userID}, "$set": bson.M{"expenses." + (*userID): userExpense}}
	}

	return mo.UpdateDocument(travelID, &update)
}

package persistence

import (
	"TravelAPI/pkg/travel/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) Gets(userID *string) (*[]domain.Travel, error) {
	ctx := context.Background()

	// Create the filter
	filter := bson.M{"travelUsers": *userID}

	// Get from database and decode
	travels := make([]domain.Travel, 0, 5)
	cursor, err := mo.Collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("database error: couldn't get object with ID: %v from database", userID)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var travel domain.Travel
		err := cursor.Decode(&travel)
		if err != nil {
			return nil, err
		}
		travels = append(travels, travel)
	}

	return &travels, nil
}

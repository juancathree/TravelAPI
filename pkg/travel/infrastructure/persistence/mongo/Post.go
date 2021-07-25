package persistence

import (
	"TravelAPI/pkg/travel/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) Post(travel *domain.Travel) (*domain.Travel, error) {
	// Inserting travel to database
	result, err := mo.Collection.InsertOne(context.Background(), *travel)
	if err != nil {
		return nil, fmt.Errorf("database error: couldn't be posted travel to database")
	}

	// Convert result to ObjectID
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("conversion error: couldn't be converted %v to ObjectID", (*result).InsertedID)
	}

	// Convert ObjectID to string
	travel.ID = oid.Hex()

	return travel, nil
}

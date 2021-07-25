package persistence

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilterID(travelID *string) (*primitive.M, error) {
	// Convert travelID to primitiveID
	id, err := primitive.ObjectIDFromHex(*travelID)
	if err != nil {
		return nil, fmt.Errorf("conversion error: couldn't be converted %v to ObjectID", *travelID)
	}

	// Create a filter for MongoDB query
	filter := bson.M{"_id": id}

	return &filter, nil
}

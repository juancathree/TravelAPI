package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilterID(userID *string) *primitive.M {
	// Create a filter for MongoDB query
	filter := bson.M{"_id": *userID}

	return &filter
}

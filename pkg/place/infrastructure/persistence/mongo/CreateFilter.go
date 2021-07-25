package persistence

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mo *MongoRepository) CreateFilter(city *string, pref, cuV *[]string) *primitive.M {
	// Create a filter for MongoDB query
	filter := bson.M{
		"city":     *city,
		"category": bson.M{"$nin": *pref},
		"name":     bson.M{"$nin": *cuV},
	}

	return &filter
}

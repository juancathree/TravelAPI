package persistence

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (mo *MongoRepository) GetCities() (*[]string, error) {

	// Get from database and decode
	result, err := mo.Collection.Distinct(context.Background(), "city", bson.M{})
	if err != nil {
		return nil, fmt.Errorf("database error: couldn't get cities from database")
	}

	resultString := make([]string, 0, len(result))

	for _, r := range result {
		resultString = append(resultString, r.(string))
	}

	return &resultString, nil
}

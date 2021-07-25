package persistence

import (
	"TravelAPI/pkg/matrixCost/domain"
	"context"
	"fmt"
)

func (mo *MongoRepository) Get(city *string) (*domain.MatrixCost, error) {
	// Create the filter
	filter := mo.CreateFilterID(city)

	// Get from database and decode
	var matrixCost domain.MatrixCost
	err := mo.Collection.FindOne(context.Background(), filter).Decode(&matrixCost)
	if err != nil {
		return nil, fmt.Errorf("database error: couldn't get object with ID: %v from database", *city)
	}

	return &matrixCost, nil
}

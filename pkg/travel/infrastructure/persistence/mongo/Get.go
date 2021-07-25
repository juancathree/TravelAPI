package persistence

import (
	"TravelAPI/pkg/travel/domain"
	"context"
	"fmt"
)

func (mo *MongoRepository) Get(travelID *string) (*domain.Travel, error) {
	// Create the filter
	filter, err := mo.CreateFilterID(travelID)
	if err != nil {
		return nil, err
	}
	// Get from database and decode
	var travel domain.Travel
	err = mo.Collection.FindOne(context.Background(), filter).Decode(&travel)
	if err != nil {
		return nil, fmt.Errorf("database error: couldn't get object with ID: %v from database", *travelID)
	}

	return &travel, nil
}

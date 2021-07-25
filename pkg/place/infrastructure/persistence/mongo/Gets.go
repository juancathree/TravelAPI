package persistence

import (
	"TravelAPI/pkg/place/domain"
	"context"
	"fmt"
)

func (mo *MongoRepository) Gets(city *string) (*[]domain.Place, error) {
	// Create the filter
	filter := mo.CreateFilterID(city)

	// Get from database and decode
	places := make([]domain.Place, 0, 20)
	cursor, err := mo.Collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("database error: couldn't get object with ID: %v from database", *city)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var place domain.Place
		err := cursor.Decode(&place)

		if err != nil {
			return nil, fmt.Errorf("couldn't be decode place")
		}

		places = append(places, place)
	}

	return &places, nil
}

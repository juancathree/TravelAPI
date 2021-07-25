package persistence

import (
	"context"
	"fmt"
)

func (mo *MongoRepository) Delete(travelID *string) error {
	// Create the filter
	filter, err := mo.CreateFilterID(travelID)
	if err != nil {
		return err
	}

	// Delete the document in MongoDB
	_, err = mo.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("database error: couldn't be deleted the travel with ID: %v", *travelID)
	}

	return nil
}

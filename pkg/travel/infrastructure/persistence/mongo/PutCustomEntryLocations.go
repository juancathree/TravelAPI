package persistence

import "go.mongodb.org/mongo-driver/bson"

func (mo *MongoRepository) PutCustomEntryLocations(cu *[]string, travelID *string) error {
	// Create the update
	update := bson.M{
		"$set": bson.M{
			"customEntryLocations": *cu,
		},
	}

	return mo.UpdateDocument(travelID, &update)
}

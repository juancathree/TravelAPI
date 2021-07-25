package application

import (
	"TravelAPI/database"
	"TravelAPI/pkg/place/domain"
	persistence "TravelAPI/pkg/place/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
	}
}

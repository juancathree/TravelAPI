package application

import (
	"TravelAPI/database"
	"TravelAPI/pkg/travel/domain"
	persistence "TravelAPI/pkg/travel/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
	}
}

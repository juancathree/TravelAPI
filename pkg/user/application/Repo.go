package application

import (
	"TravelAPI/database"
	"TravelAPI/pkg/user/domain"
	persistence "TravelAPI/pkg/user/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
	}
}

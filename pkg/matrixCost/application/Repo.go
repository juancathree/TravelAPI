package application

import (
	"TravelAPI/database"
	"TravelAPI/pkg/matrixCost/domain"
	persistence "TravelAPI/pkg/matrixCost/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo domain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		repo = persistence.NewMongoRepository()
	}
}

package application

import (
	"TravelAPI/database"
	matrixDomain "TravelAPI/pkg/matrixCost/domain"
	matrixPersistence "TravelAPI/pkg/matrixCost/infrastructure/persistence/mongo"
	placeDomain "TravelAPI/pkg/place/domain"
	placePersistence "TravelAPI/pkg/place/infrastructure/persistence/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var placeRepo placeDomain.Repository
var matrixRepo matrixDomain.Repository

func init() {
	switch database.DB.(type) {
	case *mongo.Database:
		placeRepo = placePersistence.NewMongoRepository()
		matrixRepo = matrixPersistence.NewMongoRepository()
	}
}

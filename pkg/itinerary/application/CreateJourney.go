package application

import (
	itineraryDomain "TravelAPI/pkg/itinerary/domain"
	journeyApplication "TravelAPI/pkg/journey/application"
	placeDomain "TravelAPI/pkg/place/domain"
	"sort"
)

func CreateJourney(itinerary *itineraryDomain.Itinerary) error {
	// Get places
	places, err := placeRepo.GetsCustom(&itinerary.City, &itinerary.Preferences, &itinerary.CustomVisitLocations)
	if err != nil {
		return err
	}

	// Get matrix costs
	matrixCost, err := matrixRepo.Get(&itinerary.City)
	if err != nil {
		return err
	}

	// Sort places by priority
	sort.Sort(placeDomain.ByPriority(*places))

	// Get places to visit
	visit := itinerary.GetVisitPlaces(places, &matrixCost.AverageCost)

	// getting clusters
	var cluster Cluster
	clusters := cluster.CreateCluster(visit, matrixCost)

	// getting journeys
	journey := journeyApplication.CreateJourney(clusters, places, matrixCost)

	itinerary.Journey = *CreateItineraryJourney(itinerary, journey, places, matrixCost)

	return nil
}

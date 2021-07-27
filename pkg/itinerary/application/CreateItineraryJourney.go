package application

import (
	"TravelAPI/pkg/itinerary/domain"
	journeyDomain "TravelAPI/pkg/journey/domain"
	matrixDomain "TravelAPI/pkg/matrixCost/domain"
	placeDomain "TravelAPI/pkg/place/domain"
)

func CreateItineraryJourney(itinerary *domain.Itinerary, journey *[]placeDomain.Place, places *[]placeDomain.Place, matrixCost *matrixDomain.MatrixCost) *journeyDomain.Journey {

	var j journeyDomain.Journey
	j.Initialize(&itinerary.TravelTime)

	k := 0

	for i := 0; i < itinerary.TravelTime; i++ {
		time := itinerary.Schedule.HoursOf(i)

		first := true

		for time > 0 && k < len(*journey) {

			currentPlace := (*journey)[k]

			if first {
				j[i] = append(j[i], currentPlace.Name)

				if itinerary.CustomEntryLocations[currentPlace.Name] {
					time -= currentPlace.Visit.All
				} else {
					time -= currentPlace.Visit.Outside
				}

				first = false
				k++

			} else {

				previousPlace := (*journey)[k-1]

				if itinerary.CustomEntryLocations[currentPlace.Name] {

					timeVisit := currentPlace.Visit.All + (*matrixCost).Costs[currentPlace.Name][previousPlace.Name]

					if time > timeVisit {
						time -= timeVisit
						j[i] = append(j[i], currentPlace.Name)
						k++
					} else {
						break
					}

				} else {

					timeVisit := currentPlace.Visit.Outside + (*matrixCost).Costs[currentPlace.Name][previousPlace.Name]

					if time > timeVisit {
						time -= timeVisit
						j[i] = append(j[i], currentPlace.Name)
						k++
					} else {
						break
					}
				}
			}
		}
	}

	return &j
}

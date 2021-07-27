package domain

import (
	placeDomain "TravelAPI/pkg/place/domain"
)

func (i *Itinerary) GetVisitPlaces(places *[]placeDomain.Place, averageCost *float64) *[]placeDomain.Place {
	// Get total hours
	hours := i.GetTotalHours()

	visit := make([]placeDomain.Place, 0, len(*places))

	for _, p := range *places {
		// check if it is in custom entry customEntryLocations
		if i.CustomEntryLocations[p.Name] {
			timeVisit := p.Visit.All + *averageCost
			if *hours-timeVisit > 0 {
				*hours -= timeVisit
				visit = append(visit, p)
			} else {
				break
			}
			// if isn't in customEntryLocations
		} else if p.Visit.Outside != 0 {
			timeVisit := p.Visit.Outside + *averageCost
			if *hours-timeVisit > 0 {
				*hours -= timeVisit
				visit = append(visit, p)
			} else {
				break
			}
		}
	}

	return &visit
}

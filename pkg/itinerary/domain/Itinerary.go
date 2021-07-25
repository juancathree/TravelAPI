package domain

import (
	journeyDomain "TravelAPI/pkg/journey/domain"
	scheduleDomain "TravelAPI/pkg/schedule/domain"
)

type Itinerary struct {
	City string `validate:"required,lowercase" json:"city" bson:"city"`
	scheduleDomain.Schedule
	CustomEntryLocations  map[string]bool `validate:"dive,unique" json:"customEntryLocations" bson:"customEntryLocations"`
	CustomVisitLocations  []string        `validate:"dive,unique" json:"customVisitLocations" bson:"customVisitLocations"`
	Preferences           []string        `validate:"dive,unique" json:"preferences" bson:"preferences"`
	journeyDomain.Journey `validate:"required" json:"routes" bson:"routes"`
}

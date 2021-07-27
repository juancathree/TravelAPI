package domain

import (
	"os"
	"strconv"
)

func (s *Schedule) HoursOf(day int) float64 {
	time := 0.0

	if day == 0 {
		time = float64(s.ItineraryEndTime.Hour() - s.StartDay.Hour())
	} else if day == s.TravelTime-1 {
		time = float64(s.EndDay.Hour() - s.ItineraryStartTime.Hour())
	} else {
		time = float64(s.ItineraryEndTime.Hour() - s.ItineraryStartTime.Hour())
	}

	kmh, err := strconv.ParseFloat(os.Getenv("TRAVELER_KMH"), 64)
	if err != nil {
		return 0.0
	}

	return time * kmh
}

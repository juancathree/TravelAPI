package domain

import (
	"os"
	"strconv"
)

// GetTotalHours returns the total hours
// available to visit places
func (s *Schedule) GetTotalHours() *float64 {
	// Get hours of first day
	hours := s.ItineraryEndTime.Hour() - s.StartDay.Hour()

	// Get hours of intermediate days
	hours += (s.ItineraryEndTime.Hour()-s.ItineraryStartTime.Hour())*s.TravelTime - 2

	// Get hours of last day
	hours += s.EndDay.Hour() - s.ItineraryStartTime.Hour()

	// Estimate how many kilometers can do
	kmh, err := strconv.Atoi(os.Getenv("TRAVELER_KMH"))
	if err != nil {
		return nil
	}

	hours *= kmh

	// Convert hours to float64
	totalTime := float64(hours)

	return &totalTime
}

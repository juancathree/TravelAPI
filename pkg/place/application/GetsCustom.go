package application

import "TravelAPI/pkg/place/domain"

func GetsCustom(city *string, pref, cuV *[]string) (*[]domain.Place, error) {
	return repo.GetsCustom(city, pref, cuV)
}

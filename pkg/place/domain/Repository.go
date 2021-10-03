package domain

type Repository interface {
	Gets(city *string) (*[]Place, error)
	GetsCustom(city *string, pref, cuV *[]string) (*[]Place, error)
	GetCities() (*[]string, error)
}

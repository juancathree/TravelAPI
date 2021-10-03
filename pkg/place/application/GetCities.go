package application

func GetCities() (*[]string, error) {
	return repo.GetCities()
}

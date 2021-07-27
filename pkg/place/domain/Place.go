package domain

type Place struct {
	Name        string   `validate:"required" json:"name" bson:"name"`
	Location    Location `validate:"required" json:"location" bson:"location"`
	Information string   `validate:"required,max=200" json:"information" bson:"information"`
	Visit       Visit    `validate:"required" json:"visit" bson:"visit"`
	Category    string   `validate:"required" json:"category" bson:"category"`
	Priority    int      `validate:"required" json:"priority" bson:"priority"`
}

type Location struct {
	Type        string    `validate:"required" json:"type" bson:"type"`
	Coordinates []float64 `validate:"required" json:"coordinates" bson:"coordinates"`
}
type Visit struct {
	All     float64 `json:"all,omitempty" bson:"all,omitempty"`
	Outside float64 `json:"outside,omitempty" bson:"outside,omitempty"`
}

func (p *Place) Latitude() float64 {
	return p.Location.Coordinates[0]
}

func (p *Place) Longitude() float64 {
	return p.Location.Coordinates[1]
}

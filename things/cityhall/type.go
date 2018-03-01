package cityhall

import "github.com/dpb587/go-schemaorg"

// // A city hall.
var Type = schemaorg.NewDataType("http://schema.org", "CityHall")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

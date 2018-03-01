package bustrip

import "github.com/dpb587/go-schemaorg"

// // A trip on a commercial bus line.
var Type = schemaorg.NewDataType("http://schema.org", "BusTrip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

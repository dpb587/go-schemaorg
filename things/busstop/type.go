package busstop

import "github.com/dpb587/go-schemaorg"

// // A bus stop.
var Type = schemaorg.NewDataType("http://schema.org", "BusStop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

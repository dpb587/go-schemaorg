package flight

import "github.com/dpb587/go-schemaorg"

// // An airline flight.
var Type = schemaorg.NewDataType("http://schema.org", "Flight")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

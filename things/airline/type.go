package airline

import "github.com/dpb587/go-schemaorg"

// // An organization that provides flights for passengers.
var Type = schemaorg.NewDataType("http://schema.org", "Airline")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

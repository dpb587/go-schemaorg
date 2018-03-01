package airport

import "github.com/dpb587/go-schemaorg"

// // An airport.
var Type = schemaorg.NewDataType("http://schema.org", "Airport")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

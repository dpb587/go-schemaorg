package dancegroup

import "github.com/dpb587/go-schemaorg"

// // A dance group&#x2014;for example, the Alvin Ailey Dance Theater or
// Riverdance.
var Type = schemaorg.NewDataType("http://schema.org", "DanceGroup")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

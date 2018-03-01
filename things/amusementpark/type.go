package amusementpark

import "github.com/dpb587/go-schemaorg"

// // An amusement park.
var Type = schemaorg.NewDataType("http://schema.org", "AmusementPark")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

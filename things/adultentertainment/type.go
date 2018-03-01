package adultentertainment

import "github.com/dpb587/go-schemaorg"

// // An adult entertainment establishment.
var Type = schemaorg.NewDataType("http://schema.org", "AdultEntertainment")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

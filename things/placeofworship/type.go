package placeofworship

import "github.com/dpb587/go-schemaorg"

// // Place of worship, such as a church, synagogue, or mosque.
var Type = schemaorg.NewDataType("http://schema.org", "PlaceOfWorship")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

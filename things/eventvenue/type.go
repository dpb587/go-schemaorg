package eventvenue

import "github.com/dpb587/go-schemaorg"

// // An event venue.
var Type = schemaorg.NewDataType("http://schema.org", "EventVenue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

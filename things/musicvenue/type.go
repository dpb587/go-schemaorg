package musicvenue

import "github.com/dpb587/go-schemaorg"

// // A music venue.
var Type = schemaorg.NewDataType("http://schema.org", "MusicVenue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

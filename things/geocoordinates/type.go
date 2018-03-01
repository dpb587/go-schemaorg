package geocoordinates

import "github.com/dpb587/go-schemaorg"

// // The geographic coordinates of a place or event.
var Type = schemaorg.NewDataType("http://schema.org", "GeoCoordinates")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

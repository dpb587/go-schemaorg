package geoshape

import "github.com/dpb587/go-schemaorg"

// // The geographic shape of a place. A GeoShape can be described using several
// properties whose values are based on latitude/longitude pairs. Either
// whitespace or commas can be used to separate latitude and longitude;
// whitespace should be used when writing a list of several such points.
var Type = schemaorg.NewDataType("http://schema.org", "GeoShape")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

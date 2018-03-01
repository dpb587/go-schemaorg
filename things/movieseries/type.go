package movieseries

import "github.com/dpb587/go-schemaorg"

// // A series of movies. Included movies can be indicated with the hasPart
// property.
var Type = schemaorg.NewDataType("http://schema.org", "MovieSeries")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package movie

import "github.com/dpb587/go-schemaorg"

// // A movie.
var Type = schemaorg.NewDataType("http://schema.org", "Movie")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

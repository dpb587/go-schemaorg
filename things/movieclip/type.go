package movieclip

import "github.com/dpb587/go-schemaorg"

// // A short segment/part of a movie.
var Type = schemaorg.NewDataType("http://schema.org", "MovieClip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package movietheater

import "github.com/dpb587/go-schemaorg"

// // A movie theater.
var Type = schemaorg.NewDataType("http://schema.org", "MovieTheater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

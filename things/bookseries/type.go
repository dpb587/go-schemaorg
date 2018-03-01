package bookseries

import "github.com/dpb587/go-schemaorg"

// // A series of books. Included books can be indicated with the hasPart property.
var Type = schemaorg.NewDataType("http://schema.org", "BookSeries")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

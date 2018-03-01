package book

import "github.com/dpb587/go-schemaorg"

// // A book.
var Type = schemaorg.NewDataType("http://schema.org", "Book")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

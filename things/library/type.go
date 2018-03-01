package library

import "github.com/dpb587/go-schemaorg"

// // A library.
var Type = schemaorg.NewDataType("http://schema.org", "Library")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

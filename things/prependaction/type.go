package prependaction

import "github.com/dpb587/go-schemaorg"

// // The act of inserting at the beginning if an ordered collection.
var Type = schemaorg.NewDataType("http://schema.org", "PrependAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

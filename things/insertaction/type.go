package insertaction

import "github.com/dpb587/go-schemaorg"

// // The act of adding at a specific location in an ordered collection.
var Type = schemaorg.NewDataType("http://schema.org", "InsertAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

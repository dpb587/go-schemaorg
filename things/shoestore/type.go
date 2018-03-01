package shoestore

import "github.com/dpb587/go-schemaorg"

// // A shoe store.
var Type = schemaorg.NewDataType("http://schema.org", "ShoeStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

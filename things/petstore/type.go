package petstore

import "github.com/dpb587/go-schemaorg"

// // A pet store.
var Type = schemaorg.NewDataType("http://schema.org", "PetStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

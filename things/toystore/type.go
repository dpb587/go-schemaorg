package toystore

import "github.com/dpb587/go-schemaorg"

// // A toy store.
var Type = schemaorg.NewDataType("http://schema.org", "ToyStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

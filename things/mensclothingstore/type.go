package mensclothingstore

import "github.com/dpb587/go-schemaorg"

// // A men's clothing store.
var Type = schemaorg.NewDataType("http://schema.org", "MensClothingStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

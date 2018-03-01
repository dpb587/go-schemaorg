package clothingstore

import "github.com/dpb587/go-schemaorg"

// // A clothing store.
var Type = schemaorg.NewDataType("http://schema.org", "ClothingStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

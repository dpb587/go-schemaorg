package furniturestore

import "github.com/dpb587/go-schemaorg"

// // A furniture store.
var Type = schemaorg.NewDataType("http://schema.org", "FurnitureStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

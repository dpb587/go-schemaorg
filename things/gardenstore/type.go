package gardenstore

import "github.com/dpb587/go-schemaorg"

// // A garden store.
var Type = schemaorg.NewDataType("http://schema.org", "GardenStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package bikestore

import "github.com/dpb587/go-schemaorg"

// // A bike store.
var Type = schemaorg.NewDataType("http://schema.org", "BikeStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

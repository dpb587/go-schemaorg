package store

import "github.com/dpb587/go-schemaorg"

// // A retail good store.
var Type = schemaorg.NewDataType("http://schema.org", "Store")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

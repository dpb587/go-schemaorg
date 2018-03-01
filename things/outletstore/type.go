package outletstore

import "github.com/dpb587/go-schemaorg"

// // An outlet store.
var Type = schemaorg.NewDataType("http://schema.org", "OutletStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

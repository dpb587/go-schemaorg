package electronicsstore

import "github.com/dpb587/go-schemaorg"

// // An electronics store.
var Type = schemaorg.NewDataType("http://schema.org", "ElectronicsStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

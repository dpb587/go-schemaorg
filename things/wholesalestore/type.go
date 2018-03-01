package wholesalestore

import "github.com/dpb587/go-schemaorg"

// // A wholesale store.
var Type = schemaorg.NewDataType("http://schema.org", "WholesaleStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package jewelrystore

import "github.com/dpb587/go-schemaorg"

// // A jewelry store.
var Type = schemaorg.NewDataType("http://schema.org", "JewelryStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package homegoodsstore

import "github.com/dpb587/go-schemaorg"

// // A home goods store.
var Type = schemaorg.NewDataType("http://schema.org", "HomeGoodsStore")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

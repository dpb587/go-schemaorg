package shoppingcenter

import "github.com/dpb587/go-schemaorg"

// // A shopping center or mall.
var Type = schemaorg.NewDataType("http://schema.org", "ShoppingCenter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

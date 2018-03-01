package icecreamshop

import "github.com/dpb587/go-schemaorg"

// // An ice cream shop.
var Type = schemaorg.NewDataType("http://schema.org", "IceCreamShop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

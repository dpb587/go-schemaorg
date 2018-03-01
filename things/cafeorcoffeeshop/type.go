package cafeorcoffeeshop

import "github.com/dpb587/go-schemaorg"

// // A cafe or coffee shop.
var Type = schemaorg.NewDataType("http://schema.org", "CafeOrCoffeeShop")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

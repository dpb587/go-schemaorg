package fastfoodrestaurant

import "github.com/dpb587/go-schemaorg"

// // A fast-food restaurant.
var Type = schemaorg.NewDataType("http://schema.org", "FastFoodRestaurant")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

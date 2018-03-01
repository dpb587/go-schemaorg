package foodestablishment

import "github.com/dpb587/go-schemaorg"

// // A food-related business.
var Type = schemaorg.NewDataType("http://schema.org", "FoodEstablishment")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

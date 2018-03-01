package foodservice

import "github.com/dpb587/go-schemaorg"

// // A food service, like breakfast, lunch, or dinner.
var Type = schemaorg.NewDataType("http://schema.org", "FoodService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

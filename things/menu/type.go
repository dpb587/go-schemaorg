package menu

import "github.com/dpb587/go-schemaorg"

// // A structured representation of food or drink items available from a
// FoodEstablishment.
var Type = schemaorg.NewDataType("http://schema.org", "Menu")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

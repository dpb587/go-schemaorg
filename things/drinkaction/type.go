package drinkaction

import "github.com/dpb587/go-schemaorg"

// // The act of swallowing liquids.
var Type = schemaorg.NewDataType("http://schema.org", "DrinkAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

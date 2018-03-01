package cookaction

import "github.com/dpb587/go-schemaorg"

// // The act of producing/preparing food.
var Type = schemaorg.NewDataType("http://schema.org", "CookAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

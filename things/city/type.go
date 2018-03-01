package city

import "github.com/dpb587/go-schemaorg"

// // A city or town.
var Type = schemaorg.NewDataType("http://schema.org", "City")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package restaurant

import "github.com/dpb587/go-schemaorg"

// // A restaurant.
var Type = schemaorg.NewDataType("http://schema.org", "Restaurant")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

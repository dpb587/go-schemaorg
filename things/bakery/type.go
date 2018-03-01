package bakery

import "github.com/dpb587/go-schemaorg"

// // A bakery.
var Type = schemaorg.NewDataType("http://schema.org", "Bakery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

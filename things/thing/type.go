package thing

import "github.com/dpb587/go-schemaorg"

// // The most generic type of item.
var Type = schemaorg.NewDataType("http://schema.org", "Thing")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

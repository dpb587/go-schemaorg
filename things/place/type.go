package place

import "github.com/dpb587/go-schemaorg"

// // Entities that have a somewhat fixed, physical extension.
var Type = schemaorg.NewDataType("http://schema.org", "Place")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

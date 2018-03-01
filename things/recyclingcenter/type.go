package recyclingcenter

import "github.com/dpb587/go-schemaorg"

// // A recycling center.
var Type = schemaorg.NewDataType("http://schema.org", "RecyclingCenter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

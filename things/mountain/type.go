package mountain

import "github.com/dpb587/go-schemaorg"

// // A mountain, like Mount Whitney or Mount Everest.
var Type = schemaorg.NewDataType("http://schema.org", "Mountain")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

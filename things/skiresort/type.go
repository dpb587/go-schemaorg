package skiresort

import "github.com/dpb587/go-schemaorg"

// // A ski resort.
var Type = schemaorg.NewDataType("http://schema.org", "SkiResort")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

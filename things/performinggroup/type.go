package performinggroup

import "github.com/dpb587/go-schemaorg"

// // A performance group, such as a band, an orchestra, or a circus.
var Type = schemaorg.NewDataType("http://schema.org", "PerformingGroup")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

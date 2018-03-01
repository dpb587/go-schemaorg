package performancerole

import "github.com/dpb587/go-schemaorg"

// // A PerformanceRole is a Role that some entity places with regard to a
// theatrical performance, e.g. in a Movie, TVSeries etc.
var Type = schemaorg.NewDataType("http://schema.org", "PerformanceRole")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

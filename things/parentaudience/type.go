package parentaudience

import "github.com/dpb587/go-schemaorg"

// // A set of characteristics describing parents, who can be interested in viewing
// some content.
var Type = schemaorg.NewDataType("http://schema.org", "ParentAudience")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

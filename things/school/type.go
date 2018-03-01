package school

import "github.com/dpb587/go-schemaorg"

// // A school.
var Type = schemaorg.NewDataType("http://schema.org", "School")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

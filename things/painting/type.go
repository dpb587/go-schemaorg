package painting

import "github.com/dpb587/go-schemaorg"

// // A painting.
var Type = schemaorg.NewDataType("http://schema.org", "Painting")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package photograph

import "github.com/dpb587/go-schemaorg"

// // A photograph.
var Type = schemaorg.NewDataType("http://schema.org", "Photograph")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

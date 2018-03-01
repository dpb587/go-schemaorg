package pond

import "github.com/dpb587/go-schemaorg"

// // A pond.
var Type = schemaorg.NewDataType("http://schema.org", "Pond")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

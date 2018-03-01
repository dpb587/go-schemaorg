package church

import "github.com/dpb587/go-schemaorg"

// // A church.
var Type = schemaorg.NewDataType("http://schema.org", "Church")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

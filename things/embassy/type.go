package embassy

import "github.com/dpb587/go-schemaorg"

// // An embassy.
var Type = schemaorg.NewDataType("http://schema.org", "Embassy")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

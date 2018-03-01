package hospital

import "github.com/dpb587/go-schemaorg"

// // A hospital.
var Type = schemaorg.NewDataType("http://schema.org", "Hospital")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

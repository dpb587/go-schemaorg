package dentist

import "github.com/dpb587/go-schemaorg"

// // A dentist.
var Type = schemaorg.NewDataType("http://schema.org", "Dentist")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

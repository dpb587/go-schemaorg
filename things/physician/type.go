package physician

import "github.com/dpb587/go-schemaorg"

// // A doctor's office.
var Type = schemaorg.NewDataType("http://schema.org", "Physician")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

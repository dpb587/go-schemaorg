package pharmacy

import "github.com/dpb587/go-schemaorg"

// // A pharmacy or drugstore.
var Type = schemaorg.NewDataType("http://schema.org", "Pharmacy")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

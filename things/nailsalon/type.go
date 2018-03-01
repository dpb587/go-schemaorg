package nailsalon

import "github.com/dpb587/go-schemaorg"

// // A nail salon.
var Type = schemaorg.NewDataType("http://schema.org", "NailSalon")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

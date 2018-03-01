package preschool

import "github.com/dpb587/go-schemaorg"

// // A preschool.
var Type = schemaorg.NewDataType("http://schema.org", "Preschool")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

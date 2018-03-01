package singlefamilyresidence

import "github.com/dpb587/go-schemaorg"

// // Residence type: Single-family home.
var Type = schemaorg.NewDataType("http://schema.org", "SingleFamilyResidence")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

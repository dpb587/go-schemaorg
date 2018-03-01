package ngo

import "github.com/dpb587/go-schemaorg"

// // Organization: Non-governmental Organization.
var Type = schemaorg.NewDataType("http://schema.org", "NGO")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

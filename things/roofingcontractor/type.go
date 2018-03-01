package roofingcontractor

import "github.com/dpb587/go-schemaorg"

// // A roofing contractor.
var Type = schemaorg.NewDataType("http://schema.org", "RoofingContractor")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

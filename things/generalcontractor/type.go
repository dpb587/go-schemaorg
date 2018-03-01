package generalcontractor

import "github.com/dpb587/go-schemaorg"

// // A general contractor.
var Type = schemaorg.NewDataType("http://schema.org", "GeneralContractor")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

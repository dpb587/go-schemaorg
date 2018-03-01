package employmentagency

import "github.com/dpb587/go-schemaorg"

// // An employment agency.
var Type = schemaorg.NewDataType("http://schema.org", "EmploymentAgency")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

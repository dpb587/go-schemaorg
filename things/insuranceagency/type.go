package insuranceagency

import "github.com/dpb587/go-schemaorg"

// // An Insurance agency.
var Type = schemaorg.NewDataType("http://schema.org", "InsuranceAgency")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package donateaction

import "github.com/dpb587/go-schemaorg"

// // The act of providing goods, services, or money without compensation, often
// for philanthropic reasons.
var Type = schemaorg.NewDataType("http://schema.org", "DonateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

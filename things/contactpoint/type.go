package contactpoint

import "github.com/dpb587/go-schemaorg"

// // A contact point&#x2014;for example, a Customer Complaints department.
var Type = schemaorg.NewDataType("http://schema.org", "ContactPoint")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package childcare

import "github.com/dpb587/go-schemaorg"

// // A Childcare center.
var Type = schemaorg.NewDataType("http://schema.org", "ChildCare")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

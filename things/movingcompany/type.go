package movingcompany

import "github.com/dpb587/go-schemaorg"

// // A moving company.
var Type = schemaorg.NewDataType("http://schema.org", "MovingCompany")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

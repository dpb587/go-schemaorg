package autodealer

import "github.com/dpb587/go-schemaorg"

// // An car dealership.
var Type = schemaorg.NewDataType("http://schema.org", "AutoDealer")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

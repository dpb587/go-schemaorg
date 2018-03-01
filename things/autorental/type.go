package autorental

import "github.com/dpb587/go-schemaorg"

// // A car rental business.
var Type = schemaorg.NewDataType("http://schema.org", "AutoRental")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

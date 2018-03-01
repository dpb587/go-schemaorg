package taxistand

import "github.com/dpb587/go-schemaorg"

// // A taxi stand.
var Type = schemaorg.NewDataType("http://schema.org", "TaxiStand")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

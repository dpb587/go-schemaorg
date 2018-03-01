package taxiservice

import "github.com/dpb587/go-schemaorg"

// // A service for a vehicle for hire with a driver for local travel. Fares are
// usually calculated based on distance traveled.
var Type = schemaorg.NewDataType("http://schema.org", "TaxiService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

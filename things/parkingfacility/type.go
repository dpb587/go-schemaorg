package parkingfacility

import "github.com/dpb587/go-schemaorg"

// // A parking lot or other parking facility.
var Type = schemaorg.NewDataType("http://schema.org", "ParkingFacility")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

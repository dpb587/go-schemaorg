package travelagency

import "github.com/dpb587/go-schemaorg"

// // A travel agency.
var Type = schemaorg.NewDataType("http://schema.org", "TravelAgency")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

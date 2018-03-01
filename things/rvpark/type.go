package rvpark

import "github.com/dpb587/go-schemaorg"

// // A place offering space for "Recreational Vehicles", Caravans, mobile homes
// and the like.
var Type = schemaorg.NewDataType("http://schema.org", "RVPark")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

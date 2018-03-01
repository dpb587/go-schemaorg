package travelaction

import "github.com/dpb587/go-schemaorg"

// // The act of traveling from an fromLocation to a destination by a specified
// mode of transport, optionally with participants.
var Type = schemaorg.NewDataType("http://schema.org", "TravelAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

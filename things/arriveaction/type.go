package arriveaction

import "github.com/dpb587/go-schemaorg"

// // The act of arriving at a place. An agent arrives at a destination from a
// fromLocation, optionally with participants.
var Type = schemaorg.NewDataType("http://schema.org", "ArriveAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

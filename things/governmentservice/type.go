package governmentservice

import "github.com/dpb587/go-schemaorg"

// // A service provided by a government organization, e.g. food stamps, veterans
// benefits, etc.
var Type = schemaorg.NewDataType("http://schema.org", "GovernmentService")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

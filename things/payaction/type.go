package payaction

import "github.com/dpb587/go-schemaorg"

// // An agent pays a price to a participant.
var Type = schemaorg.NewDataType("http://schema.org", "PayAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

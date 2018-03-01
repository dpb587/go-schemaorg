package eventstatustype

import "github.com/dpb587/go-schemaorg"

// // EventStatusType is an enumeration type whose instances represent several
// states that an Event may be in.
var Type = schemaorg.NewDataType("http://schema.org", "EventStatusType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

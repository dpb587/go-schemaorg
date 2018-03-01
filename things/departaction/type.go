package departaction

import "github.com/dpb587/go-schemaorg"

// // The act of  departing from a place. An agent departs from an fromLocation for
// a destination, optionally with participants.
var Type = schemaorg.NewDataType("http://schema.org", "DepartAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

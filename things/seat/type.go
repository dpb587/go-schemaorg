package seat

import "github.com/dpb587/go-schemaorg"

// // Used to describe a seat, such as a reserved seat in an event reservation.
var Type = schemaorg.NewDataType("http://schema.org", "Seat")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

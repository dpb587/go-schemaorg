package reservationpackage

import "github.com/dpb587/go-schemaorg"

// // A group of multiple reservations with common values for all sub-reservations.
var Type = schemaorg.NewDataType("http://schema.org", "ReservationPackage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

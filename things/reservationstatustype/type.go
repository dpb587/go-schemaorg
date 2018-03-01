package reservationstatustype

import "github.com/dpb587/go-schemaorg"

// // Enumerated status values for Reservation.
var Type = schemaorg.NewDataType("http://schema.org", "ReservationStatusType")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

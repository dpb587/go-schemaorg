package reservationpackage

import "github.com/dpb587/go-schemaorg"

var (
	// The individual reservations included in the package. Typically a repeated
	// property.
	SubReservation = schemaorg.NewProperty("subReservation")
)

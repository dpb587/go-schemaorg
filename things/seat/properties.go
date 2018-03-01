package seat

import "github.com/dpb587/go-schemaorg"

var (
	// The type/class of the seat.
	SeatingType = schemaorg.NewProperty("seatingType")

	// The section location of the reserved seat (e.g. Orchestra).
	SeatSection = schemaorg.NewProperty("seatSection")

	// The row location of the reserved seat (e.g., B).
	SeatRow = schemaorg.NewProperty("seatRow")

	// The location of the reserved seat (e.g., 27).
	SeatNumber = schemaorg.NewProperty("seatNumber")
)

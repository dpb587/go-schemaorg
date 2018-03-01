package house

import "github.com/dpb587/go-schemaorg"

var (
	// The number of rooms (excluding bathrooms and closets) of the acccommodation
	// or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can
	// be put in the unitText property of the QuantitativeValue.
	NumberOfRooms = schemaorg.NewProperty("numberOfRooms")
)
